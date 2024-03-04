package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline", "abc123", []string{"web-dev", "JS"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline", "bcd123", []string{"full-stack", "JS"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline", "hit123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
