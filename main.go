package main

import "fmt"

type Book struct{
	name string
	author string
}

func foo(myBook *Book){
	fmt.Println(myBook.name)
	fmt.Println(myBook.author)
}

func main() {
	theBook:=Book{name:"nameA",author:"zr"}
	foo(&theBook)
}

