package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.Mutex
var m *sync.Mutex

//go func 和主线程之间的关系是并行和竞争关系
func main() {

	m = new(sync.Mutex)

	go lock(1)
	time.Sleep(3 * time.Second)

	fmt.Printf("%s\n", "exit!")

}
func lock(i int) {
	println(i, "lock start")

	m.Lock()
	println(i, "lock")

	time.Sleep(3 * time.Second)

	m.Unlock()
	println(i, "unlock")
}
