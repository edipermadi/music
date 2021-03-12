package note

import "github.com/edipermadi/music/pkg/theory/note/accidental"

var noteAccidentals = map[Note]accidental.Accidental{
	CTripleFlat:  accidental.TripleFlat,
	CDoubleFlat:  accidental.DoubleFlat,
	CFlat:        accidental.Flat,
	CNatural:     accidental.Natural,
	CSharp:       accidental.Sharp,
	CDoubleSharp: accidental.DoubleSharp,
	CTripleSharp: accidental.TripleSharp,

	DTripleFlat:  accidental.TripleFlat,
	DDoubleFlat:  accidental.DoubleFlat,
	DFlat:        accidental.Flat,
	DNatural:     accidental.Natural,
	DSharp:       accidental.Sharp,
	DDoubleSharp: accidental.DoubleSharp,
	DTripleSharp: accidental.TripleSharp,

	ETripleFlat:  accidental.TripleFlat,
	EDoubleFlat:  accidental.DoubleFlat,
	EFlat:        accidental.Flat,
	ENatural:     accidental.Natural,
	ESharp:       accidental.Sharp,
	EDoubleSharp: accidental.DoubleSharp,
	ETripleSharp: accidental.TripleSharp,

	FTripleFlat:  accidental.TripleFlat,
	FDoubleFlat:  accidental.DoubleFlat,
	FFlat:        accidental.Flat,
	FNatural:     accidental.Natural,
	FSharp:       accidental.Sharp,
	FDoubleSharp: accidental.DoubleSharp,
	FTripleSharp: accidental.TripleSharp,

	GTripleFlat:  accidental.TripleFlat,
	GDoubleFlat:  accidental.DoubleFlat,
	GFlat:        accidental.Flat,
	GNatural:     accidental.Natural,
	GSharp:       accidental.Sharp,
	GDoubleSharp: accidental.DoubleSharp,
	GTripleSharp: accidental.TripleSharp,

	ATripleFlat:  accidental.TripleFlat,
	ADoubleFlat:  accidental.DoubleFlat,
	AFlat:        accidental.Flat,
	ANatural:     accidental.Natural,
	ASharp:       accidental.Sharp,
	ADoubleSharp: accidental.DoubleSharp,
	ATripleSharp: accidental.TripleSharp,

	BTripleFlat:  accidental.TripleFlat,
	BDoubleFlat:  accidental.DoubleFlat,
	BFlat:        accidental.Flat,
	BNatural:     accidental.Natural,
	BSharp:       accidental.Sharp,
	BDoubleSharp: accidental.DoubleSharp,
	BTripleSharp: accidental.TripleSharp,
}

// Accidental returns the accidental of a note
func (n Note) Accidental() accidental.Accidental {
	return noteAccidentals[n]
}
