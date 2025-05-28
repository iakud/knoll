.PHONY: all

kdsc:
	GOBIN=$(shell pwd)/local/bin go install ./kds/kdsc

protoc:
	#script/protoc.sh
	./script/brewInstall.sh protobuf
	script/protoc-gen-go.sh

openjdk:
	./script/brewInstall.sh openjdk@11

kdspb:
	TRACE=1 kds/kdspb/conv.sh

kdscexample:
	TRACE=1 kds/kdsc/example/conv.sh

generate:
	go generate ./...