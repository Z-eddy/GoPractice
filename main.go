package main

import "fmt"

func foo(sA string,b int)(int,string){
	return b+1,sA+"post"
}

func fooAdd(a,b int)int{
	return a+b
}

func main() {
	a,b:=foo("myTest",9)
	c:=fooAdd(2,3)
	fmt.Println(a,b,c)
}

