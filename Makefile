.PHONY: build
build:
	go build -v ./server/cmd/apiserver

.DEFAULT_GOAL := build