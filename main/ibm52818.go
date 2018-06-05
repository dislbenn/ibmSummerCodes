package main

/*
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func newServer() {
	http.HandleFunc("/", newHandler)
	log.Fatal(http.ListenAndServe(":9005", nil))
}

func newClient() {
	//fmt.Println("Hello, new client")
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("newHandler was created")
}

func main() {
	var option string
	fmt.Print("Would you like to access a new server... Enter Y/N: ") //ACCESS NEW SERVER

	fmt.Scanf("%s", &option) //SERVER OPTION

	for {
		if option == "y" || option == "Y" { //User would like to access a new server
			fmt.Print("\nConstructing new server")
			go newServer()            //Initialize Goroutine for newServer
			for i := 0; i <= 2; i++ { //TIME WAITED
				fmt.Print(".")
				time.Sleep(1000 * time.Millisecond)
			}
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("\nNEW SERVER SUCCESSFULLY INITIALIZED")

			fmt.Print("\nConstructing new client")
			go newClient()            //Initialize GoRoutine for newClient
			for i := 0; i <= 2; i++ { //TIME WAITED
				fmt.Print(".")
				time.Sleep(1000 * time.Millisecond)
			}
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("\nNEW CLIENT SUCCESSFULLY INITIALIZED")
			break

		} else if option == "n" || option == "N" {
			time.Sleep(1500 * time.Millisecond)
			fmt.Print("NO SERVER WAS SELECTED, PREPARING TO TERMINATE THE PROGRAM")
			for i := 0; i <= 2; i++ {
				fmt.Print(".")
				time.Sleep(1750 * time.Millisecond)
			}
			break

		} else {
			time.Sleep(1500 * time.Millisecond)
			fmt.Print("SYSTEM HAS DECTECTED AN ERROR")
			for i := 0; i <= 2; i++ {
				fmt.Print(".")
				time.Sleep(1500 * time.Millisecond)
			}
			fmt.Print("\nWould you like to access a new server... Enter Y/N: ")
			fmt.Scanf("%s", &option)
			fmt.Println()
		}
	}
	fmt.Println("\nTHE PROGRAM HAS NOW ENDED\n")
}
*/
