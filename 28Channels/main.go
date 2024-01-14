package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	//Receive only
	go func(myCh <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh

		if isChannelOpen {

			fmt.Println("Channel open read value")
			fmt.Println(val)

		} else {
			fmt.Println("Channel is closed")
		}
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	//send only
	go func(myCh chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		close(myCh)

		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
	// myCh <- 5
	// fmt.Println(<-myCh)
}
