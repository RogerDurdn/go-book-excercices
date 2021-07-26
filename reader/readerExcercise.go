package main

import (
	"fmt"
	"io"
	"strings"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader := strings.NewReader("hello, reader!")
	btes := make([]byte, 8)

	for {
		numb, err := reader.Read(btes)
		fmt.Printf("%v, %v, %v\n", numb, err, btes)
		fmt.Printf("b[:n] = %q\n", btes[:numb])
		if err == io.EOF {
			break
		}

	}

}
