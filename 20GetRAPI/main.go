package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Inside API Calling")
	// PerformGetReq()
	PerformPostReq()
}

func PerformGetReq() {

	//Call GET API
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)

	if err != nil {
		panic(nil)
	}

	defer response.Body.Close()
	// fmt.Println(response.Body)
	fmt.Println(response.StatusCode)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

	var responseString strings.Builder

	//to get the byte count uses strings.Builder write method
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)

	//responseString holds the content of response
	fmt.Println(responseString.String())

}

// how to make POS request with JSON data
func PerformPostReq() {
	const myUrl = "https://api.sahisavings.com:8443/sahisavingoauth/auth/signIn"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"phoneNumber":"+919821787220",
		"password":"2149"}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	databytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(databytes))
}

// how to send form data in golang
func PerformFormReq() {

	const myUrl = "http://localhost:8000/postForm"

	//form data
	data := url.Values{}

	data.Add("firstname", "kiran")
	data.Add("lastname", "sadaye")
	data.Add("email", "kiran.sadaye@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
