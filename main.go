package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	log.SetOutput(os.Stdout)

	myChl:=make(chan int)

	wait:=sync.WaitGroup{}
	wait.Add(1)//计数器增加goroutine数量

	go func()int{
		var sum int = 0
		for i:=0;i!=100000;i++{
			sum+=i
		}
		close(myChl)
		log.Println("finish")
		wait.Done()//计数器减一
		return sum
	}()

	wait.Wait()//阻塞到任务完成

	log.Println(myChl)
}

