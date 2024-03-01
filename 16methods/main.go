package main

import "fmt"

func main() {
	shub := User{
		"Shublakhan", "kaur.shublakhan@gmail.com", true, 30}
	fmt.Println("Shub details are: ", shub)
	fmt.Printf("Name is %v and Email is %v ", shub.Name, shub.Email)
	shub.getStatus()
	shub.getNewMail()
	fmt.Printf("Name is %v and Email is %v ", shub.Name, shub.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getNewMail() {
	u.Email = "newMail@text.com"
	fmt.Println("New email is: ", u.Email)
}

func (u User) getStatus() {
	fmt.Println("Is user active: ", u.Status)
}
