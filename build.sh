#!/bin/sh

IMAGE_NAME=hyperworks/hello
if [ "$1" != "" ]; then
  IMAGE_NAME="$1"
fi

docker build -t hellobuild .
docker run -v /var/run/docker.sock:/var/run/docker.sock -it --rm \
  -e IMAGE_NAME=$IMAGE_NAME \
  hellobuild
docker rmi hellobuild

