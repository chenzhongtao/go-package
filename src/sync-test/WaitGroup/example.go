package main //阻塞，直到WaitGroup中的所以过程完成。
import (
	"fmt"
	"sync"
)

func wgProcess(wg *sync.WaitGroup, id int) {
	fmt.Printf("process:%d is going!\n", id)
	//if id == 2 {
	//    return
	//}
	wg.Done()
}

func main() {
	//var wg sync.WaitGroup
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go wgProcess(wg, i)
	}
	wg.Wait()
	fmt.Println("after wait group")
}
