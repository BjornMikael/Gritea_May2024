package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || containsHelpFlag(args) {
		printHelp()
		return
	}

	var insertStr string
	var orderFlag bool

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}
		if hasPrefix(arg, "--insert=") || hasPrefix(arg, "-i=") {
			insertStr = trimPrefix(arg, "--insert=")
			insertStr = trimPrefix(insertStr, "-i=")
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		}
	}

	stringArg := getStringArg(args)

	if insertStr != "" {
		stringArg = insertString(stringArg, insertStr)
	}

	if orderFlag {
		stringArg = orderString(stringArg)
	}

	printString(stringArg)
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func getStringArg(args []string) string {
	var stringArg string
	for _, arg := range args {
		if !hasPrefix(arg, "--") && !hasPrefix(arg, "-") {
			stringArg = arg
			break
		}
	}
	return stringArg
}

func insertString(base, insert string) string {
	var result string
	for _, char := range base {
		result += string(char)
	}
	result += insert
	return result
}

func orderString(str string) string {
	chars := []rune(str)
	for i := 0; i < len(chars)-1; i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] > chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return string(chars)
}

func printString(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func trimPrefix(s, prefix string) string {
	if hasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}
