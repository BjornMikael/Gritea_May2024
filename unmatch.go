package piscine

func Unmatch(a []int) int {
	seen := make(map[int]bool)

	for _, num := range a {
		if seen[num] {
			delete(seen, num)
		} else {
			seen[num] = true
		}
	}

	for num := range seen {
		return num
	}

	return -1
}
