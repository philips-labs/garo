VERSION := 0.0.0-dev
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
GITCOMMIT := $(GITCOMMIT)-dirty
endif
CTIMEVAR=-X main.commit=$(GITCOMMIT) -X main.version=$(VERSION) -X main.date=$(shell date +%FT%TZ)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
PLANTUML_JAR_URL = https://sourceforge.net/projects/plantuml/files/plantuml.jar/download
DIAGRAMS_SRC := $(wildcard docs/diagrams/*.plantuml)
DIAGRAMS_PNG := $(addsuffix .png, $(basename $(DIAGRAMS_SRC)))
DIAGRAMS_SVG := $(addsuffix .svg, $(basename $(DIAGRAMS_SRC)))
GOPATH := $(word 2, $(subst =, , $(shell go env | grep GOPATH)))

.PHONY: help clean clean-diagrams clean-binaries diagrams png-diagrams svg-diagrams compile compile-agent compile-server test test-cover coverage-out coverage-html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

clean: clean-diagrams clean-binaries ## Clean binaries and diagrams
clean-diagrams: ## Cleans plantuml.jar and generated diagrams
	@rm -f plantuml.jar $(DIAGRAMS_PNG) $(DIAGRAMS_SVG)
clean-binaries: ## Cleans binaries
	@rm -f bin/*

diagrams: svg-diagrams png-diagrams ## Generate diagrams in SVG and PNG format
svg-diagrams: plantuml.jar $(DIAGRAMS_SVG) ## Generate diagrams in SVG format
png-diagrams: plantuml.jar $(DIAGRAMS_PNG) ## Generate diagrams in PNG format

plantuml.jar:
	@echo Downloading $@....
	@curl -sSfL $(PLANTUML_JAR_URL) -o $@

docs/diagrams/%.svg: docs/diagrams/%.plantuml
	@echo Generating $^ from plantuml....
	@java -jar plantuml.jar -tsvg $^

docs/diagrams/%.png: docs/diagrams/%.plantuml
	@echo Generating $^ from plantuml....
	@java -jar plantuml.jar -tpng $^

compile: compile-agent compile-server ## Compile garo-agent and garo-server
compile-agent: bin/garo-agent ## Compile garo-agent
compile-server: bin/garo-server ## Compile garo-server

bin/garo-agent:
	@echo Compiling $@...
	@go build -a ${GO_LDFLAGS} -o $@ ./cmd/garo-agent

bin/garo-server:
	@echo Compiling $@...
	@go build -a ${GO_LDFLAGS} -o $@ ./cmd/garo-server

test: ## Run tests
	@echo Running tests...
	@go test -v -race -count=1 ./...

test-cover: ## Run tests and coverage
	@echo Running tests and coverage...
	@go test -v -race -count=1 -covermode=atomic -coverprofile=coverage.out ./...

coverage-out: test-cover ## Show coverage in cli
	@echo Coverage details
	@go tool cover -func=coverage.out

coverage-html: test-cover ## Show coverage in browser
	@go tool cover -html=coverage.out

download: ## Fetches go.mod dependencies via go mod download
	@go mod download

install-protoc: ## Installs protoc
	@brew list | grep protobuf > /dev/null && brew upgrade protobuf > /dev/null || brew install protobuf > /dev/null

install-tools: download ## Installs tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

proto-gen: ## Generate protocol buffer implementations
	@protoc --proto_path=$(GOPATH)/pkg/mod:. --twirp_out=. --go_out=. ./rpc/garo/*.proto
	@goimports -e -local github.com/philips-labs -w rpc/garo/service.twirp.go rpc/garo/service.twirp.go
	@goimports -e -local github.com/philips-labs -w rpc/garo/service.pb.go rpc/garo/service.pb.go
