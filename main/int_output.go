package main

import (
	"C"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//export int_output
func int_output() int {
	return rand.Intn(10000)
}

func main() {
}
