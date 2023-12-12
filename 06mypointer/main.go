package main

import "fmt"

func main() {

	fmt.Println("Pointer")

	//creating a pointer
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	//& creating a pointer which is referencing to some memory
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr) //to see inside the pointer

	*ptr = *ptr * 2

	fmt.Println("New value is ", *ptr)

}
