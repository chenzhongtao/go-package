package main

import (
	"bufio"
	"fmt"
)

// Buffered 返回缓存中数据的长度
//func (b *Reader) Buffered() int

func main() {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	fmt.Println(br.Buffered())
	// 0

	br.Peek(1)
	fmt.Println(br.Buffered())
	// 18
}
