#!/bin/sh

docker run --rm -v "$PWD":/go/src/hello -w /go/src/hello \
  -e CGO_ENABLED=0 -e GOOS=linux golang:1.4 \
  go build -v -a -installsuffix cgo -o hello .

