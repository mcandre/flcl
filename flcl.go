// Package flcl provides a UNIX find-like search capability that can filter by content type.
package flcl

import (
	"path"

	"github.com/monochromegane/go-gitignore"
)

// Version is semver.
const Version = "0.0.3"

// OriginDir presents a base case for recursive file walking: the root directory.
const OriginDir = "/"

// FlexibleMatch works around go-gitignore's quite strict directory trailing slash semantics.
func FlexibleMatch(ignores gitignore.IgnoreMatcher, root string) bool {
	return ignores == nil ||
		!(ignores.Match(root, false) ||
			ignores.Match(root, true) ||
			ignores.Match(root+"/", true))
}

// Populate identifies the gitignore patterns to apply for some directory path, falling back on parent directory patterns, when available.
func Populate(visited map[string]bool, gitignores map[string]gitignore.IgnoreMatcher, dir string) {
	if !visited[dir] {
		candidate := path.Join(dir, ".gitignore")

		gitignoreMatcher, err := gitignore.NewGitIgnore(candidate, OriginDir)

		if err != nil {
			parent := path.Dir(dir)

			if parent != dir {
				Populate(visited, gitignores, parent)
				gitignores[dir] = gitignores[parent]
			}
		} else {
			gitignores[dir] = gitignoreMatcher
		}

		visited[dir] = true
	}
}
