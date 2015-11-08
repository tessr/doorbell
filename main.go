package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Dial    string   `xml:",omitempty"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/call", call)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func call(w http.ResponseWriter, r *http.Request) {
	twiml := TwiML{Dial: os.Getenv("IPHONE")}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
