GO ?= go
BINARY := bin/khs
GO_FILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

.PHONY: build test test-race vet fmt-check check

build:
	mkdir -p bin
	$(GO) build -o $(BINARY) ./cmd/khs

test:
	$(GO) test ./...

test-race:
	$(GO) test -race ./...

vet:
	$(GO) vet ./...

fmt-check:
	@test -z "$$(gofmt -l $(GO_FILES))" || { \
		echo "Go files need formatting:"; \
		gofmt -l $(GO_FILES); \
		exit 1; \
	}

check: fmt-check vet test
