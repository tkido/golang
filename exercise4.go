package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
  n, err = rot.r.Read(p)
  for i:=0; i < n; i++ {
      if p[i] > 'A' && p[i] < 'z' {
          if p[i] > 'z' - 13 {
             p[i] = p[i]-13
          } else {
             p[i] = p[i]+13
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
