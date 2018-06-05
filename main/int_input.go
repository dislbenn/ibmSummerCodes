package main

import "C"
import "fmt"

//export int_input
func int_input(num int) {
	fmt.Println(num + 2)
}

func main() {}
