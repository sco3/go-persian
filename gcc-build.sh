#!/usr/bin/env -S bash

current_path=$(readlink -f .)
go_home_path=$(readlink -f "${HOME}/go")
GOPATH="${GOPATH}:${current_path}:${go_home_path}"
export GOPATH

cd src/main || exit
go build -compiler gccgo -gccgoflags "-march=native -O3" -o ../../converter-gcc
cd - || exit
