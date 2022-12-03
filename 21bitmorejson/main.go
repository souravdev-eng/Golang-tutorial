package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"` // This will remove password from JSON Object
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []courses{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "JS"}},
		{"Micro Service", 499, "learncodeonline.in", "test354", []string{"software"}},
		{"Vue JS", 399, "learncodeonline.in", "abc367", []string{"JS", "front-end"}},
		{"Git and GitHub", 399, "learncodeonline.in", "sourav12", nil},
	}
	// package the data into JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
