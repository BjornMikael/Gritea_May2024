package piscine

// ListSize returns the number of elements in the list
func ListSize(l *List) int {
	n := l.Head
	size := 0
	for n != nil {
		size++
		n = n.Next
	}
	return size
}
