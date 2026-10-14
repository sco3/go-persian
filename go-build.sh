#!/usr/bin/env -S bash

GOPATH=${GOPATH}:$(readlink -f .):$(readlink -f "${HOME}"/go)
export GOPATH

#-c -v -gcflags "-N -l"

cd src/main || exit
go build -o ../../converter
cd - || exit
