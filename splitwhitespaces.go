package piscine

func isWhiteSpaces(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\v' || b == '\f' || b == '\r'
}

func SplitWhiteSpaces(s string) []string {
	var ret []string
	for i := 0; i < len(s); i++ {
		if !isWhiteSpaces(s[i]) {
			j := 0
			for i+j < len(s) && !isWhiteSpaces(s[i+j]) {
				j++
			}
			ret = append(ret, s[i:i+j])
			i = i + j
		}
	}
	return ret
}
