package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/twiml", twiml)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func twiml(w http.ResponseWriter, r *http.Request) {
	twiml := TwiML{Say: "Hello World!"}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
