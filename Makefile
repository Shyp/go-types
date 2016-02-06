.PHONY: install build test

install:
	go install ./...

build:
	go build ./...

test:
	go test ./...
