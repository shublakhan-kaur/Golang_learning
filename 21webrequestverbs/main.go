package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request verbs in Golang")
	PerformGetRequest()
	PerformPostRequest()
}

func PerformPostRequest() {
	myurl := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "learning",
	}
	requestBody := strings.NewReader(`
		{
			"courseName":"Golang",
			"price":0,
			"platform":"localhost:8080"
		}
	`)
	response, err := http.Post(myurl.String(), "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformGetRequest() {
	myurl := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "learning",
	}
	fmt.Println("myUrl", myurl)

	response, err := http.Get(myurl.String())
	if err != nil {
		panic(err)
	}
	fmt.Println("Response", response)

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("content", string(content))
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostReqForm() {
	myurl := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "learning",
	}

	data := url.Values{}
	data.Add("firstName", "Shublakhan")
	data.Add("lastName", "Kaur")
	data.Add("email", "kaur.shublakhan@gmail.com")
	response, err := http.PostForm(myurl.String(), data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
