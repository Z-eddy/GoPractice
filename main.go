package main

import "fmt"

func main() {
	defer fmt.Println("it's defer")

	sum:=0
	fmt.Println("start")
	for i:=0;i!=100;i++{
		sum+=i
	}
	fmt.Println("finish")
}

