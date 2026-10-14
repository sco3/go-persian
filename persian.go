package main

//#cgo CFLAGS: -I./include
//#cgo CFLAGS: -I./include/linux
//#include <jni.h>
import "C"

import (
	"os"
	"unsafe"
	"persian/v2/persian"
	"github.com/pkg/profile"
)

//export Java_sco3_kafka_Persian_persePtr
func Java_sco3_kafka_Persian_persePtr(
	env *C.JNIEnv, clazz C.jclass,
	schema C.jlong, schemaLen C.jint,
	in C.jlong, inlen C.jint,
	out C.jlong, offset C.jint, outlen C.jint) int {

	schemaBuf := persian.SliceFromPtr(uintptr(schema), int(schemaLen))
	inbuf := persian.SliceFromPtr(uintptr(in), int(inlen))
	outbuf := persian.NewMemBufferPtr(uintptr(out), int(outlen))
	outbuf.Offset = int(offset)

	p := new(persian.FldInfo)
	return persian.Convert(schemaBuf, inbuf, outbuf, p)
}

/*
	println(p.base)

	println(p.tsIdx)
	println(p.dtIdx)
	println(p.sgmIdx)
	println(p.imsiIdx)
	println(p.keyIdx)

	println(p.tsOffset)
	println(p.tsLen)

	println(p.imsiOffset)
	println(p.imsiLen)

	println(p.keyOffset)
	println(p.keyLen)

	println(p.sgm)
	println(p.ts)

	bufio.NewWriter(os.Stdout).Flush()

	p.base = 15
	p.tsIdx = 16
	p.dtIdx = 17
	p.sgmIdx = 18
	p.imsiIdx = 19
	p.keyIdx = 20
	p.tsOffset = 21
	p.tsLen = 22
	p.imsiOffset = 23
	p.imsiLen = 24
	p.keyOffset = 25
	p.keyLen = 26
	p.sgm = 27
	p.ts = 28

*/

//export Java_sco3_kafka_Persian_perseCalcPtrExt
func Java_sco3_kafka_Persian_perseCalcPtrExt(
	env *C.JNIEnv, clazz C.jclass,
	schema C.jlong, schemaLen C.jint,
	in C.jlong, inlen C.jint,
	out C.jlong, offset C.jint, outlen C.jint,
	jfldinfo C.jlong) int {

	schemaBuf := persian.SliceFromPtr(uintptr(schema), int(schemaLen))
	inbuf := persian.SliceFromPtr(uintptr(in), int(inlen))
	outbuf := persian.NewMemBufferPtr(uintptr(out), int(outlen))
	outbuf.Offset = int(offset)

	var p *persian.FldInfo = (*persian.FldInfo)(unsafe.Pointer(uintptr(jfldinfo)))

	return persian.Convert(schemaBuf, inbuf, outbuf, p)
}

func main() {
	if len(os.Args) == 1 {
		println("profiling")
		defer profile.Start().Stop()
	} else {
		println("no profiling")
	}

	persian.RunConverter()
}
