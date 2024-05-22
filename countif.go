package piscine

// CountIf function returns the number of elements in the slice tab
// for which the function f returns true
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, str := range tab {
		if f(str) {
			count++
		}
	}
	return count
}
