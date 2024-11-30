GOIMPORTS := go run golang.org/x/tools/cmd/goimports@latest
GOFUMPT := go run mvdan.cc/gofumpt@latest
GOLINES := go run github.com/segmentio/golines@latest
GOLANGCI := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest

## build: Build an application
.PHONY: build
build:
	go build -o main -v cmd/main.go

## run: Run application
.PHONY: run
run:
	go run cmd/main.go

## generate: Regenerate all required files
.PHONY: generate
generate:
	go generate ./...

## test: Launch unit tests
.PHONY: test
test:
	go generate ./...
	go test ./...

## tidy: Cleanup go.sum and go.mod files
.PHONY: tidy
tidy:
	go mod tidy

## lint: Launch project linters
.PHONY: lint
lint:
	$(GOLANGCI) run --timeout 360s

## fmt: Reformat source code
.PHONY: fmt
fmt:
	$(GOIMPORTS) -w -l .
	$(GOFUMPT) -w -l .
	$(GOLINES) -w --no-reformat-tags --max-len=120 .

.PHONY: check
check: generate fmt lint test tidy

## help: Prints help message
.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' | sort