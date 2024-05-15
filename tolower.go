package piscine

func ToLower(s string) string {
	var newString string
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			newString += string(s[i] - 'A' + 'a')
		} else {
			newString += string(s[i])
		}
	}
	return newString
}
