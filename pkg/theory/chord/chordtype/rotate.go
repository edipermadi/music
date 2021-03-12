package chordtype

// Rotate rotates chord type slice to the left
func (t Types) Rotate(k int) Types {
	length := len(t)
	if k < 0 || length == 0 {
		return t
	}

	r := k % length
	return append(t[k:], t[:r]...)
}
