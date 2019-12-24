package randomXorReaderWriter

import (
	"fmt"
	"io"
)

type FromXor struct {
	r1 io.Reader
	r2 io.Reader
}

// NewFromXor reads from two xor sources and combines them to one
func NewFromXor(r1, r2 io.Reader) *FromXor {
	return &FromXor{
		r1: r1,
		r2: r2,
	}
}

func (x FromXor) Read(p []byte) (readbytes int, err error) {
	l := len(p)

	x1buf := make([]byte, l)
	x2buf := make([]byte, l)

	x1readbytes, err := x.r1.Read(x1buf)
	if err != nil {
		return -1, fmt.Errorf(`couldn't read from r1: %w`, err)
	}

	x2readbytes, err := x.r2.Read(x2buf)
	if err != nil {
		return -2, fmt.Errorf(`couldn't read from r2: %w`, err)
	}

	if x1readbytes != x2readbytes {
		return -3, fmt.Errorf(`err: read %d bytes from r1 and %d bytes from r2`, x1readbytes, x2readbytes)
	}

	for i := 0; i < x1readbytes; i++ {
		p[i] = x1buf[i] ^ x2buf[i]
	}

	return x1readbytes, nil
}
