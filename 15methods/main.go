package main

import "fmt"

func main() {
	fmt.Println("Methods in Go")
	sourav := User{"Sourav", "sourav@gmail.com", false, 25}
	sourav.GetStatus()
	sourav.NewMail()
	fmt.Println("Email Original", sourav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Printf("Is User active: %+v\n", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is", u.Email)
}
