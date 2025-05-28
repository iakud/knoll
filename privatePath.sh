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

echo "`pwd`/local/bin:$PATH"
#echo "`pwd`/local/bin:`pwd`/local/protoc/bin:$PATH"