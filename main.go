package main

import (
	"container/list"
	"fmt"
)

func main() {
	myLi := list.New()
	myLi.PushBack(99)
	myLi.PushBack("theTwo")

	for i := 0; i != 10; i++ {
		myLi.PushBack(i)
	}

	myLi.Remove(myLi.Back())

	for beg := myLi.Front(); beg != nil; beg = beg.Next() {
		fmt.Println(beg.Value)
	}
}
