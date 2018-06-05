package main

import "fmt"

func main(){
	fmt.Println("Hello, world")

	var x int = 5
	var word string = "This is a testing demo!"

	fmt.Println(x)
	fmt.Println(word, "It worked!")
	fmt.Println(addition(10, 32))
}

func addition(x int, y int) int {
	return x + y
}