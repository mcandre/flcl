// Package main provides an executable for running content-aware flcl searches, like UNIX find.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mcandre/flcl"
	"github.com/mcandre/go-chop"
	"github.com/monochromegane/go-gitignore"
)

var flagCharsets = flag.String("charsets", "ascii,utf-8", "Limit results to comma-separated character sets")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// OriginDir presents a base case for recursive file walking: the root directory.
const OriginDir = "/"

// flexibleMatch works around go-gitignore's quite strict directory trailing slash semantics.
func flexibleMatch(ignores gitignore.IgnoreMatcher, root string) bool {
	return ignores == nil ||
		!(ignores.Match(root, false) ||
			ignores.Match(root, true) ||
			ignores.Match(root+"/", true))
}

// populate identifies the gitignore patterns to apply for some directory path, falling back on parent directory patterns, when available.
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

// process recursively identifies viable file paths, omitting those that would be ignored by git.
func process(visited map[string]bool, gitignores map[string]gitignore.IgnoreMatcher, gitignoreGlobal gitignore.IgnoreMatcher, root string, charsets []string, foundResult *bool) {
	rootInfo, err := os.Stat(root)

	if err != nil {
		log.Panic(err)
	}

	rootIsDir := rootInfo.IsDir()

	parent := path.Dir(root)
	populate(visited, gitignores, parent)
	gitignoreMatcher := gitignores[parent]

	rootBase := path.Base(root)

	rootAbs, err := filepath.Abs(root)

	if err != nil {
		log.Panic(err)
	}

	if flexibleMatch(gitignoreGlobal, rootAbs) {
		if flexibleMatch(gitignoreMatcher, rootAbs) {
			if rootIsDir && rootBase != ".git" {
				childInfos, err := ioutil.ReadDir(root)

				if err != nil {
					log.Panic(err)
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

				fmt.Println(rootQuoted)
			}
		}
	}
}

// main is the command line entry point for launching flcl commands.
func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Println(flcl.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	}

	paths := flag.Args()

	visited := make(map[string]bool)
	gitignores := make(map[string]gitignore.IgnoreMatcher)

	gitConfigCommand := exec.Command("git", "config", "--global", "core.excludesfile")
	gitConfigCommand.Stderr = os.Stderr
	gitignoreGlobalPathBytes, err := gitConfigCommand.Output()
	var gitignoreGlobal gitignore.IgnoreMatcher

	if err != nil {
		log.Println(err)
	} else {
		gitignoreGlobal, _ = gitignore.NewGitIgnore(chop.Chomp(string(gitignoreGlobalPathBytes)), OriginDir)
	}

	charsets := strings.Split(*flagCharsets, ",")

	foundResult := false

	for _, p := range paths {
		process(visited, gitignores, gitignoreGlobal, p, charsets, &foundResult)
	}

	if !foundResult {
		log.Println("No results")
	}
}
