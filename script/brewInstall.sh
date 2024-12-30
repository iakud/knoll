#!/usr/bin/env bash

###======================================================###
### Desc 使用brew工具安装
### Usage ./brewInstall.sh mongodb-community mongod
### Param-1 brew formula name
### Param-2 program name
###======================================================###

if [[ `which brew` == '' ]]; then
    echo "[misc] Start to install brew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    [[ $? -ne 0 ]] && exit 1
fi

Formula=${1}
Program=${2}

[[ $(brew list --formula | grep -e "^${Formula}$") != '' ]] && echo -e "[brew] \033[0;33m${Formula}\033[0;37m is already installed." && exit 0

[[ $(which "${Program}") != '' ]] && echo -e "[brew] \033[0;33m${Program}\033[0;37m detected, skipped." && exit 0

echo "[brew] Start to install ${Formula}..."
brew install ${Formula}
[[ $? -ne 0 ]] && exit 1

exit 0