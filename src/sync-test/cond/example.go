package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wait := sync.WaitGroup{}
	locker := new(sync.Mutex)

	//每个Cond实例都有一个相关的锁（一般是*Mutex或*RWMutex类型的值），它必须在改变条件时或者
	//调用Wait方法时保持锁定。Cond可以创建为其他结构体的字段，Cond在开始使用后不能被拷贝。
	cond := sync.NewCond(locker)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wait.Done()
			wait.Add(1)
			cond.L.Lock()
			fmt.Println("Waiting start...")
			//Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L。
			//和其他系统不同，Wait除非被Broadcast或者Signal唤醒，不会主动返回。
			cond.Wait()
			fmt.Println("Waiting end...")
			cond.L.Unlock()

			fmt.Println("Goroutine run. Number:", i)
		}(i)
	}

	time.Sleep(2e9)
	cond.L.Lock()
	// Signal唤醒等待c的一个线程（如果存在）。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定
	// 对应还有Broadcast唤醒所有等待c的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(2e9)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(2e9)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	wait.Wait()
}
