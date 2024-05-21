package piscine

import "fmt"

// ForEach applies a function f to each element of the slice a.
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// PrintNbr prints an integer without a newline.
func PrintNbr(n int) {
	fmt.Print(n)
}
