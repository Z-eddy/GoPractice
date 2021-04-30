package main

import (
	"fmt"
	"net/http"
)

type (
	HelloHandler struct{}
	WorldHandler struct{}
)

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func (theWorld WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	hello := HelloHandler{}
	//world := WorldHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: &hello,
	}
	server.ListenAndServe()
}
