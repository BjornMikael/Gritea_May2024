package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	ret := make([]int, max-min)
	for i := 0; min+i < max; i++ {
		ret[i] = min + i
	}
	return ret
}
