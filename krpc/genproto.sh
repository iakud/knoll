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

source ../var.sh

PROTO_PATH=$ROOT_PATH/krpc/knetpb
PB_PATH=$ROOT_PATH/krpc/knetpb

protoc -I=$PROTO_PATH \
    -I=$PROTOC_INCLUDE \
    --go_opt=default_api_level=API_OPAQUE \
	--go_out=paths=source_relative:$PB_PATH \
	`find $PROTO_PATH -name "*.proto"`