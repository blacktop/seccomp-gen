#!/bin/bash
set -e
set -o pipefail

main(){
	local file=$1
	local name=$(basename "$0")

	if [[ -z "$file" ]]; then
		cat >&2 <<-EOF
		${name} [strace-output-filename]

		You must pass a filename that has the strace output.
		EOF
	fi

	# get just the syscalls
	local IFS=$'\n'
	raw=( $(perl -lne 'print $1 if /([a-zA-Z_]+\()/' "$file" | sort -u) )
	unset IFS


	syscalls=( )

	tmpfile=$(mktemp /tmp/seccomp-strace.XXXXXX)

	curl -sSL -o "$tmpfile" https://raw.githubusercontent.com/torvalds/linux/master/arch/x86/entry/syscalls/syscall_64.tbl

	for syscall in "${raw[@]}"; do
		# clean the trailing (
		syscall=${syscall%(}

		if grep -R -q -w $syscall "$tmpfile"; then
			syscalls+=( $syscall )
		fi
	done

	curl -sSL -o "$tmpfile" https://raw.githubusercontent.com/torvalds/linux/master/arch/x86/entry/syscalls/syscall_32.tbl

	for syscall in "${raw[@]}"; do
		# clean the trailing (
		syscall=${syscall%(}

		if grep -R -q -w $syscall "$tmpfile"; then
			syscalls+=( $syscall )
		fi
	done

	# start the seccomp profile
	cat <<-EOF > "$tmpfile"
	{
		"defaultAction": "SCMP_ACT_ERRNO",
		"syscalls": [
		EOF

		for syscall in "${syscalls[@]}"; do
			cat <<-EOF
			{
				"name": "${syscall}",
				"action": "SCMP_ACT_ALLOW",
				"args": null
			},
			EOF
		done >> "$tmpfile"

		# remove trailing comma
		sed -i '$s/,$//' "$tmpfile"

		cat <<-EOF >> "$tmpfile"
		]
	}
	EOF

	cat "$tmpfile"
	rm "$tmpfile"
}

main $@