package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/souravdev-eng/mogoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listen on PORT 4000...")

}
