SHELL = /bin/bash
CURRENT_DIRECTORY = $(shell pwd)

# Build variables
REPORTS_DIR ?= reports
CI_SERVICE ?=
COVERALLS_TOKEN ?=

# Go variables
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go")

.PHONY: all
all: dep ## Get deps and build binary

.PHONY: clean
clean: ## Clean the working area and the project
	rm -rf $(REPORTS_DIR)

.PHONY: dep
dep: ## Install dependencies
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls
	@go mod tidy
	@go get -v -t ./...

.PHONY: format
format: ## Format the source
	@goimports -w $(GOFILES)

.PHONY: test
test: ## Run unit tests
	mkdir -p $(REPORTS_DIR)
	rm -f $(REPORTS_DIR)/*
	@go test -v -covermode=count -coverprofile="$(REPORTS_DIR)/coverage.out" ./...

.PHONY: publish-test-results
publish-test-results: ## Publish test results
	@goveralls -coverprofile="$(REPORTS_DIR)/coverage.out" -service=$(CI_SERVICE) -repotoken $(COVERALLS_TOKEN)

.PHONY: list
list: ## List all make targets
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help: ## Get help output
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)

