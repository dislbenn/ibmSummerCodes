package main

import (
	"C"         // Converts to C
	"fmt"       // Input/Output
	"io/ioutil" // Read the body of the server
	"log"       // Log Errors
	"net/http"  // Access Server
)
import "time"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server Connected...", r.URL.Path[1:])
}

//export localServer
func localServer() {
	if true {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8001", nil))
	} else {
		fmt.Println("Connection to server resulted in a failure...")
	}
}

//export getRequest
func getRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//-------------------------------------------------------------------------------
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//-------------------------------------------------------------------------------
	log.Println(string(body))
	defer resp.Body.Close()
	//-------------------------------------------------------------------------------
	return body
}

//export postRequest
func postRequest() {

}

//-------------------------------------------------------------------------------
func main() {
	fmt.Print("Initialing localServer")
	go localServer()
	for i := 0; i <= 2; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	//-------------------------------------------------------------------------------
	time.Sleep(1500 * time.Millisecond)
	//-------------------------------------------------------------------------------
	var url string

	fmt.Print("ENTER URL: ")
	fmt.Scanf("%s", &url)
	fmt.Println()
	//-------------------------------------------------------------------------------
	start := time.Now()
	//-------------------------------------------------------------------------------
	for i := 1; i <= 100; i++ {
		if i <= 25 {
			fmt.Printf("Request:#%d -- ", i)
			getRequest(url)
		}
		//-------------------------------------------------------------------------------
		secs := time.Since(start).Seconds()
	}
	fmt.Printf("\nTime Elasped: %.4f\n", secs)
	fmt.Println()

}
