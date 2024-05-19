#!/bin/bash
# Script to run tests and benchmarks for the strgen package

echo "Running tests..."
go test -v ../internal/pkg/strgen/strgen_test.go

if [ $? -ne 0 ]; then
    echo "Tests failed."
    exit 1
fi

echo "Running benchmarks..."
go test -bench=. ./path/to/strgen

