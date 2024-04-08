BINARY_NAME:=video-manager
GOBIN:=${HOME}/go/bin


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
test:
	go mod verify
	go vet ./...
	go run github.com/securego/gosec/v2/cmd/gosec@latest -quiet ./...
	go run github.com/go-critic/go-critic/cmd/gocritic@latest check -enableAll ./...
	go run github.com/google/osv-scanner/cmd/osv-scanner@latest -r .
	go test -race .

# =================================== QUALITY ================================== #

## tidy: Tidy mod file and format code
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## clean: Clean binaries
.PHONY: clean
clean: |
	go clean && rm ${BINARY_NAME}


# https://golangci-lint.run/welcome/install/
## lint: Lint code
.PHONY: lint
lint:
	$(GOBIN)/golangci-lint run -v
