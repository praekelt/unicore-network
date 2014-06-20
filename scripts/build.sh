#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
./get-dependencies.sh
go build -o ${BUILD}/unicore-network -v github.com/praekelt/unicore-network

