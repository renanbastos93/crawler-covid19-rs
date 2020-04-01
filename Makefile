GO := $(shell which go)
GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin
GOSEC ?= $(GOBIN)/gosec
GOLINT ?= $(GOBIN)/golint

all: format lint sec test

run: 
	@ $(GO) run ./cmd/main.go

sec:
	@ $(GOSEC) ./...

lint:
	@ $(GOLINT) ./...

format:
	@ $(GO)fmt -w .	

test:
	@ $(GO) test -coverprofile "coverage.html" .