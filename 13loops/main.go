package main

import "fmt"

func main() {
	fmt.Println("In Loop")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	for day := range days {
		fmt.Println(days[day])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 6 {
			rogueValue++
			continue
		}
		fmt.Println("Value is ", rogueValue)
		rogueValue++

	}
lco:
	fmt.Println("Jumping at LCO")
}
