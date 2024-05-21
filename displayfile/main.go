package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print(string(content))
}
