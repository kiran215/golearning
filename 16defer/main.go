package main

import "fmt"

func main() {
	//defer means postpone
	//execution will be done later for defer statement
	defer fmt.Println("Hello Defer")
	defer fmt.Println("One Defer")
	defer fmt.Println("Two Defer")
	fmt.Println("World Defer")
	myDefer()
}

func myDefer() {
	count := 5
	for i := 0; i < count; i++ {
		defer fmt.Println(i)
	}
}
