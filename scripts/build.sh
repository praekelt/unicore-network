#!/bin/bash
export GOPATH=`pwd`
go get ./...
go build -o ./bin/unicore_network -v github.com/praekelt/unicore_network
