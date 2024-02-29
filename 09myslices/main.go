package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on Slices")

	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Orange", "Banana")
	fmt.Println("FruitList is: ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("FruitList after slicing: ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 255
	highScores[1] = 945
	highScores[2] = 675
	highScores[3] = 885

	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//How to remove value from slices based on index

	var courses = []string{"reactjs", "javascript", "java", "python", "Angular", "Spring"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
