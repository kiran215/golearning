package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {

	fmt.Println("Mongo DB API")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Mongo DB API listening at port 4000")

}
