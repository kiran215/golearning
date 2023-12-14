package main

import "fmt"

func main() {

	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["JS"] = "JAVASCRIPT"
	languages["PY"] = "PYTHON"
	languages["GO"] = "GOLANG"

	fmt.Println("List of all langs , ", languages)

	fmt.Println("JS For , ", languages["JS"])
	delete(languages, "JS")
	fmt.Println("List of all langs , ", languages)

	//for loop to iterate map
	for key, value := range languages {
		fmt.Printf("For key %v , value is  %v\n", key, value)

	}

}
