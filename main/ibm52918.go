package main

/*
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	requestTimeout = 10 // Sets constant requestTimeout length
)

var (
	client      *http.Client //**
	atomicCount uint64       //**
)

func localServer() { // Create localServer
	http.HandleFunc("/", localHandler)
	log.Fatal(http.ListenAndServe(":3030", nil)) // Log to localhost:3030
}

func localHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handler successful constructed", r.URL.Path[1:]) //Writes message to the form
}

func localClient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
	return client
}

func makelocalRequest() {
	start := time.Now()
	resp, err := http.Get("http://localhost:3030")
	if err != nil {
		log.Fatal(err)
	}

	secs := time.Since(start).Seconds()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	log.Println(string(body))
	fmt.Printf("Seconds Elasped: %f", secs)

	_, err = os.Stdout.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client = localClient()
	go localServer()

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Constructing server...")

	makelocalRequest()
}
*/
