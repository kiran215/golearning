package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//To take the input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza")

	//Read the input
	//comma ok //err err
	//If everything is ok then input value will be return and if any error comes then err will catch
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating ,", input)
	fmt.Printf("Thanks for rating %T,", input)
	fmt.Println("Thanks for rating ,", err)

}
