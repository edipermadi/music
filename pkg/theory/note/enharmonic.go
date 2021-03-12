package note

import "fmt"

// Enharmonic represents enharmonic note information
type Enharmonic struct {
	Note        Note
	OctaveShift int
}

// Enharmonics bundles enharmonic slice
type Enharmonics []Enharmonic

// Notes returns notes from enharmonic slice
func (e Enharmonics) Notes() Notes {
	computedNotes := make(Notes, 0)
	for _, givenEnharmonic := range e {
		computedNotes = append(computedNotes, givenEnharmonic.Note)
	}
	return computedNotes
}

// Enharmonic returns enharmonic notes
func (n Note) Enharmonic() Enharmonics {
	if v, found := mapNoteToEnharmonics[n]; found {
		return v
	}

	panic(fmt.Errorf("no such enharmonics for note %s", n))
}
