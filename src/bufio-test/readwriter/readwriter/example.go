package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// ReadWriter 集成了 bufio.Reader 和 bufio.Writer
// 它实现了 io.ReadWriter 接口
//type ReadWriter struct {
//*Reader
//*Writer
//}

// NewReadWriter 封装 r 和 w 为一个 bufio.ReadWriter 对象
//func NewReadWriter(r *Reader, w *Writer) *ReadWriter
func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p)) //123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b) //asdf
}
