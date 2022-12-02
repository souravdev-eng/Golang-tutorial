package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Go")
	var fruitList = []string{"Apple", "Banana", "Orange"}

	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)

	// This [1:] means it's ignore the first element of this array

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Make
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 412
	highScores[3] = 861

	highScores = append(highScores, 666, 999, 333)
	//Sort the int Slices
	sort.Ints(highScores)
	fmt.Println(highScores)

	// return is the value is sorted or not
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from a sliced based on index
	var course = []string{"react", "vue", "next", "docker"}
	fmt.Println(course)

	var index int = 1
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
