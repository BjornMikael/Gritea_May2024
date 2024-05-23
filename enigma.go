package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Temporary storage for the values to swap them around
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	// Move values around as specified.
	*******c = tempA
	****d = tempC
	*b = tempD
	***a = tempB
}
