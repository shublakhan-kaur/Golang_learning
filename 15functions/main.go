package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()
	result := adder(3, 5)
	proRes, proMessage := proAdder(2, 5, 7, 8, 9, 1)
	fmt.Println("Result of adder function is: ", result)
	fmt.Println(proMessage, proRes)
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi Pro result is"
}

func greeter() {
	fmt.Println("Namstey from Golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
