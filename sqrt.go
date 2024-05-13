package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	ret := 1
	for ret*ret < nb && ret*ret >= ret {
		ret++
	}
	if ret*ret == nb {
		return ret
	} else {
		return 0
	}
}
