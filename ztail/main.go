package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func printError(err error) {
	fmt.Println(err)
}

func printContent(filename string, content []byte, count int) {
	if count > len(content) {
		count = len(content)
	}
	fmt.Printf("==> %s <==\n", filename)
	fmt.Print(string(content[len(content)-count:]))
	fmt.Println()
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 || args[0] != "-c" {
		fmt.Println("Usage: go run . -c <number> <file1> [<file2> ...]")
		os.Exit(1)
	}

	count, err := strconv.Atoi(args[1])
	if err != nil || count < 0 {
		fmt.Println("Invalid count value")
		os.Exit(1)
	}

	files := args[2:]

	if len(files) == 0 {
		fmt.Println("No files provided")
		os.Exit(1)
	}

	exitStatus := 0

	for i, file := range files {
		if i > 0 {
			fmt.Println()
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			printError(err)
			exitStatus = 1
		} else {
			printContent(file, content, count)
		}
	}

	os.Exit(exitStatus)
}
