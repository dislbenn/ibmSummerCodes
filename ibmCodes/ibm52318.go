package main

import "fmt"
import "C"

func title() string {return "TITLE: ibm52318"}
//--------------------------------------------
func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println(title()) //CODE NAME
/*
	The objective of the code is to 
*/
	fmt.Print("Testing")
	s := "Disaiah Bennett"

	fmt.Println(len(s))
	fmt.Println(s[0:10])
}