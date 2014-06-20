#!/bin/bash
set -e
./get-dependencies.sh
go build -o ./bin/unicore-network -v github.com/praekelt/unicore-network
redis-cli -n 0 flushdb
cd server && go test -v
cd ..
