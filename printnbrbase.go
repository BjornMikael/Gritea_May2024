package piscine

import (
	"github.com/01-edu/z01"
)

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
		if base[i] == '-' || base[i] == '+' {
			return false
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	lenBase := len(base)
	if !isValidBase(base) {
		printString("NV")
		return
	}
	var nb string
	var nbr64 uint64
	nbr64 = uint64(nbr)
	if int64(nbr64) < 0 {
		nbr64 = uint64(int64(nbr64) * -1)
		z01.PrintRune('-')
	}
	if nbr64 == 0 {
		nb += string(base[0])
	} else {
		for nbr64 > 0 {
			nb += string(base[nbr64%uint64(lenBase)])
			nbr64 /= uint64(lenBase)
		}
	}
	for i := len(nb) - 1; i >= 0; i-- {
		z01.PrintRune(rune(nb[i]))
	}
}

func printString(s string) {
	panic("unimplemented")
}
