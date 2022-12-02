package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go")

	skills := []string{"React", "Docker", "GCP", "Kubernetes", "MongoDB", "Postgres"}

	// ------ One way -------
	for i := 0; i < len(skills); i++ {
		fmt.Println(skills[i])
	}

	// ------ Scend way -------
	for i := range skills {
		fmt.Println(skills[i])
	}

	// ------ Third way -------
	for index, skill := range skills {
		fmt.Printf("index is %v and value  is %v\n", index, skill)
	}

	randValue := 1

	for randValue < 10 {
		if randValue == 2 {
			goto lco
		}

		if randValue == 5 {
			randValue++
			continue
		}
		fmt.Println("Value is ", randValue)
		randValue++
	}

lco:
	fmt.Println("Jumping at lco")
}
