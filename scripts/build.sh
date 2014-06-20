#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
./get-dependencies.sh
go build -o ${BUILDDIR}/unicore-network -v github.com/praekelt/unicore-network

