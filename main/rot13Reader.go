package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			if b[i] > 'a'+13 {
				b[i] -= 13
			} else if b[i] < 'a'+13 {
				b[i] += 13
			}
		} else if 'A' <= b[i] && b[i] <= 'Z' {
			if b[i] > 'A'+13 {
				b[i] -= 13
			} else if b[i] < 'A'+13 {
				b[i] += 13
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
