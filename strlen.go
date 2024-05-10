package piscine

func StrLen(str string) int {
	ar := []rune(str)
	var ynb int
	for i := range ar {
		ynb = i + 1
	}
	return ynb
}
