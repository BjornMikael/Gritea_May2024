package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || 'z' < s[i] {
			return false
		}
	}
	return true
}
