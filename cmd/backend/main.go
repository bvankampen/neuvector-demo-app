package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Text string `json:"text"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	action := req.URL.Query().Get("action")
	r := Response{}

	fmt.Printf("Request from: %s\n", req.RemoteAddr)
	fmt.Printf("Action: %s\n", action)

	switch action {
	case "ssn":
		r.Text = "685-14-7195"
	case "cc":
		r.Text = "VISA, 4716851810122187, 9/2028, 931"
	case "text":
		r.Text = RandStringBytes(32)
	default:
		r.Text = ""
	}

	jsonResponse, _ := json.Marshal(r)

	fmt.Printf("Response: %s\n", string(jsonResponse))

	fmt.Fprintf(w, string(jsonResponse))

}

func main() {
	http.HandleFunc("/backend", indexHandler)
	err := http.ListenAndServe(strings.Trim(os.Getenv("APP_BACKEND_IP"), " ")+":"+strings.Trim(os.Getenv("APP_BACKEND_PORT"), " "), nil)
	if err != nil {
		fmt.Println(err)
	}
}
