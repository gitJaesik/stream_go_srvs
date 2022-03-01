#!/bin/bash

# error handle function
error_handle () {
    if [ $1 -ne 0 ]; then
        echo "$2 failed"
        echo $1
        exit 0
    fi
}

echo 'rm -rf vendor'
rm -rf vendor

# cleansing module
echo 'go mod tidy'
go mod tidy

# wire re install, wire removed in above command
echo 'go get github.com/google/wire/cmd/wire: todo tidy thing make it better'
go get github.com/google/wire/cmd/wire

# wire generate
WIRE_GEN="wire_gen.go"
if [ -f "$WIRE_GEN" ]; then
    echo 'go generate'
    go generate
    error_handle $? 'go generate'
else
    echo "wire"
    wire
    error_handle $? 'wire'
fi

echo 'go build'
go build -o ./stream-auth-server-go
error_handle $? 'go build'

echo 'go mod vendor'
go mod vendor
error_handle $? 'go mod vendor'

# export STREAMGOLIB_LOG_USE_FILELOG=y
# export STREAMGOLIB_LOG_FILENAME=/log/app-%Y-%m-%d-%H.log
# export STREAMGOLIB_LOG_FILENAME_LINK=/log/app.log


echo 'GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info ./stream-auth-server-go'
GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info ./stream-auth-server-go
