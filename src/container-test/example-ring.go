package main

import (
	"container/ring"
	"fmt"
)

// For debugging - keep around.
func dump(r *ring.Ring) {
	if r == nil {
		fmt.Println("empty")
		return
	}
	i, n := 0, r.Len()
	for p := r; i < n; p = p.Next() {
		fmt.Printf("%4d: %p %v \n", i, p, p.Value)
		i++
	}
	fmt.Println()
}

func makeN(n int) *ring.Ring {
	r := ring.New(n)
	for i := 1; i <= n; i++ {
		r.Value = i
		r = r.Next()
	}
	return r
}

func sumN(n int) int { return (n*n + n) / 2 }

func TestNew() {

	r := makeN(10)
	dump(r)

}

func TestEmpty() {

	r := makeN(0) // 相当于 var r *ring.Ring  所以这里的nil还可以调用函数 所以有些函数要判断 this != nil(c++的this,go是没有的)

	fmt.Println("len: ", r.Len())

}

func main() {
	TestNew()
	TestEmpty()
}
