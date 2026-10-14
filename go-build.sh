#!/usr/bin/env -S bash

current_path=$(readlink -f .)
go_home_path=$(readlink -f "${HOME}/go")
GOPATH="${GOPATH}:${current_path}:${go_home_path}"
export GOPATH

#-c -v -gcflags "-N -l"

cd src/main || exit
go build -o ../../converter
cd - || exit
