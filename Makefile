.PHONY: default
default: test

.PHONY: help
## help: print the help info
help:
	@echo -e "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test
## test: test code base using go test
test:
	@which gotest > /dev/null 2>&1 && \
		gotest -v ./... || go test -v ./...
	
.PHONY: lint
## lint: lint code
lint:
	@golangci-lint run ./...
