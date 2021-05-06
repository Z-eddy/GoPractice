package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "No name", "输入你的姓氏")
	id := flag.Int("id", -1, "输入ID")

	flag.Parse()
	fmt.Printf("name:%s id:%d", *name, *id)
}
