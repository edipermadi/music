package mode

import (
	"fmt"
	"sort"

	"github.com/edipermadi/music/pkg/theory/mode/modetype"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/edipermadi/music/pkg/theory/scale"
)

// Mode is musical mode structure
type Mode struct {
	modeTonic       note.Note
	modeType        modetype.Type
	number          int
	canonicalNumber int
	label           string
	notes           note.Notes
}

// Modes is a slice of mode
type Modes []Mode

// New return new musical mode
func New(givenTonic note.Note, givenType modetype.Type) Mode {
	computedNotes := computeModeNotes(givenTonic, givenType)

	number := computedNotes.Checksum()
	canonicalNumber := number
	if givenTonic != note.CNatural {
		canonicalNumber = computeModeNotes(note.CNatural, givenType).Checksum()
	}

	return Mode{
		modeTonic:       givenTonic,
		modeType:        givenType,
		number:          number,
		canonicalNumber: canonicalNumber,
		label:           fmt.Sprintf("%s%s", givenTonic, givenType),
		notes:           computedNotes,
	}
}

// String return stringified mode
func (m Mode) String() string {
	return m.label
}

// GoString satisfies GoStringer
func (m Mode) GoString() string {
	return m.String()
}

// Notes returns notes
func (m Mode) Notes() note.Notes {
	return m.notes
}

// Tonic returns mode tonic
func (m Mode) Tonic() note.Note {
	return m.modeTonic
}

// Type returns mode type
func (m Mode) Type() modetype.Type {
	return m.modeType
}

// Number return mode number
func (m Mode) Number() int {
	return m.number
}

// CanonicalNumber return mode canonical number
func (m Mode) CanonicalNumber() int {
	return m.canonicalNumber
}

// Cardinality mode cardinality
func (m Mode) Cardinality() int {
	return m.modeType.Cardinality()
}

// Perfection return mode number
func (m Mode) Perfection() (int, int, []bool) {
	return m.modeType.Perfection()
}

// Transposition returns mode transposition
func (m Mode) Transposition() []int {
	return m.modeType.Transposition()
}

// Scale return scale
func (m Mode) Scale() scale.Scale {
	computedScale, _ := m.modeType.Scale()
	return computedScale
}

// Luminosity returns mode luminosity, -1 when tonic does not have sibling fifth
func (m Mode) Luminosity() int {
	for idx, currentNote := range m.Notes().CircleOfFifth().Notes() {
		if currentNote.Equal(m.Tonic()) {
			return m.Cardinality() - idx
		}
	}

	return -1
}

// AllModes returns all modes
func AllModes() Modes {
	return allModes
}

// Unique return unique modes
func (m Modes) Unique() Modes {
	uniqueModeMap := make(map[string]struct{})

	uniqueMode := make(Modes, 0)
	for _, givenMode := range m {
		modeName := givenMode.String()
		if _, found := uniqueModeMap[modeName]; !found {
			uniqueModeMap[modeName] = struct{}{}
			uniqueMode = append(uniqueMode, givenMode)
		}
	}

	return uniqueMode
}

// Sort return sorted modes
func (m Modes) Sort() Modes {
	sort.SliceStable(m, func(i, j int) bool {
		return m[i].Number() < m[j].Number()
	})

	return m
}
