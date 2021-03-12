package pitch

// Equal returns true when equal
func (p Pitch) Equal(v Pitch) bool {
	return p.MIDINoteNumber() == v.MIDINoteNumber()
}
