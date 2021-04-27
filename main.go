package main

import "fmt"

func main() {
	for i := 0;i!=10;i++{
		fmt.Println(i)
	}

	valA:=10
	for valA>0 {
		println(valA)
		valA--
	}

	strA:=[]string{"stringA","stringB"}
	for i,val:=range strA{
		fmt.Printf("index:%d value:%s\n",i,val)
	}

	arrA:=[]int{1,2,3,4,5,6}
	for i,val:=range arrA{
		fmt.Printf("index:%d value:%d\n",i,val)
	}

	/*
	i:=0
	for{
		fmt.Println(i)
		i++
	}
	 */
}

