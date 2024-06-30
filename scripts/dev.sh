#!/bin/bash

set -e

cd "$(dirname "$0")/.."

# Compile Javascript
cd gogui/js/src
go run .
cd ../../../

# Compile CSS

# Run Basic Example
cd examples/basic
go run .
