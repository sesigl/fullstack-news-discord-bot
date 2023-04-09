.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint:
	golangci-lint run ./...
.PHONY:lint

test:
	go test ./...
.PHONY:test

install-deps:
	go get ./...
.PHONY: install-deps

build: lint test
	go build src/main.go
.PHONY:build

run: lint
	go run src/main.go
.PHONY:build