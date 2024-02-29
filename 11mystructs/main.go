package main

import "fmt"

func main() {
	fmt.Println("Working with Structs")
	shub := User{"Shublakhan", "kaur.shublakhan@gmail.com", true, 30}
	fmt.Println(shub)
	fmt.Printf("User details are: %+v\n", shub)
	fmt.Printf("Name is %v and email is %v.", shub.Name, shub.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
