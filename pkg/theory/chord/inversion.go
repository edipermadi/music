package chord

import "fmt"

// Invert returns chord notes with inversion
func (c Chord) Invert() Chord {
	max := len(c.notes)
	computedNotes := c.notes.Rotate(1)
	return Chord{
		inversion:       (c.inversion + 1) % max,
		chordRoot:       c.chordRoot,
		chordType:       c.chordType,
		notes:           computedNotes,
		number:          computedNotes.Checksum(),
		canonicalNumber: c.canonicalNumber,
	}
}

// Inversion returns chord notes with inversion
func (c Chord) Inversion(pos int) Chord {
	pos = pos % len(c.notes)
	normalized := New(c.chordRoot, c.chordType)
	normalized.inversion = pos
	normalized.notes = normalized.notes.Rotate(pos)
	normalized.number = normalized.notes.Checksum()
	normalized.canonicalNumber = c.canonicalNumber
	return normalized
}

// InversionNumber return inversion number
func (c Chord) InversionNumber() int {
	return c.inversion
}

// InversionName return inversion name
func (c Chord) InversionName() string {
	switch c.inversion {
	case 0:
		return "RootPosition"
	case 1:
		return "FirstInversion"
	case 2:
		return "SecondInversion"
	case 3:
		return "ThirdInversion"
	case 4:
		return "FourthInversion"
	case 5:
		return "FifthInversion"
	case 6:
		return "SixthInversion"
	}

	panic(fmt.Errorf("no such inversion name for number %d", c.inversion))
}
