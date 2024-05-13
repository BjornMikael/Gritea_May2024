package piscine

func FindNextPrime(nb int) int {
	ret := nb
	if ret <= 0 {
		ret = 2
	}
	for !IsPrime(ret) {
		ret++
	}
	return ret
}
