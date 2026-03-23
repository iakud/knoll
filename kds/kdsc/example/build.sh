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

source ../../../var.sh

rm -rf kds/*.kds.cs kds/*.kds.go

# kdsc --out=kds --tmpl=../template/kds.go.tmpl kds/*.kds
go run ../ --out=kds --kind=go kds/*.kds
# kdsc --out=kds --tmpl=../template/kds.cs.tmpl kds/*.kds
go run ../ --out=kds --kind=cs kds/*.kds

dotnet publish -c Release -o bin --use-current-runtime .
