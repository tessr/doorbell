package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
