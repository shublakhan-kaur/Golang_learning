package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var signals = []string{"test"}

var mut sync.Mutex

func main() {
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	// go greeter("Hello")
	// greeter("World")
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	defer wg.Done()
	if err != nil {
		fmt.Println("OOPS endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
