#!/usr/bin/env bash

##################################################
# Owned by krocher. DON'T change me.
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


export http_proxy=http://127.0.0.1:1087
export https_proxy=http://127.0.0.1:1087

INSTALL_DIR="../local"
PROTOC_GEN_GO_VERSION=1.36.1

if [[ `$INSTALL_DIR/bin/protoc-gen-go --version 2>&1 | grep -e "protoc-gen-go v$PROTOC_GEN_GO_VERSION"` ]]; then
	echo -e "[misc] \033[0;33mprotoc-gen-go v$PROTOC_GEN_GO_VERSION\033[0;37m is already installed."
	exit 0
fi

rm -rf $INSTALL_DIR"/bin/protoc-gen-go"
mkdir -p $INSTALL_DIR"/bin"

echo -e "[misc] Start to install \033[0;33mprotoc-gen-go $PROTOC_VERSION\033[0;37m..."
wget -c "https://github.com/protocolbuffers/protobuf-go/releases/download/v$PROTOC_GEN_GO_VERSION/protoc-gen-go.v$PROTOC_GEN_GO_VERSION.darwin.arm64.tar.gz"
[[ $? -ne 0 ]] && exit 1

tar -xvf "protoc-gen-go.v$PROTOC_GEN_GO_VERSION.darwin.arm64.tar.gz" -C $INSTALL_DIR"/bin"
[[ $? -ne 0 ]] && exit 1

rm "protoc-gen-go.v$PROTOC_GEN_GO_VERSION.darwin.arm64.tar.gz"