package main

import (
	"encoding/json"
	"fmt"
)

// "-" will remove the tag from json data
type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Bit more json")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReacJs boot", 299, "Leanrcodeonline.in", "136", []string{"web", "dev", "js"}},
		{"ReacJs boot1", 219, "Leanrcodeonline.in", "136", []string{"web", "dev", "js"}},
		{"ReacJs boot2", 219, "Leanrcodeonline.in", "136", nil},
	}

	//Package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
