# -----------------------------------------------------------------------------
# Description: Makefile
# Author(s): retgits
# Last updated: 2019-04-19
# 
# This software may be modified and distributed under the terms of the
# MIT license. See the LICENSE file for details.
# -----------------------------------------------------------------------------

# Set the shell to bash
SHELL=/usr/bin/env bash

# GOPROXY defines which URL to use to retrieve Go Modules from
GOPROXY=https://gocenter.io

# Sets PWD to pwd_unknown if it doesn't have a value. Normally this should not
# happen. If you do see pwd_unknown showing up, you'll need to make sure your
# system understand the PWD command.
PWD ?= pwd_unknown

# PROJECT_NAME defaults to name of the current directory.
PROJECT_NAME = $(notdir $(PWD))

# VERSION either uses the current commit hash, or will default to "dev"
VERSION := $(strip $(if $(shell git rev-parse HEAD),$(shell git rev-parse HEAD),dev))

# Create a list of all packages in this repository
PACKAGES=$(shell go list ./...)

.PHONY: help
help: ## Displays the help for each target (this message)
	@echo 
	@echo Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo 

.PHONY: lint
lint: ## Lint examines Go source code and prints style mistakes for all packages.
	env GO111MODULE=on golint -set_exit_status $(ALL_PACKAGES)

.PHONY: deps
deps: ## Get all the Go dependencies.
	go get -u ./...

.PHONY: test
test: ## Run all testcases.
	mkdir -p ${PWD}/test
	mkdir -p $(GOPATH)/bin
	go get -u golang.org/x/lint/golint
	go get -u github.com/gojp/goreportcard/cmd/goreportcard-cli
	env TESTDIR=${PWD}/test go test ./...

.PHONY: score
score: ## Get a score based on GoReportcard.
	goreportcard-cli -v

.PHONY: compile
compile: ## Compiles and creates an executable in the 'out' folder.
	mkdir -p out/
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o out/${PROJECT_NAME}-linux-amd64 main.go
	env GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o out/${PROJECT_NAME}-windows-amd64.exe main.go
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o out/${PROJECT_NAME}-darwin-amd64 main.go

.PHONY: install
install: ## Compiles and installs the packages named by the import paths.
	go install