package main

import (
	"bufio"
	"fmt"
)

// WriteTo 实现了 io.WriterTo 接口
//func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
func main() {
	s := strings.NewReader("ABCEFG")
	br := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))

	br.WriteTo(b)
	fmt.Printf("%s\n", b)
	// ABCEFG
}
