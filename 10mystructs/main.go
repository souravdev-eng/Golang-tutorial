package main

import (
	"fmt"
)

func main() {
	fmt.Println("Struct in Go")
	// no inheritance in golang; No super or parent

	sourav := User{"Sourav", "sourav@gmail.com", true, 25}
	fmt.Println("User details", sourav)
	fmt.Printf("User details are: %+v\n", sourav)
	fmt.Printf("User name is %v and email is %v", sourav.Name, sourav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
