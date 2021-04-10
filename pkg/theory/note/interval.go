package note

import (
	"fmt"

	"github.com/edipermadi/music/pkg/theory/interval"
)

// Interval returns interval between two notes
func (n Note) Interval(v Note) ([]interval.Interval, error) {
	v1 := n.MIDINoteNumber(4)

	var v2 int
	if n.AbsoluteOffset() <= v.AbsoluteOffset() {
		v2 = v.MIDINoteNumber(4)
	} else {
		v2 = v.MIDINoteNumber(5)
	}

	return interval.FromSemitones(v2 - v1)
}

// NextFifth returns next fifth note
func (n Note) NextFifth() Note {
	normalized := n.Normalize()
	switch normalized {
	case CNatural:
		return GNatural
	case GNatural:
		return DNatural
	case DNatural:
		return ANatural
	case ANatural:
		return ENatural
	case ENatural:
		return BNatural
	case BNatural:
		return GFlat
	case FSharp, GFlat:
		return DFlat
	case CSharp, DFlat:
		return AFlat
	case GSharp, AFlat:
		return EFlat
	case DSharp, EFlat:
		return BFlat
	case ASharp, BFlat:
		return FNatural
	case FNatural:
		return CNatural
	default:
		panic(fmt.Errorf("invalid note %s", normalized))
	}
}

// PreviousFifth returns previous fifth note
func (n Note) PreviousFifth() Note {
	normalized := n.Normalize()
	switch normalized {
	case CNatural:
		return FNatural
	case GNatural:
		return CNatural
	case DNatural:
		return GNatural
	case ANatural:
		return DNatural
	case ENatural:
		return ANatural
	case BNatural:
		return ENatural
	case FSharp, GFlat:
		return BNatural
	case CSharp, DFlat:
		return GFlat
	case GSharp, AFlat:
		return DFlat
	case DSharp, EFlat:
		return AFlat
	case ASharp, BFlat:
		return EFlat
	case FNatural:
		return BFlat
	default:
		panic(fmt.Errorf("invalid note %s", normalized))
	}
}
