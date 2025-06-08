PWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# export PATH="`brew --prefix openjdk@11`/bin:$PATH"

PWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

LOCAL_PATH=$PWD/local
LOCAL_BIN=$LOCAL_PATH/bin
LOCAL_INCLUDE=$LOCAL_PATH/include

PROTOC_PATH=$LOCAL_PATH/protoc
PROTOC_BIN=$PROTOC_PATH/bin
PROTOC_INCLUDE=$PROTOC_PATH/include

PATH=$LOCAL_BIN:$PROTOC_BIN:$PATH