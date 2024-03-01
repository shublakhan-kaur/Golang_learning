package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in the file Hello World!"
	file, err := os.Create("./newFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile("./newFile.txt")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
