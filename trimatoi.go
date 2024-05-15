package piscine

func TrimAtoi(s string) int {
	trimAtoi := 0 // trim A to i
	isNeg := false
	index := 0
	for index < len(s) && (s[index] < '0' || s[index] > '9') {
		if s[index] == '-' {
			isNeg = true
		}
		index++
	}
	for index < len(s) {
		if '0' <= s[index] && s[index] <= '9' {
			trimAtoi *= 10
			trimAtoi += int(s[index] - '0')
		}
		index++
	}
	if isNeg {
		trimAtoi *= -1
	}
	return trimAtoi
}
