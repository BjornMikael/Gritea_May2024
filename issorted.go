package piscine

// IsSorted function checks if the slice a is sorted according to the comparison function f
func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n <= 1 {
		return true
	}

	ascending := true
	descending := true

	for i := 1; i < n; i++ {
		if f(a[i-1], a[i]) > 0 {
			ascending = false
		}
		if f(a[i-1], a[i]) < 0 {
			descending = false
		}
	}
	return ascending || descending
}
