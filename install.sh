#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src
rm -rf "$CURDIR"/pkg "$CURDIR"/bin
go install test
cd src/topNew ; go test ; cd -
# cd src/topNew ; go test ; go test -bench='.' ; cd -
"$CURDIR"/bin/test
export GOPATH="$OLDGOPATH"
echo 'ok!'