#!/usr/bin/env bash

set -e -u -o pipefail

cd transpile

# the generated code depends on the size of pointers, translate for 64-bit and 32-bit targets
go run . "$@"
go run . -goarch arm "$@"
