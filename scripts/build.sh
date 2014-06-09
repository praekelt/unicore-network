#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
go get ./...
go build -o ./bin/unicore-network -v github.com/praekelt/unicore-network
