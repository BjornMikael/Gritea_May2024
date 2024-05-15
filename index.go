package piscine

func Index(s, toFind string) int {
	lenS := len(s)
	lenF := len(toFind)
	for i := range s {
		if i+lenF > lenS {
			return -1
		}
		if s[i:i+lenF] == toFind {
			return i
		}
	}
	return -1
}
