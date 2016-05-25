package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", showHelloWorld)
	http.ListenAndServe(":8080", nil)
}

func showHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
