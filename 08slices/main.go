package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Banana", "Chickoo"}

	fmt.Printf("Type of fruitList %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Anjeer")
	fmt.Println(fruitList)

	fmt.Println(len(fruitList))
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = (fruitList[0:3])
	// //inclusive:non inclusive
	fmt.Println(fruitList)

	fruitList = (fruitList[:3])
	// //inclusive:non inclusive
	fmt.Println(fruitList)

	fruitList = (fruitList[:4])
	// //inclusive:non inclusive
	fmt.Println(fruitList)

	//memory allocation for array
	highScore := make([]int, 4)
	highScore[0] = 221
	highScore[1] = 225
	highScore[2] = 223
	highScore[3] = 224
	// highScore[4] = 225

	fmt.Println(highScore)

	//sort in increasing order
	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index

	var courses = []string{"React", "Java", "Go"}

	fmt.Println(courses)
	courses = append(courses, "Angular", "NET")
	fmt.Println(courses)
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
