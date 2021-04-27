package main

import (
	"fmt"
)

func main() {
	const a,b,c=10,"test",10.8
	switch a{
	case 0:fmt.Println("0")
	case 1:fmt.Println("1")
	case 10:
		fmt.Println("my 10")
		fallthrough
	case 2:fmt.Println("2")
	case 3:fmt.Println("3")
	}

	switch{
	case true:fmt.Println("true")
	case false:fmt.Println("false")
	}
}

