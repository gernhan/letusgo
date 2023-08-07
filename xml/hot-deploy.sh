#!/bin/bash

set -e
set -x

GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./app .
docker buildx build --tag repository.int.compax.at:5001/tm-aax-invoice-interface:perftest --platform=linux/amd64 -o type=image .
docker push repository.int.compax.at:5001/tm-aax-invoice-interface:perftest 
