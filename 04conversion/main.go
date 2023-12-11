package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")

	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating our pizza , ", input)

	//TrimSpace to remove the space in input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		if numRating > 5 {
			fmt.Println("Invalid rating")
		} else {
			fmt.Println(numRating + 1)
		}
	}
}
