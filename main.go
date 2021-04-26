package main

import (
	"fmt"
)

const (
	c=8
	d=9
)

const e=88

func main() {
	//var a,b int= 1,2//可以声明int类型,也可以隐藏
	const a,b=1,2

	//覆盖外部的
	const (
		c=18
		d=19
	)

	fmt.Println(a,b,c,d,e)
}

