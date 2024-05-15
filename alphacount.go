package piscine

func AlphaCount(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
			ret++
		}
	}
	return ret
}
