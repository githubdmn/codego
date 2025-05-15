# /home/dmn/projects/codego/Makefile

.PHONY: build run test clean

BIN_DIR := bin
BIN     := $(BIN_DIR)/api

build:
	mkdir -p ../bin
	go build -o ../bin/api ./api

run:
	go run ./api


test:
	go test ./...

clean:
	rm -rf $(BIN_DIR)

.DEFAULT_GOAL := run