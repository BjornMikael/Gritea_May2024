package piscine

import "github.com/01-edu/z01"

// ForEach applies a function f to each element of the slice a.
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// PrintNbr prints an integer using z01.PrintRune.
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}
