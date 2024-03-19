package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shublakhan-kaur/Golang_learning/mongoapi/router"
)

func main() {
	fmt.Println("Working with mongoDB and golang")
	r := router.Routes()
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000....")

}
