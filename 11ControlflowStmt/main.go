package main

import "fmt"

func main() {

	fmt.Println("If else in Golang")

	loginCount := 25
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >= 10 {
		result = "Suspicious user"
	} else {
		result = "No user"
	}

	fmt.Println("Result ", result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
}
