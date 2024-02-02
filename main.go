package main

import (
	"fmt"
	"net/http"
	"time"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/ninja":
		fmt.Fprint(w, "Ninja")
	default:
		fmt.Fprint(w, "404 Page not Found")
	}
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	// http.ListenAndServe("localhost:8080", nil)
	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      nil,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	server.ListenAndServe()
}
