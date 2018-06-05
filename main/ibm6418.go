package main //Success

import (
	"C"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//export localServer
func localServer() {
	if true {
		http.HandleFunc("/", message)
		log.Fatal(http.ListenAndServe(":8001", nil))
	} else {
		for count := 1; count < 3; count++ {
			if count != 3 {
				log.Println("Error Dectected - Server Restarting... ")
				localServer()
			} else {
				break
			}
		}
		return
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server Connected... ", r.URL.Path[1:])
}

//export makeRequestTest
func makeRequestTest(method *C.char, url *C.char) *C.char {

	//fmt.Printf("URL: %s\n", C.GoString(url)) PRINTS URL
	//fmt.Printf("METHOD: %s\n", C.GoString(method)) PRINTS METHOD

	ccURL := string(C.GoString(url))
	resp, err := http.Get(ccURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("This is body: %s", body)

	ccBODY := string(body)
	defer resp.Body.Close()

	//fmt.Printf("This is ccBODY: %s", ccBODY) CHECKS ccBODY value
	return C.CString(ccBODY)
}

func main() {
}
