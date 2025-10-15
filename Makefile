.PHONY: all

openjdk:
	script/brewInstall.sh openjdk@11

tool:
	GOBIN=$(shell pwd)/local/bin go install tool
	script/antlr4.sh
	script/protoc.sh

generate:
	go generate ./...

kdsc:
	GOBIN=$(shell pwd)/local/bin go install ./kds/kdsc

nrpc:
	GOBIN=$(shell pwd)/local/bin go install ./nrpc/protoc-gen-go-nrpc