package note

var normalizedNotes = map[Note]Note{
	CTripleFlat:  ANatural,
	CDoubleFlat:  BFlat,
	CFlat:        BNatural,
	CDoubleSharp: DNatural,
	CTripleSharp: DSharp,

	DTripleFlat:  BNatural,
	DDoubleFlat:  CNatural,
	DDoubleSharp: ENatural,
	DTripleSharp: FNatural,

	ETripleFlat:  CSharp,
	EDoubleFlat:  DNatural,
	ESharp:       FNatural,
	EDoubleSharp: FSharp,
	ETripleSharp: GNatural,

	FTripleFlat:  DNatural,
	FDoubleFlat:  EFlat,
	FFlat:        ENatural,
	FDoubleSharp: GNatural,
	FTripleSharp: GSharp,

	GTripleFlat:  ENatural,
	GDoubleFlat:  FNatural,
	GDoubleSharp: ANatural,
	GTripleSharp: ASharp,

	ATripleFlat:  FSharp,
	ADoubleFlat:  GNatural,
	ADoubleSharp: BNatural,
	ATripleSharp: CNatural,

	BTripleFlat:  GSharp,
	BDoubleFlat:  ANatural,
	BSharp:       CNatural,
	BDoubleSharp: CSharp,
	BTripleSharp: DNatural,
}

// Normalize returns normalized note
func (n Note) Normalize() Note {
	if v, found := normalizedNotes[n]; found {
		return v
	}

	return n
}
