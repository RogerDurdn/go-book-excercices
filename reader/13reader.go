package main

import (
	"fmt"
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rd.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return
}

func find13(base byte) byte {
	return 1
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	btes := make([]byte, 8)
	for {
		numb, err := r.Read(btes)
		fmt.Printf("%v, %v, %v\n", numb, err, btes)
		fmt.Printf("b[:n] = %q\n", btes[:numb])
		if err == io.EOF {
			break
		}
	}
	// io.Copy(os.Stdout, &r)
}
