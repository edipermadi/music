package modetype

// Leap returns largest semitone jump
func (t Type) Leap() int {
	computedScale, _ := t.Scale()
	return computedScale.Leap()
}
