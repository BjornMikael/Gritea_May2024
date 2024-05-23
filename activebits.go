package piscine

// ActiveBits returns the number of active bits (bits with the value 1) in the binary representation of an integer n.
func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1 // Increment count if the last bit is 1
		n >>= 1        // Right shift the bits of n
	}
	return count
}
