# Simple Makefile for Go project

# Name of the binary to create
APP_NAME := app

# Default target: build compiles the Go code into a binary
# It includes code in the current directory (main.go) and sub-packages (like cmd/)
build:
	go build -o out/$(APP_NAME) .

# run: compiles (if needed) and runs the binary
run: build
	./out/$(APP_NAME)

# test: runs all unit tests
test:
	go test ./...

# lint: checks code for common issues
lint:
	go vet ./...
	# Uncomment the next line if golangci-lint is installed
	# golangci-lint run

# clean: remove compiled binaries
clean:
	rm -rf out/


# Mark targets that are not real files as phony
.PHONY: build run test lint clean
