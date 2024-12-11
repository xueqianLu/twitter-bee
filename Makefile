.PHONY: default swag tbee docker

GOBIN = $(shell pwd)/build/bin
GO ?= latest
GOFILES_NOVENDOR := $(shell go list -f "{{.Dir}}" ./...)
TAG=latest

default: tbee

all: tbee docker

swag:
	swag init -g openapi/server.go

tbee:
	go build -o=${GOBIN}/$@ -gcflags "all=-N -l" .
	@echo "Done building."

docker:
	docker build -t tbee:${TAG} .

clean:
	rm -fr build/*
