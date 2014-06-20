#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
./get-dependencies.sh
go build -o ./bin/unicore-network -v github.com/praekelt/unicore-network
cd ..
cp -a ${REPO} ./build/