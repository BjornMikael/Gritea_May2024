package piscine

func AppendRange(min, max int) []int {
	var ret []int
	for ; min < max; min++ {
		ret = append(ret, min)
	}
	return ret
}
