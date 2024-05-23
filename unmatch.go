package piscine

func Unmatch(a []int) int {
	var result uint
	for _, num := range a {
		result ^= uint(num)
	}
	return int(result)
}
