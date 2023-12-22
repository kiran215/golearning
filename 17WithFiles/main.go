package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "Hello we are working with files!!\n"

	//Create file
	file, err := os.Create("./MyNewFile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilError(err)

	//Write string in file
	for i := 0; i < 5; i++ {
		length, err := io.WriteString(file, content)
		if err != nil {
			panic(err)
		}
		fmt.Println("Length is - ", length)
	}

	//Close file
	file.Close()

	readFile("./MyNewFile.txt")
}

func readFile(fileName string) {

	//Read file in bytes
	databyte, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Text inside file is -", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
