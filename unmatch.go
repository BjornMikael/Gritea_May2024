// unmatch.go
package piscine

func Unmatch(a []int) int {
	countMap := make(map[int]int)

	// Count occurrences of each element
	for _, num := range a {
		countMap[num]++
	}

	// Find the element with an odd count
	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}

	// If all elements have pairs, return -1
	return -1
}
