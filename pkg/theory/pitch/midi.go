package pitch

// MIDINoteNumber returns MIDI note number
func (p Pitch) MIDINoteNumber() int {
	normalized := p.Normalize()
	return normalized.Note().MIDINoteNumber(normalized.Octave())
}
