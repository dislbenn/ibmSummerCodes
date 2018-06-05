package main

/*
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Create New local server
func newLocalServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
	fmt.Println("Server is now active")
}

func makeNewRequest() {
	//resp1, _ := http.Get("www.yahoo.com") // Get the body of server 1
	time.Sleep(10) // Sleep for 10 seconds
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Testing, Testing\n")
	fmt.Println("The server test have is now conducting")
}

func main() {
	var op string            //Option
	timedStart := time.Now() //Start Timer
	fmt.Print("Would you like to access new server: Enter yes/no\n")
	fmt.Scanf("%s", &op) // Accessing new server

	secondsPassed := time.Since(timedStart).Seconds() //Seconds passed

	if op == "yes" || op == "Yes" || op == "YES" {
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\nSetting up new local server")
		go newLocalServer()
		for i := 0; i <= 2; i++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Print(".")
		}
		fmt.Println()

		fmt.Print("Setting up new client")
		for i := 0; i <= 2; i++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Print(".")
		}
		fmt.Println()

		fmt.Printf("Time Elasped: %.4fs\n\n", secondsPassed)
		makeNewRequest()

	} else if op == "no" || op == "No" || op == "NO" {
		time.Sleep(2 * time.Second)
		fmt.Println("\nNo servers are currently active... \n")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("ERROR HAS OCCURED::INVALID OPTION HAS BEEN SELECTED\n")
		time.Sleep(2 * time.Second)
	}

	time.Sleep(1500 * time.Millisecond)
	fmt.Print("Program is now closing")

	for i := 0; i <= 2; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
	fmt.Printf("Time Elasped: %.4fs\n\n", secondsPassed)

	time.Sleep(2500 * time.Millisecond)
	fmt.Println("PROGRAM CLOSED SUCCESSFULLY\n")
}
*/
