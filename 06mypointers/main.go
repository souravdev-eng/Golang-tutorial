package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// declare a pointer
	// var ptr *int
	// fmt.Println("Value of the pointer", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of pointer", ptr)
	fmt.Println("Value of pointer", *ptr)
}
