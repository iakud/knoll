#!/usr/bin/env bash

##################################################
# Owned by knoll. DON'T change me.
##################################################

[[ "$TRACE" ]] && set -x
pushd `dirname "$0"` > /dev/null
trap `popd > /dev/null` EXIT

colorful=false
tput setaf 7 > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    colorful=true
fi

function printError() {
    $colorful && tput setaf 1
    >&2 echo "Error: $@"
    $colorful && tput setaf 7
}

function printImportantMessage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

function printUsage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

#export http_proxy=http://127.0.0.1:1087
#export https_proxy=http://127.0.0.1:1087

INSTALL_DIR="../local"
PROTOC_GEN_GO_VERSION=1.36.5
PROTOC_GEN_GO_GRPC_VERSION=1.5.1

hash protoc-gen-go 2>/dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@v$PROTOC_GEN_GO_VERSION
hash protoc-gen-go-grpc 2>/dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v$PROTOC_GEN_GO_GRPC_VERSION