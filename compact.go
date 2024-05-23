package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	compactSlice := make([]string, 0, len(slice))
	count := 0

	for _, v := range slice {
		if v != "" {
			compactSlice = append(compactSlice, v)
			count++
		}
	}

	*ptr = compactSlice
	return count
}
