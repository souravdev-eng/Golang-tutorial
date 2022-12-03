package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "This needs to be the file. Add more things"

	file, err := os.Create("./myDoc.txt")
	checkNil(err)

	length, err := io.WriteString(file, content)
	checkNil(err)

	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./myDoc.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNil(err)

	fmt.Println("Text data inside the file: \n", string(dataByte))
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
