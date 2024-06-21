package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/iamSuva/mongoapi/route"
)

func main() {
	fmt.Println("Hello world with go language")
	r :=router.Router()
    fmt.Println("server is getting started")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("server is running on port 4000")
	



}