package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello web: %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, r.URL.Path)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", handler)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: multiplexer,
	}
	server.ListenAndServe()
}
