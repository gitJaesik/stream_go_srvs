#!/bin/bash

echo 'go build linux'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./dist/stream-api-server/stream-api-server .
ret=$?
if [ $ret -ne 0 ]; then
  echo $ret
  exit 0
fi
echo 'go build linux complete'
