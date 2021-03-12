package alphabet

var mapAlphabetToNext = map[Alphabet]Alphabet{
	A: B,
	B: C,
	C: D,
	D: E,
	E: F,
	F: G,
	G: A,
}

// Next returns next alphabet
func (a Alphabet) Next() Alphabet {
	return mapAlphabetToNext[a]
}
