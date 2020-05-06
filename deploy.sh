#!/bin/bash

APP_NAME=arris-exporter

RUN_FLAG="-d --restart=always"
if [ "$1" == "debug" ]; then
    RUN_FLAG="--rm"
fi

echo "Building $APP_NAME image"
docker build --no-cache -t $APP_NAME .

echo "Removing $APP_NAME container if it exists"
docker rm -f $APP_NAME

echo "Running $APP_NAME container"
docker run $RUN_FLAG -p 9300:9300 --name $APP_NAME $APP_NAME