package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		if n, valid := stringToInt(arg); valid && n >= 1 && n <= 26 {
			if upper {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

// stringToInt converts a string to an integer manually.
// It returns the integer value and a boolean indicating if the conversion was successful.
func stringToInt(s string) (int, bool) {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		n = n*10 + int(r-'0')
	}
	return n, true
}
