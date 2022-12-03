package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request in Go")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is type %T\n", res)

	// Caller's responsibility to close the connection
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Println("Response is", content)

}
