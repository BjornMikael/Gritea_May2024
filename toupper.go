package piscine

func ToUpper(s string) string {
	var newString string
	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && s[i] <= 'z' {
			newString += string(s[i] - 'a' + 'A')
		} else {
			newString += string(s[i])
		}
	}
	return newString
}
