package main

import "fmt"

func arrAdd(arr *[]int)int{
	var total int=0
	for _,val:=range *arr{
		total+=val
	}

	return total
}

func main() {
	arr0:=[]int{1,2,3}
	fmt.Println(arrAdd(&arr0))

	oneDArr:=[]int{1,2,3}
	for i,val:=range oneDArr {
		fmt.Println(i,val)
	}

	arrA:=[][]int{
		{1,2,3},
		{4,5,6},
	}
	row0:=[]int{7,8,9}
	arrA=append(arrA,row0)

	for i,val := range arrA {
		for j,jVal := range val{
			fmt.Println(i,j,jVal)
		}
	}
}

