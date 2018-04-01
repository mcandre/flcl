# flcl - what if UNIX find had -charset utf-8... wouldn't that be cool?

# EXAMPLES

```
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

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.9+
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/flcl/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.m/mcandre
$ git clone https://github.com/mcandre/flcl.git $GOPATH/src/github.com/mcandre/flcl
$ cd "$GOPATH/src/github.com/mcandre/flcl"
$ git submodule update --init --recursive
$ go install
```

# INTEGRATION TEST

```
$ mage test
```

# PORT

```
$ mage port
```

# LINT

Keep the code tidy:

```
$ mage lint
```
