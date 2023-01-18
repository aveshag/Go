package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// The io package specifies the io.Reader interface, which represents the read end of a stream of data
	// The io.Reader interface has a Read method: func (T) Read(b []byte) (n int, err error)

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
