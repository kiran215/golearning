package main

import "fmt"

func main() {
	fmt.Println("Functions ")
	greeter()

	result := adder(3, 5)
	fmt.Println("Value of addition is - ", result)

	resultAdderPro := proAdder(3, 5, 8, 7)
	fmt.Println("Value of addition is - ", resultAdderPro)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

//... can take multiple values
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total = total + val
	}
	return total
}
func greeter() {
	fmt.Println("Namaste from Golang")
}
