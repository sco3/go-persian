#!/usr/bin/env -S bash

go build -compiler gccgo -gccgoflags "-march=native -O3" -o converter-gcc
