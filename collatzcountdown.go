package piscine

// CollatzCountdown returns the number of steps to reach 1 using the Collatz sequence.
// It returns -1 if happdthen the start is zero or negative.
func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	steps := 0
	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		steps++
	}

	return steps
}
