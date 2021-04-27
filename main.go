package main

import "fmt"

func send(a uint32,c chan uint32){
	var i uint32
	for i= 0;i!=a;i++ {
		c <- i
	}

	close(c)
}

func main() {
	c:=make(chan uint32,100)

	send(20,c)

	for v:=range c{
		fmt.Println(v)
	}
}

