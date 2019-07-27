# Go parameters
GOCMD=go

.SILENT:
install:
		$(GOCMD) mod vendor

local:
		$(GOCMD) run cmd/apiname/apiname.go

build:
		$(GOCMD) build cmd/apiname/apiname.go

test:
		$(GOCMD) test ./... -cover -v -coverprofile=coverage.out

cover:
		$(GOCMD) tool cover -html=coverage.out