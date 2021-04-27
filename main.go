package main

import "fmt"

func fooAdd(a *int){
	*a += 1
}

func main() {
	a:=3
	fooAdd(&a)
	fmt.Println(a)
}

