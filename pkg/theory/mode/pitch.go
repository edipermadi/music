package mode

import (
	"github.com/edipermadi/music/pkg/theory/pitch"
)

// Pitches build pitches
func (m Mode) Pitches(givenOctave int) pitch.Pitches {
	var pitches pitch.Pitches

	var previousOffset int
	for idx, givenNote := range m.notes {
		if idx > 0 && previousOffset >= givenNote.AbsoluteOffset() {
			givenOctave++
		}

		pitches = append(pitches, pitch.Make(givenNote, givenOctave))
		previousOffset = givenNote.AbsoluteOffset()
	}

	return pitches
}
