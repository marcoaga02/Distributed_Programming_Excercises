package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return n, err
	}
	for i:=0; i<len(b); i++ {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func rot13(c byte) byte {
	switch  {
	case 'A' <= c && c <= 'Z':
		if (c + 13) > 90 {
			return (c - 13)
		} else {
			return c + 13
		}
	case 'a' <= c && c <= 'z':
		if (c + 13) > 122 {
			return (c - 13)
		} else {
			return (c + 13)
		}
	default:
		return c
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
