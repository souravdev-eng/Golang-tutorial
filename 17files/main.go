package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "This needs to be the file. Add more things"

	file, err := os.Create("./myDoc.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", length)
	file.Close()
}
