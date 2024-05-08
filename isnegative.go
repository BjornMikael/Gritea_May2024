package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {

	trueval := 'T'
	falsval := 'F'
	if nb >= 0 {
		z01.PrintRune(falsval)
		z01.PrintRune(10)
	} else {
		z01.PrintRune(trueval)
		z01.PrintRune(10)
	}

}
