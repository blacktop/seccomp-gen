package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/apex/log"
	clihander "github.com/apex/log/handlers/cli"
	"github.com/blacktop/seccomp-gen/seccomp"
	"github.com/blacktop/seccomp-gen/seccomp/syscalls"
)

var requiredSyscalls = []string{
	"capget",
	"capset",
	"chdir",
	"fchown",
	"futex",
	"getdents64",
	"getpid",
	"getppid",
	"lstat",
	"openat",
	"prctl",
	"setgid",
	"setgroups",
	"setuid",
	"stat",
}

func unique(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func init() {
	log.SetHandler(clihander.Default)
}

func main() {

	var sc string
	var scs []string

	verbosePtr := flag.Bool("verbose", false, "verbose output")
	flag.Parse()

	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	}

	re := regexp.MustCompile(`^[a-zA-Z_]+\(`)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sc = re.FindString(scanner.Text())
		if len(sc) > 0 {
			if syscalls.IsValid(strings.TrimRight(sc, "(")) {
				scs = unique(scs, strings.TrimRight(sc, "("))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for _, sc := range scs {
		log.Debugf("found syscall: %s", sc)
	}

	requiredSyscalls = append(requiredSyscalls, scs...)
	sort.Strings(requiredSyscalls)

	// write out to file
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f := filepath.Join(wd, "seccomp.json")

	// write the default profile to the file
	b, err := json.MarshalIndent(seccomp.DefaultProfile(requiredSyscalls), "", "\t")
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(f, b, 0644); err != nil {
		panic(err)
	}
}
