package main

import "fmt"

func main() {
	fmt.Println("My struts")

	kiran := User{"Kiran", "kiran@gmail.com", true, 28}
	fmt.Println(kiran)

	//+v to get details with parameters
	fmt.Printf("user details are %+v\n", kiran)

	fmt.Printf("Name %v\n", kiran.Name)
	fmt.Printf("Name %v\n", kiran.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
