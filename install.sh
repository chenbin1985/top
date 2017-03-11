#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src
go install top
go build test
export GOPATH="$OLDGOPATH"
echo 'ok!'