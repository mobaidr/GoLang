package main

import (
	"fmt"
	"net/http"
)

// This small program is to show case how to debug GoLang code
// running in Docker from VSCode
// https://dev.to/bruc3mackenzi3/debugging-go-inside-docker-using-vscode-4f67

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/Hello endpoint called")
	fmt.Fprintf(w, "Hello with Happy Debugging\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server Up and Listening....")
	http.ListenAndServe(":8090", nil)
}
