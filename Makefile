.PHONY: all

kdsc:
	GOBIN=$(shell pwd)/local/bin go install ./kds/kdsc

tool:
	GOBIN=$(shell pwd)/local/bin go install tool
	script/antlr4.sh
	script/protoc.sh

openjdk:
	./script/brewInstall.sh openjdk@11

kdspb:
	TRACE=1 kds/kdspb/conv.sh

kdscexample:
	TRACE=1 kds/kdsc/example/conv.sh

generate:
	go generate ./...