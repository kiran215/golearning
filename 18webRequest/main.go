package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const url = "https://lco.dev"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is %T\n", response)
	fmt.Println("Response is ,", response.StatusCode)

	defer response.Body.Close() // Caller's responsibilty to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
