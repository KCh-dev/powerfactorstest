#!/bin/bash

# Set default number of connections to 3
num_connections=3

# Check if an argument is provided
if [ "$#" -eq 1 ]; then
    num_connections="$1"
elif [ "$#" -gt 1 ]; then
    echo "Usage: $0 [<number of connections>]"
    exit 1
fi

../bin/client -n "$num_connections"