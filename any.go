package piscine

// Any function returns true if the function f returns true for at least one element in the slice a
func Any(f func(string) bool, a []string) bool {
	for _, str := range a {
		if f(str) {
			return true
		}
	}
	return false
}
