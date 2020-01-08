#!/bin/bash

go get -d ./...

if [ $? != 0 ]; then
    echo "could not get dependencies"
    exit 1
fi

go test ./...

if [ $? != 0 ]; then
echo "tests failed"
    exit 1
fi

go build add.go

if [ $? != 0 ]; then
    echo "could not build app"
    exit 1
fi
