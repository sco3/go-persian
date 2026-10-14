#!/usr/bin/env -S bash

#javah -d ../c/include/ -classpath "../../../target/test-classes/:../../../target/dependency/*" com.tnf.avro.TestAvro

go build -buildmode=c-shared -o libpersian-go.so

ret=$?

if [[ ${ret} -ne 0 ]]; then
	exit
fi

strip libpersian-go.so

stat libpersian-go.so

nm -D libpersian-go.so | grep -i java || true
