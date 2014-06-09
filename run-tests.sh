#!/bin/bash
set -e
go build -o ./bin/unicore-network -v github.com/praekelt/unicore-network
go test -v github.com/praekelt/unicore-network/...
