package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	c := make([]byte, 8)
	n, err := rot.r.Read(c)

	for i, v := range c {
		switch {
		case 'A' <= v && v <= 'Z':
			b[i] = 'A' + (v-'A'+13)%26
		case 'a' <= v && v <= 'z':
			b[i] = 'a' + (v-'a'+13)%26
		default:
			b[i] = v
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
