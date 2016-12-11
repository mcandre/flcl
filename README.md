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

$ flcl -h
Usage:
  flcl [options] <path>...
  flcl -h --help
  flcl -v --version

  Arguments:
    <path>                    A file path. Directories are traversed recursively. Nearby .gitignore's are applied.
  Options:
    -c --charsets <charsets>  Limit results to certain character sets [default: ascii,utf-8]
    -h --help                 Show usage information
    -v --version              Show version information
```

# DOWNLOADS

https://github.com/mcandre/flcl/releases

# DEVELOPMENT REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [editorconfig-tools](https://www.npmjs.com/package/editorconfig-tools)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/flcl/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.m/mcandre
$ git clone git@github.com:mcandre/go-chop.git $GOPATH/src/github.com/mcandre/go-chop
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-chop/cmd/chop && go install"
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-chop/cmd/chomp && go install"
```

# TEST REMOTELY

```
$ go test github.com/mcandre/go-chop
```

# TEST LOCALLY

```
$ go test
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
