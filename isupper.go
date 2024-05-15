package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || 'Z' < s[i] {
			return false
		}
	}
	return true
}
