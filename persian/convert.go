package persian

import (
	"bytes"
	"strconv"

	"github.com/go-avro/avro"
)

// WriteDouble write dbl
func WriteDouble(encoder avro.Encoder, field []byte) {
	d, err := strconv.ParseFloat(string(field), 64)
	if err != nil {
		encoder.WriteInt(0)
	} else {
		encoder.WriteInt(1)
		encoder.WriteDouble(d)
	}
}

// WriteLong write long
func WriteLong(encoder avro.Encoder, field []byte) {
	l, err := strconv.ParseInt(string(field), 10, 64)
	if err != nil {
		encoder.WriteInt(0)
	} else {
		encoder.WriteInt(1)
		encoder.WriteLong(l)
	}

}

// WriteString write string
func WriteString(encoder avro.Encoder, field []byte) {
	encoder.WriteInt(1)
	encoder.WriteBytes(field)
}

// FldInfo exchange data betwen java and native/go
type FldInfo struct {
	Base int32

	TsIdx   int32
	DtIdx   int32
	SgmIdx  int32
	ImsiIdx int32
	KeyIdx  int32

	TsOffset int32
	TsLen    int32

	ImsiOffset int32
	ImsiLen    int32

	KeyOffset int32
	KeyLen    int32

	Sgm int64
	Ts  int64
}

// Convert creates avro from csv
func Convert(schema []byte, in []byte, out *MemBuffer, fld *FldInfo) int {
	encoder := avro.NewBinaryEncoder(out)
	//fields := bytes.Split(in, []byte("|"))
	var end int = 0
	var field []byte
	var rest []byte = in

	for _, t := range schema {

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

		if bytes.Compare(field, []byte("\\N")) == 0 {
			field = nil
		}
		length := len(field)
		if length == 0 { // empty field - > null
			encoder.WriteInt(0)
		} else if t == 's' {
			WriteString(encoder, field)
		} else if t == 'l' {
			WriteLong(encoder, field)
		} else if t == 'd' {
			WriteDouble(encoder, field)
		}
	}
	return out.Offset
}
