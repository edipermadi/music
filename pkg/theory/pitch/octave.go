package pitch

var pitchOctaves = map[Pitch]int{
	CN1Natural: -1, CN1Sharp: -1, CN1DoubleSharp: -1,
	DN1DoubleFlat: -1, DN1Flat: -1, DN1Natural: -1, DN1Sharp: -1, DN1DoubleSharp: -1,
	EN1DoubleFlat: -1, EN1Flat: -1, EN1Natural: -1, EN1Sharp: -1, EN1DoubleSharp: -1,
	FN1DoubleFlat: -1, FN1Flat: -1, FN1Natural: -1, FN1Sharp: -1, FN1DoubleSharp: -1,
	GN1DoubleFlat: -1, GN1Flat: -1, GN1Natural: -1, GN1Sharp: -1, GN1DoubleSharp: -1,
	AN1DoubleFlat: -1, AN1Flat: -1, AN1Natural: -1, AN1Sharp: -1, AN1DoubleSharp: -1,
	BN1DoubleFlat: -1, BN1Flat: -1, BN1Natural: -1, BN1Sharp: 0, BN1DoubleSharp: 0,

	C0DoubleFlat: -1, C0Flat: -1, C0Natural: 0, C0Sharp: 0, C0DoubleSharp: 0,
	D0DoubleFlat: 0, D0Flat: 0, D0Natural: 0, D0Sharp: 0, D0DoubleSharp: 0,
	E0DoubleFlat: 0, E0Flat: 0, E0Natural: 0, E0Sharp: 0, E0DoubleSharp: 0,
	F0DoubleFlat: 0, F0Flat: 0, F0Natural: 0, F0Sharp: 0, F0DoubleSharp: 0,
	G0DoubleFlat: 0, G0Flat: 0, G0Natural: 0, G0Sharp: 0, G0DoubleSharp: 0,
	A0DoubleFlat: 0, A0Flat: 0, A0Natural: 0, A0Sharp: 0, A0DoubleSharp: 0,
	B0DoubleFlat: 0, B0Flat: 0, B0Natural: 0, B0Sharp: 1, B0DoubleSharp: 1,

	C1DoubleFlat: 0, C1Flat: 0, C1Natural: 1, C1Sharp: 1, C1DoubleSharp: 1,
	D1DoubleFlat: 1, D1Flat: 1, D1Natural: 1, D1Sharp: 1, D1DoubleSharp: 1,
	E1DoubleFlat: 1, E1Flat: 1, E1Natural: 1, E1Sharp: 1, E1DoubleSharp: 1,
	F1DoubleFlat: 1, F1Flat: 1, F1Natural: 1, F1Sharp: 1, F1DoubleSharp: 1,
	G1DoubleFlat: 1, G1Flat: 1, G1Natural: 1, G1Sharp: 1, G1DoubleSharp: 1,
	A1DoubleFlat: 1, A1Flat: 1, A1Natural: 1, A1Sharp: 1, A1DoubleSharp: 1,
	B1DoubleFlat: 1, B1Flat: 1, B1Natural: 1, B1Sharp: 2, B1DoubleSharp: 2,

	C2DoubleFlat: 1, C2Flat: 1, C2Natural: 2, C2Sharp: 2, C2DoubleSharp: 2,
	D2DoubleFlat: 2, D2Flat: 2, D2Natural: 2, D2Sharp: 2, D2DoubleSharp: 2,
	E2DoubleFlat: 2, E2Flat: 2, E2Natural: 2, E2Sharp: 2, E2DoubleSharp: 2,
	F2DoubleFlat: 2, F2Flat: 2, F2Natural: 2, F2Sharp: 2, F2DoubleSharp: 2,
	G2DoubleFlat: 2, G2Flat: 2, G2Natural: 2, G2Sharp: 2, G2DoubleSharp: 2,
	A2DoubleFlat: 2, A2Flat: 2, A2Natural: 2, A2Sharp: 2, A2DoubleSharp: 2,
	B2DoubleFlat: 2, B2Flat: 2, B2Natural: 2, B2Sharp: 3, B2DoubleSharp: 3,

	C3DoubleFlat: 2, C3Flat: 2, C3Natural: 3, C3Sharp: 3, C3DoubleSharp: 3,
	D3DoubleFlat: 3, D3Flat: 3, D3Natural: 3, D3Sharp: 3, D3DoubleSharp: 3,
	E3DoubleFlat: 3, E3Flat: 3, E3Natural: 3, E3Sharp: 3, E3DoubleSharp: 3,
	F3DoubleFlat: 3, F3Flat: 3, F3Natural: 3, F3Sharp: 3, F3DoubleSharp: 3,
	G3DoubleFlat: 3, G3Flat: 3, G3Natural: 3, G3Sharp: 3, G3DoubleSharp: 3,
	A3DoubleFlat: 3, A3Flat: 3, A3Natural: 3, A3Sharp: 3, A3DoubleSharp: 3,
	B3DoubleFlat: 3, B3Flat: 3, B3Natural: 3, B3Sharp: 4, B3DoubleSharp: 4,

	C4DoubleFlat: 3, C4Flat: 3, C4Natural: 4, C4Sharp: 4, C4DoubleSharp: 4,
	D4DoubleFlat: 4, D4Flat: 4, D4Natural: 4, D4Sharp: 4, D4DoubleSharp: 4,
	E4DoubleFlat: 4, E4Flat: 4, E4Natural: 4, E4Sharp: 4, E4DoubleSharp: 4,
	F4DoubleFlat: 4, F4Flat: 4, F4Natural: 4, F4Sharp: 4, F4DoubleSharp: 4,
	G4DoubleFlat: 4, G4Flat: 4, G4Natural: 4, G4Sharp: 4, G4DoubleSharp: 4,
	A4DoubleFlat: 4, A4Flat: 4, A4Natural: 4, A4Sharp: 4, A4DoubleSharp: 4,
	B4DoubleFlat: 4, B4Flat: 4, B4Natural: 4, B4Sharp: 5, B4DoubleSharp: 5,

	C5DoubleFlat: 4, C5Flat: 4, C5Natural: 5, C5Sharp: 5, C5DoubleSharp: 5,
	D5DoubleFlat: 5, D5Flat: 5, D5Natural: 5, D5Sharp: 5, D5DoubleSharp: 5,
	E5DoubleFlat: 5, E5Flat: 5, E5Natural: 5, E5Sharp: 5, E5DoubleSharp: 5,
	F5DoubleFlat: 5, F5Flat: 5, F5Natural: 5, F5Sharp: 5, F5DoubleSharp: 5,
	G5DoubleFlat: 5, G5Flat: 5, G5Natural: 5, G5Sharp: 5, G5DoubleSharp: 5,
	A5DoubleFlat: 5, A5Flat: 5, A5Natural: 5, A5Sharp: 5, A5DoubleSharp: 5,
	B5DoubleFlat: 5, B5Flat: 5, B5Natural: 5, B5Sharp: 6, B5DoubleSharp: 6,

	C6DoubleFlat: 5, C6Flat: 5, C6Natural: 6, C6Sharp: 6, C6DoubleSharp: 6,
	D6DoubleFlat: 6, D6Flat: 6, D6Natural: 6, D6Sharp: 6, D6DoubleSharp: 6,
	E6DoubleFlat: 6, E6Flat: 6, E6Natural: 6, E6Sharp: 6, E6DoubleSharp: 6,
	F6DoubleFlat: 6, F6Flat: 6, F6Natural: 6, F6Sharp: 6, F6DoubleSharp: 6,
	G6DoubleFlat: 6, G6Flat: 6, G6Natural: 6, G6Sharp: 6, G6DoubleSharp: 6,
	A6DoubleFlat: 6, A6Flat: 6, A6Natural: 6, A6Sharp: 6, A6DoubleSharp: 6,
	B6DoubleFlat: 6, B6Flat: 6, B6Natural: 6, B6Sharp: 7, B6DoubleSharp: 7,

	C7DoubleFlat: 6, C7Flat: 6, C7Natural: 7, C7Sharp: 7, C7DoubleSharp: 7,
	D7DoubleFlat: 7, D7Flat: 7, D7Natural: 7, D7Sharp: 7, D7DoubleSharp: 7,
	E7DoubleFlat: 7, E7Flat: 7, E7Natural: 7, E7Sharp: 7, E7DoubleSharp: 7,
	F7DoubleFlat: 7, F7Flat: 7, F7Natural: 7, F7Sharp: 7, F7DoubleSharp: 7,
	G7DoubleFlat: 7, G7Flat: 7, G7Natural: 7, G7Sharp: 7, G7DoubleSharp: 7,
	A7DoubleFlat: 7, A7Flat: 7, A7Natural: 7, A7Sharp: 7, A7DoubleSharp: 7,
	B7DoubleFlat: 7, B7Flat: 7, B7Natural: 7, B7Sharp: 8, B7DoubleSharp: 8,

	C8DoubleFlat: 7, C8Flat: 7, C8Natural: 8, C8Sharp: 8, C8DoubleSharp: 8,
	D8DoubleFlat: 8, D8Flat: 8, D8Natural: 8, D8Sharp: 8, D8DoubleSharp: 8,
	E8DoubleFlat: 8, E8Flat: 8, E8Natural: 8, E8Sharp: 8, E8DoubleSharp: 8,
	F8DoubleFlat: 8, F8Flat: 8, F8Natural: 8, F8Sharp: 8, F8DoubleSharp: 8,
	G8DoubleFlat: 8, G8Flat: 8, G8Natural: 8, G8Sharp: 8, G8DoubleSharp: 8,
	A8DoubleFlat: 8, A8Flat: 8, A8Natural: 8, A8Sharp: 8, A8DoubleSharp: 8,
	B8DoubleFlat: 8, B8Flat: 8, B8Natural: 8, B8Sharp: 9, B8DoubleSharp: 9,

	C9DoubleFlat: 8, C9Flat: 8, C9Natural: 9, C9Sharp: 9, C9DoubleSharp: 9,
	D9DoubleFlat: 9, D9Flat: 9, D9Natural: 9, D9Sharp: 9, D9DoubleSharp: 9,
	E9DoubleFlat: 9, E9Flat: 9, E9Natural: 9, E9Sharp: 9, E9DoubleSharp: 9,
	F9DoubleFlat: 9, F9Flat: 9, F9Natural: 9, F9Sharp: 9, F9DoubleSharp: 9,
	G9DoubleFlat: 9, G9Flat: 9, G9Natural: 9, G9Sharp: 9, G9DoubleSharp: 9,
	A9DoubleFlat: 9, A9Flat: 9, A9Natural: 9, A9Sharp: 9, A9DoubleSharp: 9,
	B9DoubleFlat: 9, B9Flat: 9, B9Natural: 9, B9Sharp: 10, B9DoubleSharp: 10,

	C10DoubleFlat: 9, C10Flat: 9, C10Natural: 10, C10Sharp: 10, C10DoubleSharp: 10,
	D10DoubleFlat: 10, D10Flat: 10, D10Natural: 10, D10Sharp: 10, D10DoubleSharp: 10,
	E10DoubleFlat: 10, E10Flat: 10, E10Natural: 10, E10Sharp: 10, E10DoubleSharp: 10,
	F10DoubleFlat: 10, F10Flat: 10, F10Natural: 10, F10Sharp: 10, F10DoubleSharp: 10,
	G10DoubleFlat: 10, G10Flat: 10, G10Natural: 10, G10Sharp: 10, G10DoubleSharp: 10,
	A10DoubleFlat: 10, A10Flat: 10, A10Natural: 10, A10Sharp: 10, A10DoubleSharp: 10,
	B10DoubleFlat: 10, B10Flat: 10, B10Natural: 10,
}

// Octave returns octave number from a pitch
func (p Pitch) Octave() int {
	return pitchOctaves[p]
}
