package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	digits := "0123456789"
	combination := make([]rune, n)
	printCombinations(digits, combination, 0, 0, n)
	z01.PrintRune('\n')
}

func printCombinations(digits string, combination []rune, start, index, n int) {
	if index == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(combination[i])
		}
		if combination[0] != '9'-rune(n-1) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	for i := start; i < len(digits); i++ {
		combination[index] = rune(digits[i])
		printCombinations(digits, combination, i+1, index+1, n)
	}
}
