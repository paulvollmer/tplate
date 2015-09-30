.PHONY: build test clean

VERSION := 0.1.0
REPO_REV := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date +%Y-%m-%d:%H:%M:%S)
BUILD_FLAGS := -ldflags="-X main.version=$(VERSION) -X main.gitRev=$(REPO_REV) -X main.buildDate=$(BUILD_DATE)"

build:
	@go build $(BUILD_FLAGS)

install:
	@go install $(BUILD_FLAGS)

test:
	@go test -cover -v

clean:
	@rm -f tplate
	@rm -f tplate_darwin_*
	@rm -f tplate_freebsd_*
	@rm -f tplate_linux_*
	@rm -f tplate_netbsd_*
	@rm -f tplate_openbsd_*
	@rm -f tplate_windows_*
