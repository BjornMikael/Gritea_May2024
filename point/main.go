package main

import (
	"github.com/01-edu/z01"
)

// Define the point struct
type point struct {
	x int
	y int
}

// setPoint function sets the values of x and y
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// printNumber prints an integer using z01.PrintRune
func printNumber(n int) {
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
		digits = append(digits, rune('0'+(n%10)))
		n = n / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

// printString prints a string using z01.PrintRune
func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printString("x = ")
	printNumber(points.x)
	printString(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}
