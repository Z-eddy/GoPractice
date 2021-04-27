package main

import "fmt"

func main() {
	//直接赋值创建
	dirArr:=[]int{1,2,3}
	//make创建8个int元素的数组
	arr0:=make([]int,20)
	//复制整个引用
	arr1:=dirArr[:]
	arr2:=arr0
	//引用部分索引
	arr3:=arr2[1:]
	arr4:=arr3[:15]
	arr5:=arr4[0:12]

	//元素访问
	arr1[2]=12
	arr0[4]=88

	fmt.Println(arr5)

	//查看len/cap
	//fmt.Println(len(arr0),cap(arr0))
	//fmt.Println(len(arr2),cap(arr2))
	//fmt.Println(len(arr3),cap(arr3))
	//fmt.Println(len(arr4),cap(arr4))
	//fmt.Println(len(arr5),cap(arr5))

	//拷贝数组
	arrDest:=make([]int,len(arr5),cap(arr5)*2)
	copy(arrDest,arr5)
	arr5[1]=11
	fmt.Println(arrDest,arr5)
}

