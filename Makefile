
# Makefile

PACKAGES=$(shell go list ./... | grep -v '/simulation')
#VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
VERSION := $(shell echo $(shell git describe --always) | sed 's/^v//')
#"git log -1": This returns the most recent commit on a branch.
COMMIT := $(shell git log -1 --format='%H')
DAEMON_BINARY = middlewared
PRJ_NAME = middlewared
GOBIN = $(GOPATH)/bin

ldflags = -X github.com/reapchain/reapchain-sdk/version.Name=reapchain \
		  -X github.com/reapchain/reapchain-sdk/version.ServerName=$(DAEMON_BINARY) \
		  -X github.com/reapchain/reapchain-sdk/version.Version=$(VERSION) \
		  -X github.com/reapchain/reapchain-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

build: go.sum
	@echo "--> Building middlewared to ./build"
	@echo "version:" $(VERSION) "commit:" $(COMMIT)
	@go build -mod=readonly -o build/$(DAEMON_BINARY) $(BUILD_FLAGS) ./cmd/$(PRJ_NAME)

install: go.sum
	@echo "--> Installing middlewared"
	@echo "version:" $(VERSION) "commit:" $(COMMIT)
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/$(PRJ_NAME)
	@mv $(GOBIN)/$(PRJ_NAME) $(GOBIN)/$(DAEMON_BINARY)

clean:
	@rm -rf ./build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
