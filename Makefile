GO_BIN := $(GOPATH)/bin
GOIMPORTS := go run golang.org/x/tools/cmd/goimports@latest
GOFUMPT := go run mvdan.cc/gofumpt@latest
GOLINES := go run github.com/segmentio/golines@latest
GOLANGCI := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest
OAPI_CODEGEN := $(GO_BIN)/oapi-codegen
OAPI_MERGER := $(GO_BIN)/oapi-merger
MERGED_OAPI_V1=$(PWD)/api/openapi/v1/merged.json

## build: Build an application
.PHONY: build
build: docs
	go build -o main -v cmd/main.go

## run: Run application
.PHONY: run
run: docs
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

## docs: Regenerate openapi docs
.PHONY: docs
docs: openapi_merge openapi_http

.PHONY: openapi_merge
openapi_merge: $(OAPI_MERGER)
	oapi-merger -wdir api/openapi/v1 -spec openapi.yaml -o $(MERGED_OAPI_V1)

.PHONY: openapi_http
openapi_http: $(OAPI_CODEGEN)
	oapi-codegen --old-config-style  -generate types,skip-prune -o ./internal/view/types.gen.go -package view $(MERGED_OAPI_V1)
	oapi-codegen --old-config-style -generate spec -o ./internal/view/spec.gen.go -package view $(MERGED_OAPI_V1)
	rm -f $(MERGED_OAPI_V1)

$(OAPI_CODEGEN):
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0

$(OAPI_MERGER):
	go install github.com/felicson/oapi-merger/cmd/oapi-merger@v0.0.2

## help: Prints help message
.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' | sort