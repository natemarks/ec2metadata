.DEFAULT_GOAL := help

# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
PKG := github.com/natemarks/ec2metadata
DEFAULT_BRANCH := main
VERSION := 0.0.0
COMMIT := $(shell git rev-parse HEAD)
SHELL := $(shell which bash)
CDIR = $(shell pwd)
CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DEFAULT_BRANCH := main


help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

clean-venv: ## re-create virtual env
	rm -rf .venv
	python3 -m venv .venv
	( \
       source .venv/bin/activate; \
       pip install --upgrade pip setuptools; \
    )

unittest: ## run test that don't require deployed resources
	go test -v ./... -tags unit

deploymenttest: ##  run all tests
	go test -v ./...

static: lint ## run fmt, vet, goimports, gocyclo
	( \
			 gofmt -w  -s .; \
			 test -z "$$(go vet ./...)"; \
			 go install golang.org/x/tools/cmd/goimports@latest; \
			 goimports -w .; \
			 go install github.com/fzipp/gocyclo/cmd/gocyclo@latest; \
			 test -z "$$(gocyclo -over 25 .)"; \
			 go install honnef.co/go/tools/cmd/staticcheck@latest ; \
			 staticcheck ./... ; \
    )

lint:  ##  run golint
	( \
			 go install golang.org/x/lint/golint@latest; \
			 golint ./...; \
			 test -z "$$(golint ./...)"; \
    )

bump: git-status clean-venv  ## bump version in main branch
ifeq ($(CURRENT_BRANCH), $(DEFAULT_BRANCH))
	( \
	   source .venv/bin/activate; \
	   pip install bump2version; \
	   bump2version $(part); \
	)
else
	@echo "UNABLE TO BUMP - not on Main branch"
	$(info Current Branch: $(CURRENT_BRANCH), main: $(DEFAULT_BRANCH))
endif


git-status: ## require status is clean so we can use undo_edits to put things back
	@status=$$(git status --porcelain); \
	if [ ! -z "$${status}" ]; \
	then \
		echo "Error - working directory is dirty. Commit those changes!"; \
		exit 1; \
	fi

.PHONY: build static test	