#!/usr/bin/env bash

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

source ../var.sh

protoc -I=. -I=$PROTOC_INCLUDE --go_out=paths=source_relative:. `find . -name "actor.proto"`

protoc -I=. -I=$PROTOC_INCLUDE --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. `find remote -name "remote.proto"`