#!/bin/bash
cd "${WORKSPACE}/${REPO}"
export GOPATH=`pwd`
go get github.com/go-martini/martini
go get github.com/fzzy/radix/redis
go get gopkg.in/yaml.v1
go get github.com/praekelt/unicore-network
go build -o ./bin/unicore-network -v github.com/praekelt/unicore-network
