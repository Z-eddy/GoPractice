package main

import (
	"fmt"
)

const (
	a=iota
	b
	c="test"
	d
	e=100
	f
	g=iota
	h
)

func main() {
	fmt.Println(a,b,c,d,e,f,g,h)
}

