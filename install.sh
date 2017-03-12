#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src
rm -rf "$CURDIR"/pkg "$CURDIR"/bin
go install test
"$CURDIR"/bin/test
cd src/top ; go test ; cd -
cd src/topNew ; go test ; cd -
# cd src/topNew ; go test ; go test -bench='.' ; cd -
export GOPATH="$OLDGOPATH"
echo 'ok!'