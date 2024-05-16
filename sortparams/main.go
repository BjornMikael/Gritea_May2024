package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s []rune) {
	n := 0
	for n < len(s) {
		z01.PrintRune(s[n])
		n++
	}
}

func endIndex(s []rune, toFind rune) int {
	ret := 0
	for nextIndex := 0; nextIndex < len(s); nextIndex++ {
		if s[nextIndex] == toFind {
			ret = nextIndex + 1
		}
	}
	return ret
}

func isSorted(args []string) bool {
	for i := 1; i < len(args)-1; i++ {
		if args[i] > args[i+1] {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args
	for !isSorted(args) {
		for i := 1; i < len(args)-1; i++ {
			if args[i] > args[i+1] {
				tmp := args[i]
				args[i] = args[i+1]
				args[i+1] = tmp
			}
		}
	}
	for i := 1; i < len(args); i++ {
		arg := []rune(args[i])
		printStr(arg[endIndex(arg, '/'):])
		z01.PrintRune('\n')
	}
}
