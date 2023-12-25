package main

import (
	"fmt"
	"net/url"
	"strings"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=react&paymentid=123"

func main() {

	fmt.Println("Into the Handling URL , ", myUrl)

	//parsing url
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)

	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(strings.Split(result.RawQuery, "&"))

	qparams := result.Query() //return map
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams.Get("paymentid"))

	for _, val := range qparams {
		fmt.Println("Param is , ", val)
	}

	//generate url

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=kiran",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
