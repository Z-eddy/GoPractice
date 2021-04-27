package main

import "fmt"

func main() {
	func(a int){
		fmt.Println(a+1)
	}(5)
}

