package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)
	for index, value := range days {
		fmt.Printf("Index is %v and value is %v \n", index, value)
	}

	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at label")
}
