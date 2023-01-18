package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rotr.r.Read(b)
	for i, v := range b {
		if v >= 65 && v <= 90 {
			b[i] = 65 + ((b[i]-65)+13)%26
		} else if v >= 97 && v <= 122 {
			b[i] = 97 + ((b[i]-97)+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
