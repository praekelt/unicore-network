#!/bin/bash
export GOPATH="${WORKSPACE}/${REPO}"
go get ./...
go build -o ./bin/unicore_network -v github.com/praekelt/unicore_network
