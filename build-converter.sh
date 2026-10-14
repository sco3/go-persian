#!/usr/bin/env -S bash

export GOPATH="${HOME}/go"

#-c -v -gcflags "-N -l"

go build -o converter
