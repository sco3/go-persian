#!/usr/bin/env -S bash

go build -buildmode=c-shared -o libpersian-go.so

ret=$?

if [[ ${ret} -ne 0 ]]; then
	exit
fi

strip libpersian-go.so

stat libpersian-go.so

nm -D libpersian-go.so | grep -i java || true
