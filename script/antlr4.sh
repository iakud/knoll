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


export http_proxy=http://127.0.0.1:1087
export https_proxy=http://127.0.0.1:1087

INSTALL_DIR="../local"
ANTLR4_VERSION=4.13.2

source ../var.sh

if [[ `java -jar "$INSTALL_DIR/lib/antlr-$ANTLR4_VERSION-complete.jar" 2>&1 | grep -e "ANTLR Parser Generator  Version $ANTLR4_VERSION"` ]]; then
    echo -e "[misc] \033[0;antlr4 $ANTLR4_VERSION\033[0;37m is already installed."
	exit 0
fi

rm -rf "$INSTALL_DIR/lib/antlr-$ANTLR4_VERSION-complete.jar"
mkdir -p "$INSTALL_DIR/lib"

echo -e "[misc] Start to install \033[0;33mantlr4 $ANTLR4_VERSION\033[0;37m..."

wget -c "https://www.antlr.org/download/antlr-$ANTLR4_VERSION-complete.jar"
[[ $? -ne 0 ]] && exit 1

mv "antlr-$ANTLR4_VERSION-complete.jar" "$INSTALL_DIR/lib"