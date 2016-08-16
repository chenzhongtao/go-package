package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

//  sync.Once.Do(f func())是一个挺有趣的东西,能保证once只执行一次，
//无论你是否更换once.Do(xx)这里的方法,这个sync.Once块只会执行一次。
func main() {

	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			// onced函数不会被调用
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
