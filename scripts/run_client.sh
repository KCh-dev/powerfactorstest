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

# Navigate to the root directory of the project
cd "$(dirname "$0")/.." || exit

# Run the client with the specified or default number of connections
./bin/client -n "$num_connections"