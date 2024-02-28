package main

import "fmt"

const LoginToken string = "fsdfdsfdsf" // Public in Go lang all the functions and
//variables with capital letter in the beginning
//are by default public

func main() {
	fmt.Println("Variables")
	var username string = "Shublakhan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.454545465
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
