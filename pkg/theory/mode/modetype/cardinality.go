package modetype

// Cardinality returns count of note in one octave
func (t Type) Cardinality() int {
	computedScale, _ := t.Scale()
	return computedScale.Cardinality()
}
