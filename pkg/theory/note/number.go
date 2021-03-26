package note

// Number returns pitch class number, CNatural will be 0, CSharp will be 1 and so on
func (n Note) Number() int {
	switch n {
	case CNatural, DDoubleFlat, ATripleSharp, BSharp:
		return 0
	case CSharp, DFlat, ETripleFlat, BDoubleSharp:
		return 1
	case DNatural, CDoubleSharp, EDoubleFlat, FTripleFlat, BTripleSharp:
		return 2
	case DSharp, CTripleSharp, EFlat, FDoubleFlat:
		return 3
	case ENatural, DDoubleSharp, FFlat, GTripleFlat:
		return 4
	case FNatural, DTripleSharp, ESharp, GDoubleFlat:
		return 5
	case FSharp, EDoubleSharp, GFlat, ATripleFlat:
		return 6
	case GNatural, ETripleSharp, FDoubleSharp, ADoubleFlat:
		return 7
	case GSharp, FTripleSharp, AFlat, BTripleFlat:
		return 8
	case ANatural, CTripleFlat, GDoubleSharp, BDoubleFlat:
		return 9
	case ASharp, CDoubleFlat, GTripleSharp, BFlat:
		return 10
	case BNatural, CFlat, DTripleFlat, ADoubleSharp:
		return 11
	default:
		panic("invalid note")
	}
}
