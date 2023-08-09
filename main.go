package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB setup for Golang")
	r := router.Router()
	fmt.Println("Server Is Getting Started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening 4000 port...")
}
