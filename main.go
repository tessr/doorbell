package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
