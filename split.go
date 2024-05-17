package piscine

func isSep(s, sep string) bool {
	if len(s) >= len(sep) && s[:len(sep)] == sep {
		return true
	} else {
		return false
	}
}

func Split(s, sep string) []string {
	var ret []string
	for i := 0; i < len(s); {
		if !isSep(s[i:], sep) {
			j := 0
			for i+j < len(s) && !isSep(s[i+j:], sep) {
				j++
			}
			ret = append(ret, s[i:i+j])
			i += j
		} else {
			i += len(sep)
		}
	}
	return ret
}
