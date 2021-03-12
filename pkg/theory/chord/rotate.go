package chord

// Rotate rotates chord slice to the left
func (c Chords) Rotate(k int) Chords {
	length := len(c)
	if k < 0 || length == 0 {
		return c
	}

	r := k % length
	return append(c[k:], c[:r]...)
}
