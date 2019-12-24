package randomXorReaderWriter

import (
	"crypto/rand"
	"fmt"
	"io"
)

type ToXor struct {
	r io.Reader
}

// NewToXor reads from one source file and outputs two randomly xor'ed bytes
func NewToXor(r io.Reader) *ToXor {
	return &ToXor{
		r: r,
	}
}

func (x ToXor) Read(xbuf1, xbuf2 []byte) (readbytes int, err error) {
	l := len(xbuf1)

	if l != len(xbuf2) {
		return -1, fmt.Errorf(`buffer len mismatch`)
	}

	randBytes := make([]byte, l)
	readbytes, err = rand.Read(randBytes)
	if err != nil {
		return -2, fmt.Errorf(`couldn't read random bytes: %w`, err)
	}

	tmp := make([]byte, l)

	readbytes, err = x.r.Read(tmp)
	if err != nil {
		return readbytes, err
	}

	for i := 0; i < readbytes; i++ {
		xbuf1[i] = tmp[i] ^ randBytes[i]
		xbuf2[i] = randBytes[i]
	}

	return readbytes, nil
}
