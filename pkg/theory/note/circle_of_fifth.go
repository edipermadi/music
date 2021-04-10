package note

// CircleOfFifth represents circle of fifth, true when the note is used by a scale
type CircleOfFifth map[Note]bool

var circleOfFifthNotes = []Note{
	CNatural, GNatural, DNatural, ANatural, ENatural, BNatural, GFlat, DFlat, AFlat, EFlat, BFlat, FNatural,
}

// CircleOfFifth returns circles of fifth from notes
func (n Notes) CircleOfFifth() CircleOfFifth {
	circle := make(map[Note]bool)
	for _, k := range circleOfFifthNotes {
		circle[k] = n.Contains(k)
	}
	return circle
}

// Contains returns true when note is active in circle of fifth
func (c CircleOfFifth) Contains(givenNote Note) bool {
	for currentNote := range c {
		if currentNote.Equal(givenNote) {
			return true
		}
	}

	return false
}

// HasLeftSibling returns true when previous fifth is in the circle of fifth
func (c CircleOfFifth) HasLeftSibling(givenNote Note) bool {
	return c.Contains(givenNote.PreviousFifth())
}

// HasRightSibling returns true when next fifth is in the circle of fifth
func (c CircleOfFifth) HasRightSibling(givenNote Note) bool {
	return c.Contains(givenNote.NextFifth())
}

// HasSibling returns true when next or previous fifth is in the circle of fifth
func (c CircleOfFifth) HasSibling(givenNote Note) bool {
	return c.HasLeftSibling(givenNote) || c.HasRightSibling(givenNote)
}

// Leftmost returns the left most note within the circle of fifth
func (c CircleOfFifth) Leftmost() Note {
	for currentNote, active := range c {
		if !active || !c.HasSibling(currentNote) {
			continue
		} else if !c.HasLeftSibling(currentNote) {
			return currentNote
		}
	}

	return Invalid
}

// Rightmost returns the right most note within the circle of fifth
func (c CircleOfFifth) Rightmost() Note {
	for currentNote, active := range c {
		if !active || !c.HasSibling(currentNote) {
			continue
		} else if !c.HasRightSibling(currentNote) {
			return currentNote
		}
	}

	return Invalid
}

// Notes return notes sorted clockwise
func (c CircleOfFifth) Notes() Notes {
	notes := make(Notes, 0)
	for currentNote := range c {
		if c.HasSibling(currentNote) {
			notes = append(notes, currentNote)
		}
	}

	lastNote := Invalid
	sortedNotes := make(Notes, 0)
	if len(notes) == 0 {
		return sortedNotes
	}

	for count := 0; count < len(notes); {
		for _, currentNote := range notes {
			if lastNote.Equal(Invalid) && c.Leftmost().Equal(currentNote) {
				sortedNotes = append(sortedNotes, currentNote)
				lastNote = currentNote
				count++
			} else if !lastNote.Equal(Invalid) && lastNote.NextFifth().Equal(currentNote) {
				sortedNotes = append(sortedNotes, currentNote)
				lastNote = currentNote
				count++
			}
		}
	}

	return sortedNotes
}
