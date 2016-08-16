package main

import (
	"bufio"
	"fmt"
)

// ReadFrom 实现了 io.ReaderFrom 接口
//func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("Hello 世界！")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)
	//bw.Flush()            //ReadFrom无需使用Flush，其自己已经写入．
	fmt.Println(b) // Hello 世界！
}
