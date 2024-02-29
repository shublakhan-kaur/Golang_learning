package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	//fruitList[2] = "Pear"
	fruitList[3] = "Grapes"
	fmt.Println("Fruit List", fruitList)
	fmt.Println("Length of fruit is", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Length of Veg is", len(vegList))

}
