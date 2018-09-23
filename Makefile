.PHONY: install test prep deps

#--- Help ---
help:
	@echo Makefile for gh
	@echo  
	@echo usage: make [target]
	@echo 
	@echo Test targets:
	@echo - test: Runs all test cases
	@echo 
	@echo Setup targets:
	@echo - prep: Make preparations to run tests
	@echo - deps: Get all dependencies
	@echo
	@echo Run targets
	@echo - install : Installs the executable in your GOPATH/bin	

#--- Setup targets ---
prep:
	mkdir -p test
deps:
	go get -u github.com/spf13/cobra
	go get -u github.com/stretchr/testify/assert

#--- Test targets ---
test:
	export TESTDIR=`pwd`/test && go test ./...

#--- Run targets ---
install:
	go install