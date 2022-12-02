package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map in Go")
	// (map[key]value)
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["GO"] = "Golang"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS short for", languages["JS"])
	// delete value
	delete(languages, "PY")
	fmt.Println("List of all languages", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("FOR key %v, value is %v\n", key, value)
	}
}
