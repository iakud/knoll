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

NEW_PATH=`../../../privatePath.sh`
[[ $? -ne 0 ]] && exit 1
PATH="$NEW_PATH"

rm -rf kds
mkdir -p kds

rm -rf proto
mkdir -p proto

kdsc --go_out=kds --proto_out=proto *.kds
[[ $? -ne 0 ]] && exit 1

rm -rf kdspb
mkdir -p kdspb

protoc -I=./proto -I=../../../local/protoc/include \
	--go_out=paths=source_relative:./kdspb \
	`find proto -name "*.proto"`