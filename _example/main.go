package main

import (
	"github.com/raspi/randomXorReaderWriter"
	"io"
	"log"
	"os"
)

// Create two xor files from one file
func plainTo2Xor(inFileName, out1Filename, out2Filename string) {
	f, err := os.Open(inFileName)
	if err != nil {
		panic(err)
	}

	outf1, err := os.Create(out1Filename)
	if err != nil {
		panic(err)
	}

	outf2, err := os.Create(out2Filename)
	if err != nil {
		panic(err)
	}

	x := randomXorReaderWriter.NewToXor(f)

	for {
		x1b := make([]byte, 1024)
		x2b := make([]byte, 1024)
		rb, err := x.Read(x1b, x2b)
		if err != nil {
			if err == io.EOF {
				return
			}

			panic(err)
		}
		log.Printf(`got %d bytes`, rb)

		outf1.Write(x1b[0:rb])
		outf2.Write(x2b[0:rb])

		for i := 0; i < rb; i++ {
			log.Printf(`%c`, x1b[i]^x2b[i])
		}

	}

}

// Combine two xor files to one file
func xorToPlain(in1Filename, in2Filename, outFilename string) {
	src1, err := os.Open(in1Filename)
	if err != nil {
		panic(err)
	}

	src2, err := os.Open(in2Filename)
	if err != nil {
		panic(err)
	}

	targetf, err := os.Create(outFilename)
	if err != nil {
		panic(err)
	}

	x := randomXorReaderWriter.NewFromXor(src1, src2)

	for {
		buf := make([]byte, 1024)
		rb, err := x.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}

			panic(err)
		}
		log.Printf(`got %d bytes`, rb)

		log.Printf(`%v`, string(buf[0:rb]))
		targetf.Write(buf[0:rb])
	}
}

func main() {
	plainTo2Xor(`hello.txt`, `hello.txt.x1`, `hello.txt.x2`)
	xorToPlain(`hello.txt.x1`, `hello.txt.x2`, `hello.new.txt`)
}
