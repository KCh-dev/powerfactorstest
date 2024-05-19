#!/bin/bash
# Script to run tests and benchmarks for the strgen package

PKG_PATH="../pkg/util"

echo "Running tests..."
go test -v $PKG_PATH

if [ $? -ne 0 ]; then
    echo "Tests failed."
    exit 1
fi

echo "Running benchmarks..."
go test -bench=. $PKG_PATH