package main

import (
	"bufio"
	"fmt"
)

// NewWriterSize 将 wr 封装成一个拥有 size 大小缓存的 bufio.Writer 对象
// 如果 wr 的基类型就是 bufio.Writer 类型，而且拥有足够的缓存
// 则直接将 wr 转换为基类型并返回
//func NewWriterSize(wr io.Writer, size int) *Writer

// NewWriter 相当于 NewWriterSize(wr, 4096)
//func NewWriter(wr io.Writer) *Writer

//------------------------------------------------------------

// Flush 将缓存中的数据提交到底层的 io.Writer 中
//func (b *Writer) Flush() error

// Available 返回缓存中的可以空间
//func (b *Writer) Available() int

// Buffered 返回缓存中未提交的数据长度
//func (b *Writer) Buffered() int

// Write 将 p 中的数据写入 b 中，返回写入的字节数
// 如果写入的字节数小于 p 的长度，则返回一个错误信息
//func (b *Writer) Write(p []byte) (nn int, err error)

// WriteString 同 Write，只不过写入的是字符串
//func (b *Writer) WriteString(s string) (int, error)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGH")
	fmt.Println(bw.Available()) // 4088
	fmt.Println(bw.Buffered())  // 8
	fmt.Printf("%q\n", b)       // ""

	bw.Flush()
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0
	fmt.Printf("%q\n", b)       // "ABCEFG"
}
