#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src
rm -rf "$CURDIR"/pkg "$CURDIR"/bin
go install test
export GOPATH="$OLDGOPATH"
echo 'ok!'