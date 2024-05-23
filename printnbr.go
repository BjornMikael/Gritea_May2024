package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -2147483648 {
			z01.PrintRune('2')
			n = 147483648
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
