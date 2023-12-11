package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//Need t ouse this format of date only
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.April, 5, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
