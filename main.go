package main

import (
	"fmt"
	"log"
	"os"
	pb "practice/myPackage"
)

func main() {
	fmt.Println(pb.Num0)
	log.SetOutput(os.Stdout)
	log.Println("temp",pb.Num0)

	fmt.Println("finish")
}

