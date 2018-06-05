package main //success

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	resqtimeout = 10
)

var (
	client *http.Client
)

func newClient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(resqtimeout) * time.Second,
	}
	return client
}

func messageToServer(w http.ResponseWriter, r *http.Request) { //Show that the server is live/connected
	fmt.Fprintln(w, "Server Connected", r.URL.Path[1:])
}

func newServer() { //Initialize server to use
	if true {
		http.HandleFunc("/", messageToServer)
		log.Fatal(http.ListenAndServe(":1894", nil))
	} else {
		fmt.Println("Server Connection Has Failed... Unknown Error Has Occured")
	}
}

func requestPanel(method string, url string) {
	time.Sleep(5000 * time.Millisecond)
	switch method {
	case "GET":
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(body))
		defer resp.Body.Close()
	default:
	}
}

func main() {
	client = newClient() //Declare Custom Client
	start := time.Now()  //Program Timer
	if true {
		go newServer() //Gorountine: 1 [newServer]
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Initializaing Connection To Server")
		for i := 0; i <= 2; i++ {
			fmt.Print(".")
			time.Sleep(750 * time.Millisecond)
		}
		fmt.Println()
	}
	time.Sleep(2000 * time.Millisecond)

	var currentRequest, currentURL string
	fmt.Print("\nENTER REQUEST: ")
	fmt.Scanf("%s", &currentRequest)

	fmt.Print("ENTER URL: ")
	fmt.Scanf("%s", &currentURL)

	fmt.Print("\nInitializing Request To Server")
	go requestPanel(currentRequest, currentURL)
	for i := 0; i <= 2; i++ {
		fmt.Print(".")
		time.Sleep(750 * time.Millisecond)
	}
	fmt.Println()

	time.Sleep(4000 * time.Millisecond)
	secs := time.Since(start).Seconds()
	fmt.Printf("\nTotal Time Elasped: %.4f\n", secs)
}
