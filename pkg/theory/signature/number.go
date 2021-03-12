package signature

import (
	"github.com/edipermadi/music/pkg/theory/note"
)

// Number returns signature number
func (s Signature) Number() int {
	return Number(s.Notes())
}

// Number returns signature number from notes
func Number(givenNotes note.Notes) int {
	var computedNumber int
	for _, givenNote := range givenNotes.Unique() {
		givenNote = givenNote.Normalize()
		computedNumber = computedNumber + (1 << givenNote.AbsoluteOffset())
	}

	return computedNumber
}

// FromNotes deduces signature from note
func FromNotes(notes note.Notes) Signatures {
	if v, found := mapNumberToSignature[Number(notes)]; found {
		return v
	}
	return Signatures{CNaturalMajor}
}
