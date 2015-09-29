.PHONY: build test clean

VERSION := v0.1.0
REPO_REV := $(shell git rev-parse HEAD)

build:
	@go build -ldflags "-X main.version=$(VERSION) -X main.gitRev=$(REPO_REV)"

test:
	@go test -cover -v

clean:
	@rm tplate
