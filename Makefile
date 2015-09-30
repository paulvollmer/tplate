NAME := tplate
VERSION := 0.1.0
REPO_REV := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date +%Y-%m-%d:%H:%M:%S)
BUILD_FLAGS := -ldflags="-X main.version=$(VERSION) -X main.gitRev=$(REPO_REV) -X main.buildDate=$(BUILD_DATE)"

define buildRelease
@gox -os="$1" -arch="$2" -output="$(NAME)"
@tar cfvz $(NAME)-$(VERSION)-$1-$2.tar.gz $(NAME)
endef

build:
	@go build $(BUILD_FLAGS)

install:
	@go install $(BUILD_FLAGS)

test:
	@go tool vet *.go
	@go test -cover -v

release:
	$(MAKE) release-darwin-386
	$(MAKE) release-darwin-amd64
	$(MAKE) release-linux-386
	$(MAKE) release-linux-amd64
	$(MAKE) release-linux-arm
	$(MAKE) release-windows-386
	$(MAKE) release-windows-amd64

release-darwin-386:
	$(call buildRelease,darwin,386)

release-darwin-amd64:
	$(call buildRelease,darwin,amd64)

release-linux-386:
	$(call buildRelease,linux,386)

release-linux-amd64:
	$(call buildRelease,linux,amd64)

release-linux-arm:
	$(call buildRelease,linux,arm)

release-windows-386:
	$(call buildRelease,windows,386)

release-windows-amd64:
	$(call buildRelease,windows,amd64)

clean:
	@rm -f *.tar.gz
	@rm -f $(NAME)
	@rm -f $(NAME).exe
	@rm -f $(NAME)_darwin_*
	@rm -f $(NAME)_freebsd_*
	@rm -f $(NAME)_linux_*
	@rm -f $(NAME)_netbsd_*
	@rm -f $(NAME)_openbsd_*
	@rm -f $(NAME)_windows_*
