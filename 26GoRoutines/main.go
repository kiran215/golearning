package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"signal1"}

// wg -wait group
var wg sync.WaitGroup // pointers
var mut sync.Mutex    // pointers

func main() {
	// go greeter("Hello")
	// greeter("Kiran")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://fb.com1",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

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
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Failed to get status code for %s\n", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf(" %d status code for %s\n", res.StatusCode, endpoint)
	}
}
