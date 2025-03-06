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
PROTOC_VERSION=30.0

if [[ `$INSTALL_DIR/protoc/bin/protoc --version 2>&1 | grep -e "libprotoc $PROTOC_VERSION"` ]]; then
	echo -e "[misc] \033[0;33mprotoc $PROTOC_VERSION\033[0;37m is already installed."
	exit 0
fi

rm -rf $INSTALL_DIR"/protoc"
mkdir -p $INSTALL_DIR

echo -e "[misc] Start to install \033[0;33mprotoc $PROTOC_VERSION\033[0;37m..."
wget -c	"https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-osx-aarch_64.zip"
[[ $? -ne 0 ]] && exit 1

unzip "protoc-$PROTOC_VERSION-osx-aarch_64.zip" -d $INSTALL_DIR"/protoc"
[[ $? -ne 0 ]] && exit 1

rm "protoc-$PROTOC_VERSION-osx-aarch_64.zip"