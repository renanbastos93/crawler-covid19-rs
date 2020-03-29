GO := $(shell which go)
GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin
GOSEC ?= $(GOBIN)/gosec
GOLINT ?= $(GOBIN)/golint

all: format lint sec

sec:
	$(GOSEC) ./...

lint:
	$(GOLINT) ./...

format:
	$(GO)fmt -w .	
