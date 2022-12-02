package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	result := adder(3, 3)
	result2, myMessage := proAdder(3, 3, 6, 10, 900)
	fmt.Println(result)
	fmt.Println("Value of ProAdder", result2)
	fmt.Println("Message of ProAdder", myMessage)
	greeter("Sourav")
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Hello from proAdder"
}

func greeter(name string) {
	fmt.Printf("Welcome %v to golang", name)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
