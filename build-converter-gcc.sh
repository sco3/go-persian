#!/usr/bin/env -S bash

export GOPATH="${HOME}/go"

go build -compiler gccgo -gccgoflags "-march=native -O3" -o converter-gcc
