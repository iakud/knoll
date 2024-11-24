PROTOC_VERSION=28.3
GOGOPROTO_VERSION=1.3.2

.PHONY: all

kdsc:
	GOBIN=$(shell pwd)/local/bin go install ./kds/kdsc

protoc:
	script/protoc.sh
	script/protoc-gen-go.sh

kdspb:
	TRACE=1 kds/kdspb/conv.sh

kdscexample:
	TRACE=1 kds/kdsc/example/conv.sh