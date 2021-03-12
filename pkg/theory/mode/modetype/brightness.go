package modetype

var typeBrighters = map[Type]Type{
	Ionian:     Lydian,
	Dorian:     Mixolydian,
	Phrygian:   Aeolian,
	Lydian:     Lydian,
	Mixolydian: Ionian,
	Aeolian:    Dorian,
	Locrian:    Phrygian,
}

var typeDarkers = map[Type]Type{
	Ionian:     Mixolydian,
	Dorian:     Aeolian,
	Phrygian:   Locrian,
	Lydian:     Ionian,
	Mixolydian: Dorian,
	Aeolian:    Phrygian,
	Locrian:    Locrian,
}

// Brighten return next brighter mode
func (t Type) Brighten() Type {
	return typeBrighters[t]
}

// Darken return next darker mode
func (t Type) Darken() Type {
	return typeDarkers[t]
}
