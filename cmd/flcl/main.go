package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/mcandre/go-chop"
	"github.com/mcandre/flcl"
	"github.com/monochromegane/go-gitignore"
)

const Usage = `Usage:
  flcl [options] <path>...
  flcl -h --help
  flcl -v --version

  Arguments:
    <path>                    A file path.

                              - Directories are traversed recursively.
                              - Nearby .gitignore's are applied.
  Options:
    -c --charsets <charsets>  Limit results to certain character sets [default: ascii,utf-8]
    -h --help                 Show usage information
    -v --version              Show version information
`

const OriginDir = "/"

//
// Work around go-gitignore's overly strict directory trailing slash semantics
//
func flexibleMatch(ignores gitignore.IgnoreMatcher, root string) bool {
	return ignores == nil ||
		(!ignores.Match(root, false) &&
			!ignores.Match(root, true) &&
			!ignores.Match(root + "/", true))
}

func populate(visited map[string]bool, gitignores map[string]gitignore.IgnoreMatcher, dir string) {
	if !visited[dir] {
		candidate := path.Join(dir, ".gitignore")

		gitignoreMatcher, err := gitignore.NewGitIgnore(candidate, OriginDir)

		if err != nil {
			parent := path.Dir(dir)

			if parent != dir {
				populate(visited, gitignores, parent)
				gitignores[dir] = gitignores[parent]
			}
		} else {
			gitignores[dir] = gitignoreMatcher
		}

		visited[dir] = true
	}
}

func process(visited map[string]bool, gitignores map[string]gitignore.IgnoreMatcher, gitignoreGlobal gitignore.IgnoreMatcher, root string, charsets []string, foundResult *bool) {
	rootInfo, err := os.Stat(root)

	if err != nil {
		log.Printf("Cannot access path: %s\n", root)
		return
	}

	rootIsDir := rootInfo.IsDir()

	parent := path.Dir(root)
	populate(visited, gitignores, parent)
	gitignoreMatcher := gitignores[parent]

	rootBase := path.Base(root)

	rootAbs, err := filepath.Abs(root)

	if err != nil {
		panic(err)
	}

	if flexibleMatch(gitignoreGlobal, rootAbs) {
		if flexibleMatch(gitignoreMatcher, rootAbs) {
			if rootIsDir && rootBase != ".git" {
				childInfos, err := ioutil.ReadDir(root)

				if err != nil {
					panic(err)
				}

				for _, childInfo := range childInfos {
					childRel := path.Join(root, childInfo.Name())
					process(visited, gitignores, gitignoreGlobal, childRel, charsets, foundResult)
				}
			} else if rootBase != ".gitmodules" {
				*foundResult = true

				var rootQuoted string
				if strings.ContainsAny(root, " '\"") {
					rootQuoted = strconv.Quote(root)
				} else {
					rootQuoted = root
				}

				fmt.Printf("%s\n", rootQuoted)
			}
		}
	}
}

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, flcl.Version, false)

	if err != nil {
		panic(Usage)
	}

	paths, _ := arguments["<path>"].([]string)

	charsetCommas, _ := arguments["--charsets"].(string)

	visited := make(map[string]bool)
	gitignores := make(map[string]gitignore.IgnoreMatcher)

	gitConfigCommand := exec.Command("git", "config", "--global", "core.excludesfile")
	gitConfigCommand.Stderr = os.Stderr
	gitignoreGlobalPathBytes, err := gitConfigCommand.Output()
	var gitignoreGlobal gitignore.IgnoreMatcher

	if err != nil {
		log.Printf("Unable to identify global gitignore: %v\n", err)
	} else {
		gitignoreGlobal, _ = gitignore.NewGitIgnore(chop.Chomp(string(gitignoreGlobalPathBytes)), OriginDir)
	}

	charsets := strings.Split(charsetCommas, ",")

	foundResult := false

	for _, p := range paths {
		process(visited, gitignores, gitignoreGlobal, p, charsets, &foundResult)
	}

	if !foundResult {
		log.Printf("No results")
	}
}
