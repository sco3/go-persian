package persian

import (
	"errors"
	"unsafe"
)

// MemBuffer is memory writer
type MemBuffer struct {
	Buf    []byte
	Offset int
}

// SliceFromPtr creates slice
func SliceFromPtr(addr uintptr, size int) (t []byte) {
	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(addr), size, size}
	b := *(*[]byte)(unsafe.Pointer(&sl))
	return b
}

// NewMemBufferPtr created from address
func NewMemBufferPtr(addr uintptr, size int) (t *MemBuffer) {
	b := SliceFromPtr(addr, size)
	return &MemBuffer{b, 0}
}

// AllocMemBuffer constructor from string
func AllocMemBuffer(size int) (t *MemBuffer) {
	return &MemBuffer{make([]byte, size), 0}
}

func (t *MemBuffer) Write(src []byte) (n int, err error) {
	n = len(src)
	if t.Offset+n > len(t.Buf) {
		err = errors.New("Not enough space")
	} else {
		for i := 0; i < n; i++ {
			t.Buf[t.Offset+i] = src[i]
		}
		t.Offset += n
	}
	return n, err
}
