package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("Maths in Golang")
	// TO generate a random number usinh math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// Generate random number using crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
