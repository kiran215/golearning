package main

import "fmt"

const jwtToken = 3000 // Public

func main() {
	var username string = "Kiran"
	fmt.Println(username)
	fmt.Println("Variables")
	fmt.Printf("Variable is of Type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	var smallValue int16 = 256
	fmt.Println(smallValue)
	fmt.Printf("Variable is of Type: %T \n", smallValue)

	var smallFloatValue float32 = 255.223336666666
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of Type: %T \n", smallFloatValue)

	//default vaalues and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of Type: %T \n", anotherVar)

	var anotherStringVar string
	fmt.Println(anotherStringVar)
	fmt.Printf("Variable is of Type: %T \n", anotherStringVar)

	//implicit type - without mentioning datatype

	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of Type: %T \n", website)

	// no var style - need to use : before = // volorous operator-Will not work outside main function

	websitee := "www.google.com"
	fmt.Println(websitee)
	fmt.Printf("Variable is of Type: %T \n", websitee)

	fmt.Println(jwtToken)
	fmt.Printf("Variable is of Type: %T \n", jwtToken)

}
