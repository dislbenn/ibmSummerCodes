package main

import (
	"C"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	requestTimer = 10
)

var (
	client *http.Client
)

func newclient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(requestTimer) * time.Second,
	}
	return client
}

//export server
func server() {
	if true {
		http.HandleFunc("/", message)
		log.Fatal(http.ListenAndServe(":4040", nil))
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Server Connection Complete...", r.URL.Path[1:])
}

//export requestToServer
func requestToServer(method string, url string) {

	switch method {
	case "GET":
		resp, err := http.Get("http://localhost:4040")
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(body))
		defer resp.Body.Close()
	case "POST":
		resp, err := http.Post("http://localhost:4040", "POST", nil)
		http.Post("http://localhost:4040", "Body", body)
	default:
		fmt.Println("SYSTEM HAS DETECTED AN ERROR...")
	}
}

func main() {
	client := newclient()
	server()
}
