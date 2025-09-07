ROOT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

LOCAL_PATH=$ROOT_PATH/local
LOCAL_BIN=$LOCAL_PATH/bin
LOCAL_LIB=$LOCAL_PATH/lib
LOCAL_INCLUDE=$LOCAL_PATH/include

PROTOC_PATH=$LOCAL_PATH/protoc
PROTOC_BIN=$PROTOC_PATH/bin
PROTOC_INCLUDE=$PROTOC_PATH/include

PATH=$LOCAL_BIN:$PROTOC_BIN:$PATH

if [[ `uname -m | grep -e "arm64"` ]]; then
	PATH="`brew --prefix openjdk@11`/bin:$PATH"
fi