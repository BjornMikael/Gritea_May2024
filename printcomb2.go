package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			printTwoDigitNumber(i)
			z01.PrintRune(' ')
			printTwoDigitNumber(j)
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigitNumber(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}
