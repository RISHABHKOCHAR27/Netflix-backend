package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/route"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Server is getting started...")
	r := route.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running...")
}
