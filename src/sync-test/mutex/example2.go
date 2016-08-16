package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name   string
	Locker *sync.Mutex
}

func (u *User) SetName(wati *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		u.Locker.Unlock()
		wati.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock set name:", name)
	time.Sleep(1 * time.Second)
	u.Name = name
}

func (u *User) GetName(wati *sync.WaitGroup) {
	defer func() {
		fmt.Println("Unlock get name:", u.Name)
		u.Locker.Unlock()
		wati.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock get name:", u.Name)
	time.Sleep(1 * time.Second)
}

func main() {
	user := User{}
	user.Locker = new(sync.Mutex)
	wait := &sync.WaitGroup{}
	names := []string{"a", "b", "c"}
	for _, name := range names {
		wait.Add(2)
		go user.SetName(wait, name)
		go user.GetName(wait)
	}

	wait.Wait()
}
