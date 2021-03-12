package pitch

import (
	"github.com/edipermadi/music/pkg/theory/note"
)

// Pitches bundles a list of pitch
type Pitches []Pitch

// Unique return unique pitches
func (p Pitches) Unique() Pitches {
	pitches := make([]Pitch, 0)
	uniqueMap := make(map[Pitch]struct{})
	for _, v := range p {
		if _, found := uniqueMap[v]; !found {
			uniqueMap[v] = struct{}{}
			pitches = append(pitches, v)
		}
	}

	return pitches
}

// Contains returns true when given pitch is in the slice
func (p Pitches) Contains(v Pitch) bool {
	for _, x := range p {
		if x.Equal(v) {
			return true
		}
	}
	return false
}

// Strings return stringified pitches
func (p Pitches) Strings() []string {
	pitches := make([]string, 0)
	for _, v := range p {
		pitches = append(pitches, v.String())
	}

	return pitches
}

// PitchesFromNotes return pitches form notes
func PitchesFromNotes(givenOctave int, notes ...note.Note) Pitches {
	pitches := make([]Pitch, 0)

	var previousOffset int
	for idx, givenNote := range notes {
		if idx > 0 && givenNote.AbsoluteOffset() <= previousOffset {
			givenOctave++
		}

		pitches = append(pitches, Make(givenNote, givenOctave))
		previousOffset = givenNote.AbsoluteOffset()
	}
	return pitches
}
