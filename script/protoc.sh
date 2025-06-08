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

PROTOC_VERSION=30.2

if [[ `$PROTOC_BIN/protoc --version 2>&1 | grep -e "libprotoc $PROTOC_VERSION"` ]]; then
    echo -e "[misc] \033[0;33mprotoc $PROTOC_VERSION\033[0;37m is already installed."
    exit 0
fi

case $(uname -m) in # 检测系统架构
    x86_64) ARCH="x86_64" ;;
    aarch64) ARCH="aarch_64" ;;
    arm64) ARCH="aarch_64" ;;   # Mac M1/M2
    *) ARCH="x86_64" ;;         # 默认x86_64
esac

case $(uname -s) in # 检测操作系统
    Linux) OS="linux" ;;
    Darwin) OS="osx" ;;
    *) echo "Unsupported OS"; exit 1 ;;
esac

PROTOC_FILE="protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip"

echo -e "[misc] Start to install \033[0;33mprotoc $PROTOC_VERSION\033[0;37m..."
wget -c "https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_FILE"
[[ $? -ne 0 ]] && exit 1

rm -rf $PROTOC_PATH
unzip $PROTOC_FILE -d $PROTOC_PATH
[[ $? -ne 0 ]] && exit 1

rm -rf $PROTOC_FILE

echo -e "[misc] Start to install \033[0;33mprotoc-gen-go\033[0;37m..."
rm -rf $LOCAL_BIN/protoc-gen-go
GOBIN=$LOCAL_BIN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

echo -e "[misc] Start to install \033[0;33mprotoc-gen-go-grpc\033[0;37m..."
rm -rf $LOCAL_BIN/protoc-gen-go-grpc
GOBIN=$LOCAL_BIN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest