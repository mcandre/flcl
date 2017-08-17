APP=flcl
VERSION=0.0.2

.PHONY: port clean clean-ports

all: integration-test

integration-test: bin
	flcl bin 2>&1 | grep 'No results'

govet:
	find . -path "*/vendor*" -prune -o -name "*.go" -type f -exec go tool vet -shadow {} \;

golint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec golint {} \;

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

errcheck:
	errcheck -blank

opennota-check:
	aligncheck
	structcheck
	varcheck

megacheck:
	megacheck

editorconfig:
	sh editorconfig.sh

lint: govet golint gofmt goimport errcheck opennota-check megacheck editorconfig

port: archive-ports

archive-ports: bin
	zipc -C bin "$(APP)-$(VERSION).zip" "$(APP)-$(VERSION)"

bin:
	gox -output="bin/$(APP)-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd...

clean: clean-ports

clean-ports:
	rm -rf bin
