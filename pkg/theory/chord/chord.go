package chord

import (
	"fmt"
	"sort"

	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/edipermadi/music/pkg/theory/note"
)

// Chord chord represents chord type
type Chord struct {
	chordRoot       note.Note
	chordType       chordtype.Type
	inversion       int
	number          int
	canonicalNumber int
	label           string
	notes           note.Notes
}

// New returns a new chord
func New(givenRoot note.Note, givenType chordtype.Type) Chord {
	givenRoot = givenRoot.Normalize()
	computedNotes, found := mapRootToMapTypeToNotes[givenRoot][givenType]
	if !found {
		panic(fmt.Errorf("no such chord notes for root %s and type %s", givenRoot, givenType))
	}

	// compute checksum
	number := computedNotes.Checksum()
	canonicalNumber := number
	if givenRoot != note.CNatural {
		computedCanonicalNotes, found := mapRootToMapTypeToNotes[note.CNatural][givenType]
		if !found {
			panic(fmt.Errorf("no such chord notes for root %s and type %s", givenRoot, givenType))
		}

		canonicalNumber = computedCanonicalNotes.Checksum()
	}

	return Chord{
		inversion:       0,
		chordRoot:       givenRoot,
		chordType:       givenType,
		notes:           computedNotes,
		number:          number,
		label:           fmt.Sprintf("%s%s", givenRoot, givenType),
		canonicalNumber: canonicalNumber,
	}
}

// Chords bundles chord
type Chords []Chord

// Unique returns unique chord slice
func (c Chords) Unique() Chords {
	uniqueChords := make(Chords, 0)
	uniqueChordMap := make(map[int]struct{})
	for _, givenChord := range c {
		if _, found := uniqueChordMap[givenChord.number]; !found {
			uniqueChordMap[givenChord.number] = struct{}{}
			uniqueChords = append(uniqueChords, givenChord)
		}
	}

	return uniqueChords
}

// Sort sorts chord slice
func (c Chords) Sort() Chords {
	sort.SliceStable(c, func(i, j int) bool {
		return c[i].canonicalNumber < c[j].canonicalNumber
	})
	return c
}

// String returns stringified chord
func (c Chord) String() string {
	return c.label
}

// GoString satisfies GoStringer
func (c Chord) GoString() string {
	return c.String()
}

// Name returns chord name
func (c Chord) Name() string {
	if c.chordType == chordtype.Invalid || c.chordRoot == note.Invalid {
		return "Invalid"
	}

	return fmt.Sprintf("%s%s", c.Root().Name(), c.Type().Name())
}

// Root return chord root
func (c Chord) Root() note.Note {
	return c.chordRoot
}

// Type return chord type
func (c Chord) Type() chordtype.Type {
	return c.chordType
}

// AllChords returs all chords
func AllChords() []Chord {
	return allChords
}

// Notes return chord
func (c Chord) Notes() note.Notes {
	return c.notes
}

// Number returns chord number
func (c Chord) Number() int {
	return c.number
}

// CanonicalNumber returns canonical chord number
func (c Chord) CanonicalNumber() int {
	return c.canonicalNumber
}
