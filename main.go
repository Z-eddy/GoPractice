package main

import (
	"errors"
	"fmt"
)

func divideErr(a,b float64)(float64,error){
	if b!=0 {
		return a / b, nil
	}else{
		return 0,errors.New("error!!!")
	}
}

func main() {
	fmt.Println(divideErr(5,0))
}

