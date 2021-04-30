package main

import "fmt"

type (
	user struct {
		name string
	}
)

func (u user) notify(arg string) {
	u.name += arg
	fmt.Println("u", u)
}

func main() {
	myUser := new(user)
	myUser.name = "test"
	myUser.notify("other")

	fmt.Println("myUser", myUser)
}
