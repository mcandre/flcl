# flcl - what if UNIX find had -charset utf-8... wouldn't that be cool?

# EXAMPLES

```console
$ find .
.
./.editorconfig
./.envrc
./.envrc.sample
./.git
./.git/COMMIT_EDITMSG
./.git/config
./.git/description
./.git/HEAD
./.git/hooks
...

$ flcl .
.editorconfig
.envrc.sample
.git
.gitignore
.node-version
Makefile
README.md
cmd/flcl/main.go
drbrule.gif
editorconfig.sh
flcl.go
package.json

$ flcl -help
  -charsets string
        Limit results to comma-separated character sets (default "ascii,utf-8")
  -help
        Show usage information
  -version
        Show version information
```

# DOWNLOADS

https://github.com/mcandre/flcl/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/flcl

# RUNTIME REQUIREMENTS

(None)

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/flcl/cmd/flcl@latest
```

# CONTRIBUTING

For more information on developing flcl itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
