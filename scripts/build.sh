#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
./get-dependencies.sh
mkdir ${BUILDDIR}/${REPO}
go build -o ${BUILDDIR}/${REPO}/unicore-network -v github.com/praekelt/unicore-network

