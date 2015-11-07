package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func call(w http.ResponseWriter, r *http.Request) {
	accountSid := os.Getenv("SID")
	authToken := os.Getenv("TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Calls.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", os.Getenv("IPHONE"))
	v.Set("From", os.Getenv("TWPHONE"))
	v.Set("Url", "http://"+os.Getenv("NAME")+"herokuapp.com/twiml")
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make request
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
		w.Write([]byte("womp. womp."))
	}

}
