package note

// Note represents musical note
type Note int

// Note enumerations
const (
	Invalid Note = iota

	CTripleFlat  Note = iota
	CDoubleFlat  Note = iota
	CFlat        Note = iota
	CNatural     Note = iota
	CSharp       Note = iota
	CDoubleSharp Note = iota
	CTripleSharp Note = iota

	DTripleFlat  Note = iota
	DDoubleFlat  Note = iota
	DFlat        Note = iota
	DNatural     Note = iota
	DSharp       Note = iota
	DDoubleSharp Note = iota
	DTripleSharp Note = iota

	ETripleFlat  Note = iota
	EDoubleFlat  Note = iota
	EFlat        Note = iota
	ENatural     Note = iota
	ESharp       Note = iota
	EDoubleSharp Note = iota
	ETripleSharp Note = iota

	FTripleFlat  Note = iota
	FDoubleFlat  Note = iota
	FFlat        Note = iota
	FNatural     Note = iota
	FSharp       Note = iota
	FDoubleSharp Note = iota
	FTripleSharp Note = iota

	GTripleFlat  Note = iota
	GDoubleFlat  Note = iota
	GFlat        Note = iota
	GNatural     Note = iota
	GSharp       Note = iota
	GDoubleSharp Note = iota
	GTripleSharp Note = iota

	ATripleFlat  Note = iota
	ADoubleFlat  Note = iota
	AFlat        Note = iota
	ANatural     Note = iota
	ASharp       Note = iota
	ADoubleSharp Note = iota
	ATripleSharp Note = iota

	BTripleFlat  Note = iota
	BDoubleFlat  Note = iota
	BFlat        Note = iota
	BNatural     Note = iota
	BSharp       Note = iota
	BDoubleSharp Note = iota
	BTripleSharp Note = iota
)

// AllNotes return a slice of notes
func AllNotes() Notes {
	return allNotes
}

// AllSimpleNotes return simple notes
func AllSimpleNotes() Notes {
	return Notes{CNatural, CSharp, DFlat, DNatural, DSharp, EFlat, ENatural, FNatural, FSharp, GFlat, GNatural, GSharp, AFlat, ANatural, ASharp, BFlat, BNatural}
}
