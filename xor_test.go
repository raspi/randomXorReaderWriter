package randomXorReaderWriter

import (
	"bytes"
	"testing"
)

func TestToXor(t *testing.T) {
	source := []byte(`Hello, world!`)

	toXor := NewToXor(bytes.NewReader(source))

	xorBuffer1 := make([]byte, len(source))
	xorBuffer2 := make([]byte, len(source))

	rBytes, err := toXor.Read(xorBuffer1, xorBuffer2)
	if err != nil {
		t.Fail()
	}

	if rBytes != len(source) {
		t.Fail()
	}

	for i := 0; i < rBytes; i++ {
		if source[i] != xorBuffer1[i]^xorBuffer2[i] {
			t.Fail()
		}
	}

}

func TestFromXor(t *testing.T) {
	source1 := []byte{0x3b, 0xca, 0x19, 0x86, 0x43, 0x15, 0x28, 0x4a, 0x15, 0x12, 0x27, 0xda, 0xee}
	source2 := []byte{0x73, 0xaf, 0x75, 0xea, 0x2c, 0x39, 0x08, 0x3d, 0x7a, 0x60, 0x4b, 0xbe, 0xcf}
	expected := []byte(`Hello, world!`)

	fromXor := NewFromXor(bytes.NewReader(source1), bytes.NewReader(source2))

	actual := make([]byte, len(source1))
	rBytes, err := fromXor.Read(actual)
	if err != nil {
		t.Fail()
	}

	if rBytes != len(expected) {
		t.Fail()
	}

	if string(actual) != string(expected) {
		t.Fail()
	}
}
