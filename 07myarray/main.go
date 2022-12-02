package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Go")

	var myArr [4]string
	myArr[0] = "Apple"
	myArr[1] = "Orange"
	// myArr[2] = "Banana"
	myArr[3] = "Grep"

	fmt.Println("Fruits list is: ", myArr)
	// This len() is refer the array deceleration size not the array length
	fmt.Println("Fruits list is: ", len(myArr))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", vegList)
}
