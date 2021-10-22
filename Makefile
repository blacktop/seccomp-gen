REPO=blacktop
NAME=seccomp-gen
CUR_VERSION=$(shell svu current)
NEXT_VERSION=$(shell svu patch)


.PHONY: test
test: ## Run all the tests
	gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=30s

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: run
run: dry_release
	cat wdef.strace | dist/darwin_amd64/scgen --verbose

.PHONY: run
run_docker: dry_release
	docker run --rm --security-opt="seccomp=unconfined" $(REPO)/$(NAME):test 2>&1 1>/dev/null | tail -n +3 | head -n -2 | awk '{print $(NF)}' | dist/darwin_amd64/scgen --verbose
	docker run --rm --security-opt="no-new-privileges" --security-opt="seccomp=seccomp.json" $(REPO)/$(NAME):test

.PHONY: dry_release
dry_release: ## Run goreleaser without releasing/pushing artifacts to github
	@echo " > Creating Pre-release Build ${NEXT_VERSION}"
	@goreleaser build --rm-dist --skip-validate --id darwin

.PHONY: snapshot
snapshot: ## Run goreleaser snapshot
	@echo " > Creating Snapshot ${NEXT_VERSION}"
	@goreleaser --rm-dist --snapshot

.PHONY: release
release: ## Create a new release from the NEXT_VERSION
	@echo " > Creating Release ${NEXT_VERSION}"
	@hack/make/release ${NEXT_VERSION}
	@goreleaser --rm-dist

.PHONY: destroy
destroy: ## Remove release for the CUR_VERSION
	@echo " > Deleting Release"
	git tag -d ${CUR_VERSION}
	git push origin :refs/tags/${CUR_VERSION}

ci: lint test ## Run all the tests and code checks

clean: ## Clean up artifacts
	rm seccomp.json || true
	rm -rf dist

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help