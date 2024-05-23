package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			n = -n
		}
	}

	printDigits(n)
}

func printDigits(n int) {
	if n == 0 {
		return
	}
	printDigits(n / 10)
	z01.PrintRune(rune(n%10 + '0'))
}
