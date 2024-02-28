package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)
	myNumber := 23
	var ptr = &myNumber //reference
	fmt.Println("Value of actual pointer is :", ptr)
	fmt.Println("Value of actual pointer is :", *ptr) //value inside the pointer

	*ptr = *ptr + 2
	fmt.Println("New Value is:", myNumber)
}
