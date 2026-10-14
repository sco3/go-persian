package persian

import (
	"encoding/hex"
	"fmt"
	"unsafe"

	"testing"

	"log"

	"bytes"
	"github.com/go-avro/avro"
)

func TestSimple(t *testing.T) {
	var b = make([]byte, 20)
	var buffer = NewMemBufferPtr(uintptr(unsafe.Pointer(&b[0])), len(b))

	encoder := avro.NewBinaryEncoder(buffer)
	encoder.WriteLong(0)
	encoder.WriteLong(1)
	encoder.WriteLong(3)
	fmt.Println(hex.Dump(buffer.Buf))
}

func TestIterate(t *testing.T) {
	var in = []byte("|asdf1|asdf2|asdf3|")
	log.Printf("%s\n", string(in))
	var end int = 0
	var field []byte
	var rest []byte = in
	for i := 0; i < 6; i++ {
		if end < 0 {
			field = nil
		} else {
			rest = rest[end:]
			end = bytes.IndexByte(rest, '|')
			if end < 0 {
				field = rest
			} else {
				field = rest[:end]
				end++
			}
		}
		log.Printf("field %s\n", string(field))
	}

}
