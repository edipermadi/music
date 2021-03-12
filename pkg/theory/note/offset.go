package note

import (
	"github.com/edipermadi/music/pkg/theory/limit"
)

var noteOffsets = map[Note]int{
	CTripleFlat: -3, CDoubleFlat: -2, CFlat: -1, CNatural: 0, CSharp: 1, CDoubleSharp: 2, CTripleSharp: 3,
	DTripleFlat: -1, DDoubleFlat: 0, DFlat: 1, DNatural: 2, DSharp: 3, DDoubleSharp: 4, DTripleSharp: 5,
	ETripleFlat: 1, EDoubleFlat: 2, EFlat: 3, ENatural: 4, ESharp: 5, EDoubleSharp: 6, ETripleSharp: 7,
	FTripleFlat: 2, FDoubleFlat: 3, FFlat: 4, FNatural: 5, FSharp: 6, FDoubleSharp: 7, FTripleSharp: 8,
	GTripleFlat: 4, GDoubleFlat: 5, GFlat: 6, GNatural: 7, GSharp: 8, GDoubleSharp: 9, GTripleSharp: 10,
	ATripleFlat: 6, ADoubleFlat: 7, AFlat: 8, ANatural: 9, ASharp: 10, ADoubleSharp: 11, ATripleSharp: 12,
	BTripleFlat: 8, BDoubleFlat: 9, BFlat: 10, BNatural: 11, BSharp: 12, BDoubleSharp: 13, BTripleSharp: 14,
}

// AbsoluteOffset return pitch absolute offset
func (n Note) AbsoluteOffset() int {
	return ((noteOffsets[n] % limit.SemitonesInOneOctave) + limit.SemitonesInOneOctave) % limit.SemitonesInOneOctave
}

func (n Note) octaveShift() int {
	off := noteOffsets[n]
	switch {
	case off < 0:
		return -1
	case off > 11:
		return 1
	default:
		return 0
	}
}
