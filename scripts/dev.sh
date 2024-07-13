#!/bin/bash

set -e

cd "$(dirname "$0")/.."

# Compile Javascript
cd lib/js/src
go run .
cd ../../../

# Compile CSS

# Run Basic Example
cd examples/"$1"
go run .
