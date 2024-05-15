package piscine

func Join(strs []string, sep string) string {
	var newString string
	for i := 0; i < len(strs); i++ {
		if i > 0 {
			newString += sep
		}
		newString += strs[i]
	}
	return newString
}
