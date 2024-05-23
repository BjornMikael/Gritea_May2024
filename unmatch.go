package piscine

func Unmatch(a []int) int {
	result := -1
	for _, num := range a {
		result ^= num
	}
	return result
}
