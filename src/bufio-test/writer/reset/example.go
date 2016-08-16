package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//func (b *Writer) Reset(w io.Writer)
//Reset丢弃任何没有写入的缓存数据，清除任何错误并且重新将b指定它的输出结果指向w
func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}
