package main

import "fmt"

func sum(a uint32,c chan uint32){
	var i,total uint32
	total=0
	for i= 0;i!=a;i++ {
		total+=i
	}

	c <- total
}

func main() {
	var a,b,d uint32
	a=0
	b=10
	d=20

	c:=make(chan uint32)

	go sum(a,c)
	go sum(b,c)
	go sum(d,c)

	x,y,z:=<-c,<-c,<-c

	fmt.Println(x,y,z)
}

