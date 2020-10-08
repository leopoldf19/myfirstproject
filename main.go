package main

import (
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
