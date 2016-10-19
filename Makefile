.PHONY: install build test
BUMP_VERSION := $(shell command -v bump_version)
GODOCDOC := $(shell command -v godocdoc)

install:
	go get ./...
	go install ./...

build:
	go build ./...

lint:
	go vet ./...

test:
	go test ./...

race-test:
	go test -race -v ./...

release: race-test
ifndef BUMP_VERSION
	go get github.com/Shyp/bump_version
endif
	bump_version minor types.go

docs:
ifndef GODOCDOC
	go get -u github.com/kevinburke/godocdoc
endif
	godocdoc
