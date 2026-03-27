.PHONY: all

openjdk:
	script/brewInstall.sh openjdk@11

tool:
	GOBIN=$(shell pwd)/local/bin go install tool
	script/antlr4.sh
	script/protoc.sh

generate:
	chmod +x actor/build.sh
	chmod +x actor/remote/build.sh
	chmod +x kdsync/kdsgen/build.sh
	chmod +x krpc/knetpb/build.sh
	chmod +x krpc/krpcgen/build.sh
	go generate ./...

example:
	chmod +x actor/examples/chat/message/build.sh
	chmod +x kdsync/example/build.sh
	chmod +x nrpc/examples/hello/hello/build.sh

	actor/examples/chat/message/build.sh
	kdsync/example/build.sh
	nrpc/examples/hello/hello/build.sh

kdsgen:
	GOBIN=$(shell pwd)/local/bin go install ./kdsync/kdsgen

protoc-gen-nrpc:
	GOBIN=$(shell pwd)/local/bin go install ./nrpc/protoc-gen-nrpc