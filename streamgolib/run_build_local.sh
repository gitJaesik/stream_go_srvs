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

echo 'go mod tidy'
go mod tidy

echo 'go get github.com/google/wire/cmd/wire: todo tidy thing make it better'
go get github.com/google/wire/cmd/wire

echo 'go generate'
go generate
error_handle $? 'go generate'

echo 'go build'
go build
error_handle $? 'go build'

echo 'go mod vendor'
go mod vendor
error_handle $? 'go mod vendor'

echo 'go build local complete'
