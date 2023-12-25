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
	DecodeJson()
}

func DecodeJson() {

	finalJsonFromWeb := []byte(`{
		"courseName": "ReacJs boot2",
		"price": 219,
		"website": "Leanrcodeonline.in",
		"tags": [
                        "web",
                        "dev",
                        "js"
						]
}
`)

	var lcoCourse course
	if json.Valid(finalJsonFromWeb) {
		fmt.Println("JSON is valid")
		json.Unmarshal(finalJsonFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
		fmt.Println(lcoCourse.Name)
	} else {
		fmt.Println("JSON invalid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(finalJsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v \n", k, v)
	}

}
