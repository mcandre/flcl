APP=flcl
VERSION=0.0.1

.PHONY: port clean clean-ports

all: integration-test

integration-test: bin
	flcl bin 2>&1 | grep 'No results'

govet:
	go list ./... | grep -v vendor | xargs go vet -v

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

editorconfig:
	sh editorconfig.sh

lint: govet gofmt goimport

port: archive-ports

archive-ports: bin
	zipc -C bin "$(APP)-$(VERSION).zip" "$(APP)-$(VERSION)"

bin:
	gox -output="bin/$(APP)-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd...

clean: clean-ports

clean-ports:
	rm -rf bin
