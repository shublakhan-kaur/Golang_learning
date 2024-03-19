package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myChannel := make(chan int, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel
		fmt.Println("isChannel OPen: ", isChannelOpen)
		fmt.Println("Value from channel:", val)
		// fmt.Println("Go routine Getting the value from the channel:", <-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		fmt.Println("GO routine Adding value to the Channel")
		myChannel <- 5
		// myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)
	wg.Wait()
	// /myChannel <- 5
	// fmt.Println(<-myChannel)
}
