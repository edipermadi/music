package accidental

import "fmt"

// Accidental represents note accidental type
type Accidental int

// Accidental enumeration
const (
	Invalid     Accidental = iota
	TripleFlat  Accidental = iota
	DoubleFlat  Accidental = iota
	Flat        Accidental = iota
	Natural     Accidental = iota
	Sharp       Accidental = iota
	DoubleSharp Accidental = iota
	TripleSharp Accidental = iota
)

// String returns stringified accidental
func (a Accidental) String() string {
	switch a {
	case TripleFlat:
		return "TripleFlat"
	case DoubleFlat:
		return "DoubleFlat"
	case Flat:
		return "Flat"
	case Natural:
		return "Natural"
	case Sharp:
		return "Sharp"
	case DoubleSharp:
		return "DoubleSharp"
	case TripleSharp:
		return "TripleSharp"
	default:
		return "Invalid"
	}
}

// GoString satisfies GoStringer
func (a Accidental) GoString() string {
	return a.String()
}

// Name returns accidental name
func (a Accidental) Name() string {
	switch a {
	case TripleFlat:
		return "bbb"
	case DoubleFlat:
		return "bb"
	case Flat:
		return "b"
	case Natural:
		return ""
	case Sharp:
		return "#"
	case DoubleSharp:
		return "##"
	case TripleSharp:
		return "###"
	default:
		return "Invalid"
	}
}

// Natural returns true when there is no accidental
func (a Accidental) Natural() bool {
	return a == Natural
}

// Single returns true when there is single accidental
func (a Accidental) Single() bool {
	return (a == Flat) || (a == Sharp)
}

// Double returns true when there is double accidental
func (a Accidental) Double() bool {
	return (a == DoubleFlat) || (a == DoubleSharp)
}

// Triple returns true when there is triple accidental
func (a Accidental) Triple() bool {
	return (a == TripleFlat) || (a == TripleSharp)
}

// Displacement returns semitone displacement
func (a Accidental) Displacement() int {
	switch a {
	case TripleFlat:
		return -3
	case DoubleFlat:
		return -2
	case Flat:
		return -1
	case Natural:
		return 0
	case Sharp:
		return 1
	case DoubleSharp:
		return 2
	case TripleSharp:
		return 3
	default:
		panic(fmt.Errorf("no such displacement for %v", a.Name()))
	}
}

// AllAccidentals return all accidentals
func AllAccidentals() []Accidental {
	return []Accidental{TripleFlat, DoubleFlat, Flat, Natural, Sharp, DoubleSharp, TripleSharp}
}
