// +build mage

package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/flcl"
	"github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// Test executes the integration test suite.
func Test() error {
	mg.Deps(Install)

	if err := os.MkdirAll(artifactsPath, os.ModeDir|0775); err != nil {
		return err
	}

	var out bytes.Buffer

	cmd := exec.Command("flcl", "bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = bufio.NewWriter(&out)

	if err := cmd.Run(); err != nil {
		return err
	}

	errorOutput := out.String()

	if !strings.Contains(errorOutput, "No results") {
		return errors.New(fmt.Sprintf("Expected \"No results\", got %v", errorOutput))
	}

	return nil
}

// GoVet runs go tool vet.
func GoVet() error { return mageextras.GoVet("-shadow") }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("flcl-%s", flcl.Version)

// Port archives build artifacts.
func Port() error { mg.Deps(Gox); return mageextras.Archive(portBasename, artifactsPath) }

// Gox cross-compiles Go binaries.
func Gox() error {
	return mageextras.Gox(
		artifactsPath,
		strings.Join(
			[]string{
				portBasename,
				"{{.OS}}",
				"{{.Arch}}",
				"{{.Dir}}",
			},
			mageextras.PathSeparatorString,
		),
	)
}

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("flcl") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
