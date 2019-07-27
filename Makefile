# Go parameters
PROJECTNAME=$(shell basename "$(PWD)")
ENTRYFILE="cmd/$(PROJECTNAME)/$(PROJECTNAME).go"

GOCMD=go

MAKEFLAGS += --silent

## run: Runs your application
run:
	@echo " > Running..."
	$(GOCMD) run $(ENTRYFILE)

## install: Installs your dependencies
install:
	@echo " > Installing dependencies in vendor/"
	$(GOCMD) mod vendor
	@echo " > Done."

## build: Builds your application
build:
	@echo " > Building..."
	$(GOCMD) build -a -ldflags '-extldflags "-static"' $(ENTRYFILE)
	@echo " > Done."

## test: Runs your tests if any
test:
	@echo " > Running tests..."
	$(GOCMD) test ./... -cover -v -coverprofile=coverage.out
	@echo " > Done."

## cover: Checks your code coverage
cover:
	@echo " > Checking coverage..."
	$(GOCMD) tool cover -html=coverage.out

.PHONY: help
all: help
help: Makefile
	@echo
	@echo "Choose a command to run in "$(PROJECTNAME)":"
	@echo
	sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo