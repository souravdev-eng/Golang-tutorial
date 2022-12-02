package main

import "fmt"

// Public variable start with capital letter
const LoggedInToken string = "this-is-a-token"

func main() {
	var username string = "Sourav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45555555555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "sourav.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	// ":=" called valours operator
	numberOfUser := 322
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	// Use public variable
	fmt.Println(LoggedInToken)
	fmt.Printf("Variable is of type: %T \n", LoggedInToken)
}
