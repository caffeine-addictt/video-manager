BINARY_NAME:=main.out


# =================================== DEFAULT =================================== #

default: all

## default: Runs build and test
.PHONY: default
all: build test

# =================================== HELPERS =================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# =================================== DEVELOPMENT =================================== #

## build: builds the binary
.PHONY: build
build: |
	go build -o $(BINARY_NAME) main.go

## test: Test the program
.PHONY: test
test: |
	go test -v main.go

# =================================== QUALITY ================================== #

## tidy: Tidy mod file
.PHONY: tidy
tidy: |
	go mod tidy -v

## clean: Clean binaries
.PHONY: clean
clean: |
	go clean && rm ${BINARY_NAME}

## fmt: Format code
.PHONY: fmt
fmt:
	go fmt ./...
