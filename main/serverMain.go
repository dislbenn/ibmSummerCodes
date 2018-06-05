package main

/*
import (
	"fmt"       // I/O
	"io/ioutil" //Reading the HTTP response body
	"log"       //
	"net/http"  //Making http requests
	"time"      //Time
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connection was successful, welcome %s!", r.URL.Path[1:])
}

func makeRequest() {
	start := time.Now() //Starts the timer

	resp1, _ := http.Get("https://yahoo.com") //Getting the body
	resp2, _ := http.Get("https://google.com")

	secs := time.Since(start).Seconds() //The amount of time passed

	body, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	fmt.Printf("Seconds Passed:  %f", secs)

	body, err = ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	fmt.Printf("Seconds Passed:  %f", secs)
}

func main() {
	http.HandleFunc("/", handler)
	makeRequest()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
