package piscine

import "github.com/01-edu/z01"

func printStr(s string) {
	n := 0
	for n < len(s) {
		z01.PrintRune(rune(s[n]))
		n++
	}
}

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		printStr(a[i])
		z01.PrintRune('\n')
	}
}
