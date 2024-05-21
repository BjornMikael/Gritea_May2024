package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printError(err error) {
	errorMessage := "ERROR: " + err.Error() + "\n"
	for _, r := range errorMessage {
		z01.PrintRune(r)
	}
}

func printContent(content string) {
	for _, r := range content {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Read from standard input if no arguments are provided
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
	} else {
		for _, fileName := range args {
			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			printContent(string(content))
		}
	}
}
