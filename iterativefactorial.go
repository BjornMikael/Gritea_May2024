package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	ret := 1
	for ; nb > 1; nb-- {
		if ret*nb < ret {
			return 0
		}
		ret *= nb
	}
	return ret
}
