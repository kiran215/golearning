package main

import "fmt"

func main() {
	fmt.Println("In methods")

	kiran := User{"Kiran", "kiran@gmail.com", true, 28}
	fmt.Println(kiran)

	//+v to get details with parameters
	fmt.Printf("user details are %+v\n", kiran)

	fmt.Printf("Name %v\n", kiran.Name)
	fmt.Printf("Name %v\n", kiran.Email)
	kiran.GetStatus()
	kiran.SetEmail("Uday.sadaye@gmail.com")
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//method
func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)

}

func (u User) GetEmail() {
	fmt.Println("User email : ", u.Email)

}

func (u User) SetEmail(email string) {
	u.Email = email
	fmt.Println("User email ,", u.Email)
}
