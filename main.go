package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	//大写暴露出来,才能被reflect感知
	FooMethod0()
}

type Foo struct {
	string
}

func (f Foo) FooMethod0() {
	fmt.Println(f.string)
}

func main() {
	myFoo := Foo{"test"}

	//类信息获得
	myType := reflect.TypeOf(myFoo)
	for i := 0; i != myType.NumField(); i++ {
		item := myType.Field(i)
		fmt.Println(item)
	}

	//接口方法获得
	var theInterface MyInterface = &myFoo
	myInterfaceType := reflect.TypeOf(theInterface)
	for i := 0; i != myInterfaceType.NumMethod(); i++ {
		item := myType.Method(i)
		fmt.Println(item)
	}
}
