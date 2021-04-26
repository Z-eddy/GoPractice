package main

import (
	"fmt"
)

func log(msg string){
	for i := 0 ; i < 100; i++ {
		fmt.Println(msg,i)
	}
}

func init(){
	fmt.Println("initialize...")
}

func init(){
	fmt.Println("init 1")
}

func main() {
	//自动多线程
	go log("A:")
	go log("B:")

	fmt.Println("finish")
}
