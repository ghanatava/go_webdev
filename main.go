package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Hello world")
}

func main(){
	http.HandleFunc("/",HelloWorld)
	http.ListenAndServe("localhost:8080",nil)
}