package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to golang web server")
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"
	// form data

	data := url.Values{}
	data.Add("firstName", "Sourav")
	data.Add("lastName", "Majumdar")

	response, err := http.PostForm(myUrl, data)
	ThrowPanic(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content of POST Form :", string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake JSON data
	requestBody := strings.NewReader(`
		{
			"name": "Sourav",
			"role": "Full Stack Developer",
			"exp" :	3
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	ThrowPanic(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content of POST :", string(content))
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	ThrowPanic(err)

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())

}

func ThrowPanic(err error) {
	if err != nil {
		panic(err)
	}
}
