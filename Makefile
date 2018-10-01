.PHONY: install test prep deps

#--- Help ---
help:
	@echo 
	@echo Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo 

#--- Setup targets ---
prep: ## Make preparations to run the tests
	mkdir -p test
deps: ## Get all the dependencies
	go get -u github.com/spf13/cobra
	go get -u github.com/stretchr/testify/assert
	go get -u github.com/mattn/go-sqlite3
	# TODO: This needs to be changed once Travis supports Go mods too
	go get -u github.com/google/go-github/...

#--- Test targets ---
test: ## Run all testcases
	export TESTDIR=`pwd`/test && go test ./...

#--- Run targets ---
install: ## Install the executable in your $GOPATH/bin folder
	go install