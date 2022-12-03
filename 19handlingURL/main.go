package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hjfhjdhf"

func main() {
	fmt.Println("Welcome to handling URL in golang")
	// fmt.Println(myUrl)
	// parsing url
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("Type of the query parameters is %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=sourav",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("Another URL", anotherURL)
}
