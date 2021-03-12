package alphabet

var mapAlphabetPrevious = map[Alphabet]Alphabet{
	A: G, B: A, C: B, D: C, E: D, F: E, G: F,
}

// Previous returns previous alphabet
func (a Alphabet) Previous() Alphabet {
	return mapAlphabetPrevious[a]
}
