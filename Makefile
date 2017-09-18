# VARIABLES
PACKAGE="github.com/bretanac93/logindex"
BINARY_NAME="logindex"

# export GOPATH=$(shell pwd)

default: help

# Clean .o files and binary
clean: ## Remove binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@rm -rf bin/$(BINARY_NAME) 2> /dev/null
	@echo "Clean OK"

# Compile sources and build binary
install: clean ## Compile sources and build binary
	@echo "--> installing..."
	@go install $(PACKAGE) || (echo "Compilation error" && exit 1)
	@echo "Install OK"

# Run your application
run: clean install ## Build application and run it
	@echo "--> running application..."
	@./bin/$(BINARY_NAME)

# Test your application
test: ## Run all tests of your app
	@echo "--> testing..."
	@go test -v $(PACKAGE)/...

# Help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
