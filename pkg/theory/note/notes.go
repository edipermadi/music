package note

import (
	"sort"

	"github.com/edipermadi/music/pkg/theory/alphabet"
)

// Notes is a bundle of note
type Notes []Note

func (n Notes) IndexOf(v Note) int {
	for idx, givenNote := range n {
		if givenNote.Equal(v) {
			return idx
		}
	}
	return -1
}

func (n Notes) Sort() Notes {
	sort.SliceStable(n, func(i int, j int) bool {
		return n[i].AbsoluteOffset() < n[j].AbsoluteOffset()
	})
	return n
}

// LeastAccidental find least accidental
func (n Notes) LeastAccidental() Note {
	for _, v := range n {
		if v.Accidental().Natural() {
			return v
		}
	}

	for _, v := range n {
		if v.Accidental().Single() {
			return v
		}
	}

	for _, v := range n {
		if v.Accidental().Double() {
			return v
		}
	}

	return n[0]
}

// MatchAlphabet returns note matching given alphabet
func (n Notes) MatchAlphabet(expectedAlphabet alphabet.Alphabet) Note {
	for _, givenNote := range n {
		if givenNote.Alphabet() == expectedAlphabet {
			return givenNote
		}
	}

	return n[0]
}

// Normalized returns normalized notes
func (n Notes) Normalized() Notes {
	notes := make(Notes, 0)
	for _, v := range n {
		notes = append(notes, v.Normalize())
	}

	return notes
}

// Unique return uniqe notes
func (n Notes) Unique() Notes {
	uniqueNotes := make(map[Note]struct{})

	notes := make(Notes, 0)
	for _, v := range n {
		if _, found := uniqueNotes[v]; !found {
			uniqueNotes[v] = struct{}{}
			notes = append(notes, v)
		}
	}
	return notes
}

// Contains returns true if it contatains given note
func (n Notes) Contains(v Note) bool {
	for _, w := range n {
		if w.Equal(v) {
			return true
		}
	}

	return false
}

// ContainsAll returns true when all notes are in the list
func (n Notes) ContainsAll(notes ...Note) bool {
	mask := Notes(notes).Checksum()
	return n.Checksum()&mask == mask
}

// ContainsAny returns true when one of given notes is in the list
func (n Notes) ContainsAny(notes ...Note) bool {
	if len(notes) == 0 {
		return false
	}

	for _, v := range notes {
		if n.Contains(v) {
			return true
		}
	}

	return false
}

// Names return note names
func (n Notes) Names() []string {
	names := make([]string, 0)
	for _, v := range n {
		names = append(names, v.Name())
	}
	return names
}

// Rotate rotates note slice
func (n Notes) Rotate(k int) Notes {
	if k < 0 {
		panic("invalid rotation")
	}

	r := k % len(n)
	return append(n[k:], n[:r]...)
}

// Checksum returns note slice checksum
func (n Notes) Checksum() int {
	var value int

	for _, givenNote := range n {
		value |= (1 << givenNote.AbsoluteOffset())
	}

	return value
}
