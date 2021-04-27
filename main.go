package main

import "fmt"

func main() {
	//myMap:=make(map[int]string)
	myMap:=map[int]string{}
	myMap[0]="my0"
	myMap[1]="my1"
	myMap[2]="my2"

	for key,val:=range myMap {
		fmt.Println(key,val)
	}

	delete(myMap,1)
	fmt.Println(myMap)
}

