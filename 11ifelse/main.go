package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in Go")

	loginCount := 10
	var message string

	if loginCount <= 10 {
		message = "Regular User"
	} else if loginCount > 10 && loginCount <= 22 {
		message = "Watch Out"
	} else {
		message = "Default message"
	}

	fmt.Println(message)

}
