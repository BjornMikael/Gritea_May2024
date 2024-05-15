package piscine

func Capitalize(s string) string {
	var newString string
	isFirst := true
	for i := 0; i < len(s); i++ {
		if IsAlpha(string(s[i])) {
			if isFirst {
				newString += ToUpper(string(s[i]))
			} else {
				newString += ToLower(string(s[i]))
			}
			isFirst = false
		} else {
			newString += string(s[i])
			isFirst = true
		}
	}
	return newString
}
