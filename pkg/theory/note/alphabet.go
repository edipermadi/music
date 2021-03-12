package note

import "github.com/edipermadi/music/pkg/theory/alphabet"

var noteAlphabets = map[Note]alphabet.Alphabet{
	CTripleFlat: alphabet.C, CDoubleFlat: alphabet.C, CFlat: alphabet.C, CNatural: alphabet.C, CSharp: alphabet.C, CDoubleSharp: alphabet.C, CTripleSharp: alphabet.C,
	DTripleFlat: alphabet.D, DDoubleFlat: alphabet.D, DFlat: alphabet.D, DNatural: alphabet.D, DSharp: alphabet.D, DDoubleSharp: alphabet.D, DTripleSharp: alphabet.D,
	ETripleFlat: alphabet.E, EDoubleFlat: alphabet.E, EFlat: alphabet.E, ENatural: alphabet.E, ESharp: alphabet.E, EDoubleSharp: alphabet.E, ETripleSharp: alphabet.E,
	FTripleFlat: alphabet.F, FDoubleFlat: alphabet.F, FFlat: alphabet.F, FNatural: alphabet.F, FSharp: alphabet.F, FDoubleSharp: alphabet.F, FTripleSharp: alphabet.F,
	GTripleFlat: alphabet.G, GDoubleFlat: alphabet.G, GFlat: alphabet.G, GNatural: alphabet.G, GSharp: alphabet.G, GDoubleSharp: alphabet.G, GTripleSharp: alphabet.G,
	ATripleFlat: alphabet.A, ADoubleFlat: alphabet.A, AFlat: alphabet.A, ANatural: alphabet.A, ASharp: alphabet.A, ADoubleSharp: alphabet.A, ATripleSharp: alphabet.A,
	BTripleFlat: alphabet.B, BDoubleFlat: alphabet.B, BFlat: alphabet.B, BNatural: alphabet.B, BSharp: alphabet.B, BDoubleSharp: alphabet.B, BTripleSharp: alphabet.B,
}

// Alphabet return note alphabet
func (n Note) Alphabet() alphabet.Alphabet {
	return noteAlphabets[n]
}
