package main

import (
	"fmt"
	"sync"
)

var rwLocker = sync.RWMutex{}
var wait = sync.WaitGroup{}
var gNum = 0

func rFoo(num int) {
	rwLocker.RLock()
	defer rwLocker.RUnlock()
	defer wait.Done()
	fmt.Println("read", num, gNum)
	gNum++
}

func wFoo(num int) {
	rwLocker.Lock()
	defer rwLocker.Unlock()
	defer wait.Done()
	fmt.Println("write", num, gNum)
	gNum++
}

func main() {
	wait.Add(20)
	for i := 0; i != 10; i++ {
		go rFoo(i)
	}
	for i := 0; i != 10; i++ {
		go wFoo(i)
	}
	wait.Wait()
}
