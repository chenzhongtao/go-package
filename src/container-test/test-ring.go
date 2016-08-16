//ring是一个环形链表
package main

import (
	"container/ring" //闭环包引入，详见/usr/local/go/src/pkg/container/ring
	"fmt"
)

func main() {
	//创建10个元素的闭环
	r := ring.New(10)

	//给闭环中的元素赋值
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	//循环打印闭环中的元素值
	r.Do(
		func(p interface{}) {
			println(p)
		})

	//获得当前元素之后的第5个元素
	r5 := r.Move(5)
	fmt.Println(r5)
	fmt.Println(r)

	//链接当前元素r与r5，相当于删除了r与r5之间的元素
	r1 := r.Link(r5)
	fmt.Println(r1)
	fmt.Println(r)
}

/*
(0x4b94a0,0xc820074188)
(0x4b94a0,0xc8200741b0)
(0x4b94a0,0xc8200741b8)
(0x4b94a0,0xc8200741c0)
(0x4b94a0,0xc8200741c8)
(0x4b94a0,0xc8200741d0)
(0x4b94a0,0xc8200741d8)
(0x4b94a0,0xc8200741e0)
(0x4b94a0,0xc8200741e8)
(0x4b94a0,0xc8200741f0)
&{0xc82006e140 0xc82006e100 6}
&{0xc82006e0a0 0xc82006e1a0 1}
&{0xc82006e0c0 0xc82006e100 2}
&{0xc82006e120 0xc82006e1a0 1}

*/
