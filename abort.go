package piscine

import "sort"

// Abort returns the median of five int arguments.
func Abort(a, b, c, d, e int) int {
	// Create a slice with the five integers
	numbers := []int{a, b, c, d, e}

	// Sort the slice
	sort.Ints(numbers)

	// Return the middle element
	return numbers[2]
}
