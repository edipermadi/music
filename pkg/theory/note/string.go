package note

var mapNoteToString = map[Note]string{
	Invalid:     "Invalid",
	CTripleFlat: "CTripleFlat", CDoubleFlat: "CDoubleFlat", CFlat: "CFlat", CNatural: "CNatural", CSharp: "CSharp", CDoubleSharp: "CDoubleSharp", CTripleSharp: "CTripleSharp",
	DTripleFlat: "DTripleFlat", DDoubleFlat: "DDoubleFlat", DFlat: "DFlat", DNatural: "DNatural", DSharp: "DSharp", DDoubleSharp: "DDoubleSharp", DTripleSharp: "DTripleSharp",
	ETripleFlat: "ETripleFlat", EDoubleFlat: "EDoubleFlat", EFlat: "EFlat", ENatural: "ENatural", ESharp: "ESharp", EDoubleSharp: "EDoubleSharp", ETripleSharp: "ETripleSharp",
	FTripleFlat: "FTripleFlat", FDoubleFlat: "FDoubleFlat", FFlat: "FFlat", FNatural: "FNatural", FSharp: "FSharp", FDoubleSharp: "FDoubleSharp", FTripleSharp: "FTripleSharp",
	GTripleFlat: "GTripleFlat", GDoubleFlat: "GDoubleFlat", GFlat: "GFlat", GNatural: "GNatural", GSharp: "GSharp", GDoubleSharp: "GDoubleSharp", GTripleSharp: "GTripleSharp",
	ATripleFlat: "ATripleFlat", ADoubleFlat: "ADoubleFlat", AFlat: "AFlat", ANatural: "ANatural", ASharp: "ASharp", ADoubleSharp: "ADoubleSharp", ATripleSharp: "ATripleSharp",
	BTripleFlat: "BTripleFlat", BDoubleFlat: "BDoubleFlat", BFlat: "BFlat", BNatural: "BNatural", BSharp: "BSharp", BDoubleSharp: "BDoubleSharp", BTripleSharp: "BTripleSharp",
}

// String retuns stringified note
func (n Note) String() string {
	return mapNoteToString[n]
}

// GoString satisfies GoStringer
func (n Note) GoString() string {
	return n.String()
}
