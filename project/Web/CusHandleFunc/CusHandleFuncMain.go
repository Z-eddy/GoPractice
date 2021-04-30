package main

import (
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func worldHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	//myHandle := http.NewServeMux()
	//myHandle.HandleFunc("/hello", helloHandle)
	//myHandle.HandleFunc("/world", worldHandle)
	server := http.Server{
		Addr: ":8080",
		//Handler: myHandle,
	}
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/world", worldHandle)

	server.ListenAndServe()
}
