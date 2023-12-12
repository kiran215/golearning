package main

import "fmt"

func main() {

	fmt.Println("Array in Golang")

	var fruitList [7]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"
	fruitList[6] = "Banana"

	fmt.Println(fruitList)
	fmt.Println("Lenght of array ", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato"}
	fmt.Println("Lenght of array ", len(vegList))

}
