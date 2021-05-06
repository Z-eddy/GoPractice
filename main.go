package main

import (
	"fmt"
	"sync"
)

var myMap = sync.Map{}
var theWait = sync.WaitGroup{}

func wFoo(num int) {
	val := fmt.Sprintf("test %d", num)
	myMap.Store(num, val)
	theWait.Done()
}

func main() {
	theWait.Add(10)

	for i := 0; i != 10; i++ {
		go wFoo(i)
	}

	theWait.Wait()

	//读取单个的值
	val, _ := myMap.Load(3)
	fmt.Println(val)

	//遍历
	myMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)

		//断言实现类型转换
		theStr := value.(string)
		str := theStr + " other"
		fmt.Println(str)

		return true
	})
}
