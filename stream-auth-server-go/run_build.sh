#!/bin/bash

# error handle function
error_handle () {
    if [ $1 -ne 0 ]; then
        echo "$2 failed"
        echo $1
        exit 0
    fi
}


# export GOOS=linux
# export GOARCH=amd64
# export NAMESPACE=dice

# echo 'go mod tidy'
# go mod tidy
# go get github.com/google/wire/cmd/wire
# go generate

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./dist/dice-api-server/dice-api-server .
error_handle $? 'go build'
echo 'Build complete'
