package piscine

func UltimateDivMod(a *int, b *int) {
	dm1 := *a / *b
	dm2 := *a % *b
	*a = dm1
	*b = dm2
}
