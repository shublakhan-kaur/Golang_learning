package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentId=fsdf4414"

func main() {
	fmt.Println("Welcome to handling urls in Golang")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path", result.Path)
	fmt.Println("Port", result.Port())
	fmt.Println("RawQuery", result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("The type of query params is %T\n", queryParams)
	fmt.Println(queryParams["coursename"])

	for _, param := range queryParams {
		fmt.Println("Param is", param)
	}

	// TO create a url

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost:8080",
		Path:     "learning",
		RawQuery: "name=Shub",
	}

	createUrl := partsOfUrl.String()
	fmt.Println("Created Url:", createUrl)
}
