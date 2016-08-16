package main

import (
	"container/list"
	"fmt"
)

func Example() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	l.PushBack("a")

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		//i := e.Value.(int)
		//fmt.Println(i)
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// a
}

func main() {
	Example()
}
