#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
go get -t github.com/go-martini/martini
go get -t github.com/fzzy/radix/redis
go get -t gopkg.in/yaml.v1
go build -o ./bin/unicore-network -v github.com/praekelt/unicore-network
