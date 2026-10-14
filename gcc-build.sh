#!/usr/bin/env -S bash

GOPATH=${GOPATH}:$(readlink -f .):$(readlink -f "${HOME}"/go)
export GOPATH

cd src/main || exit
go build -compiler gccgo -gccgoflags "-march=native -O3" -o ../../converter-gcc
cd - || exit
