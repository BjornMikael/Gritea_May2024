package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maxVal := a[0]
	for _, val := range a {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
