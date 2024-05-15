package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'A' || 'Z' < s[i]) && (s[i] < 'a' || 'z' < s[i]) && (s[i] < '0' || '9' < s[i]) {
			return false
		}
	}
	return true
}
