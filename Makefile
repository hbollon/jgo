NAME          := <project_name> 
FILES         := $(wildcard */*.go)
VERSION       := $(shell git describe --always)
BIN_DIR 	    := bin/

export GO111MODULE=on

## setup: Install required libraries/tools for build tasks
.PHONY: setup
setup:
	@command -v goimports 2>&1 >/dev/null || GO111MODULE=off go get -u -v golang.org/x/tools/cmd/goimports
	@command -v golangci-lint 2>&1 >/dev/null || GO111MODULE=off go get -v github.com/golangci/golangci-lint/cmd/golangci-lint

## fmt: Format all sources files
.PHONY: fmt
fmt: setup 
	goimports -w $(FILES)

## lint: Run all lint related tests against the codebase (will use the .golangci.yml config)
.PHONY: lint
lint: setup
	golangci-lint run

## test: Run the tests against the codebase
.PHONY: test
test:
	go test -v -race ./...

## build: Build the binary for Linux environement
.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 \
		go build \
		-o ./bin/output-linux-amd64 .

## build-all: Build binaries for all supported platforms
.PHONY: build-all
build-all:
	env GOOS=linux GOARCH=amd64 \
		go build \
		-o ./bin/output-linux-amd64 \
		.

	env GOOS=windows GOARCH=amd64 \
		go build \
		-o ./bin/output-windows-amd64.exe \
		.

	env GOOS=darwin GOARCH=amd64 \
		go build \
		-o ./bin/output-darwin-amd64 \
		.

## release: Build binaries for all platforms
.PHONY: release
release: build-all
	cd ./bin && \
		zip -r output-linux-amd64.zip output-linux-amd64 && \
		zip -r output-windows-amd64.zip output-windows-amd64 && \
		zip -r output-darwin-amd64.zip output-darwin-amd64


## install: Install go dependencies
.PHONY: install
install:
	go get ./...

# vendor: Vendor go modules
.PHONY: vendor
vendor:
	go mod vendor

## coverage: Generates coverage report
.PHONY: coverage
coverage:
	rm -f coverage.out
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out

## clean: Remove binaries (go binaries, bundles and vue dist folder) if they exist
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: all
all: lint test build-all

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

