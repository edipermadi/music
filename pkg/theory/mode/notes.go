package mode

import (
	"github.com/edipermadi/music/pkg/theory/mode/modetype"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/edipermadi/music/pkg/theory/note/accidental"
)

func computeModeNotes(givenTonic note.Note, givenType modetype.Type) note.Notes {
	templateNotesWithSharp := note.Notes{note.CNatural, note.CSharp, note.DNatural, note.DSharp, note.ENatural, note.FNatural, note.FSharp, note.GNatural, note.GSharp, note.ANatural, note.ASharp, note.BNatural}
	templateNotesWithFlat := note.Notes{note.CNatural, note.DFlat, note.DNatural, note.EFlat, note.ENatural, note.FNatural, note.GFlat, note.GNatural, note.AFlat, note.ANatural, note.BFlat, note.BNatural}

	var templateNotes note.Notes

	givenTonic = givenTonic.Normalize()
	if givenTonic.Accidental() == accidental.Flat {
		templateNotes = templateNotesWithFlat
	} else {
		templateNotes = templateNotesWithSharp
	}

	var currentIndex int
	computedNotes := make(note.Notes, 0)
	for idx, givenNote := range templateNotes {
		if givenNote == givenTonic {
			currentIndex = idx
			computedNotes = append(computedNotes, givenNote)
			break
		}
	}

	previousNote := givenTonic
	computedTransposition := givenType.Transposition()
	for _, intervalSemitone := range computedTransposition[:len(computedTransposition)-1] {
		currentIndex = (currentIndex + intervalSemitone) % len(templateNotes)
		currentNote := templateNotes[currentIndex]
		nextAlphabet := previousNote.Alphabet().Next()
		computedCardinality := givenType.Cardinality()
		canBeUnique := (computedCardinality == 6) || (computedCardinality == 7)
		if currentNote.Alphabet() == nextAlphabet || !canBeUnique {
			previousNote = currentNote
		} else {
			previousNote = currentNote.Enharmonic().Notes().MatchAlphabet(nextAlphabet)
		}
		computedNotes = append(computedNotes, previousNote)
	}

	return append(computedNotes, givenTonic)
}
