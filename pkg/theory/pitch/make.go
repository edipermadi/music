package pitch

import (
	"fmt"

	"github.com/edipermadi/music/pkg/theory/note"
)

var pitchBuilder = map[note.Note]map[int]Pitch{
	note.CDoubleFlat:  {0: C0DoubleFlat, 1: C1DoubleFlat, 2: C2DoubleFlat, 3: C3DoubleFlat, 4: C4DoubleFlat, 5: C5DoubleFlat, 6: C6DoubleFlat, 7: C7DoubleFlat, 8: C8DoubleFlat, 9: C9DoubleFlat, 10: C10DoubleFlat},
	note.CFlat:        {0: C0Flat, 1: C1Flat, 2: C2Flat, 3: C3Flat, 4: C4Flat, 5: C5Flat, 6: C6Flat, 7: C7Flat, 8: C8Flat, 9: C9Flat, 10: C10Flat},
	note.CNatural:     {-1: CN1Natural, 0: C0Natural, 1: C1Natural, 2: C2Natural, 3: C3Natural, 4: C4Natural, 5: C5Natural, 6: C6Natural, 7: C7Natural, 8: C8Natural, 9: C9Natural, 10: C10Natural},
	note.CSharp:       {-1: CN1Sharp, 0: C0Sharp, 1: C1Sharp, 2: C2Sharp, 3: C3Sharp, 4: C4Sharp, 5: C5Sharp, 6: C6Sharp, 7: C7Sharp, 8: C8Sharp, 9: C9Sharp, 10: C10Sharp},
	note.CDoubleSharp: {-1: CN1DoubleSharp, 0: C0DoubleSharp, 1: C1DoubleSharp, 2: C2DoubleSharp, 3: C3DoubleSharp, 4: C4DoubleSharp, 5: C5DoubleSharp, 6: C6DoubleSharp, 7: C7DoubleSharp, 8: C8DoubleSharp, 9: C9DoubleSharp, 10: C10DoubleSharp},
	note.DDoubleFlat:  {-1: DN1DoubleFlat, 0: D0DoubleFlat, 1: D1DoubleFlat, 2: D2DoubleFlat, 3: D3DoubleFlat, 4: D4DoubleFlat, 5: D5DoubleFlat, 6: D6DoubleFlat, 7: D7DoubleFlat, 8: D8DoubleFlat, 9: D9DoubleFlat, 10: D10DoubleFlat},
	note.DFlat:        {-1: DN1Flat, 0: D0Flat, 1: D1Flat, 2: D2Flat, 3: D3Flat, 4: D4Flat, 5: D5Flat, 6: D6Flat, 7: D7Flat, 8: D8Flat, 9: D9Flat, 10: D10Flat},
	note.DNatural:     {-1: DN1Natural, 0: D0Natural, 1: D1Natural, 2: D2Natural, 3: D3Natural, 4: D4Natural, 5: D5Natural, 6: D6Natural, 7: D7Natural, 8: D8Natural, 9: D9Natural, 10: D10Natural},
	note.DSharp:       {-1: DN1Sharp, 0: D0Sharp, 1: D1Sharp, 2: D2Sharp, 3: D3Sharp, 4: D4Sharp, 5: D5Sharp, 6: D6Sharp, 7: D7Sharp, 8: D8Sharp, 9: D9Sharp, 10: D10Sharp},
	note.DDoubleSharp: {-1: DN1DoubleSharp, 0: D0DoubleSharp, 1: D1DoubleSharp, 2: D2DoubleSharp, 3: D3DoubleSharp, 4: D4DoubleSharp, 5: D5DoubleSharp, 6: D6DoubleSharp, 7: D7DoubleSharp, 8: D8DoubleSharp, 9: D9DoubleSharp, 10: D10DoubleSharp},
	note.EDoubleFlat:  {-1: EN1DoubleFlat, 0: E0DoubleFlat, 1: E1DoubleFlat, 2: E2DoubleFlat, 3: E3DoubleFlat, 4: E4DoubleFlat, 5: E5DoubleFlat, 6: E6DoubleFlat, 7: E7DoubleFlat, 8: E8DoubleFlat, 9: E9DoubleFlat, 10: E10DoubleFlat},
	note.EFlat:        {-1: EN1Flat, 0: E0Flat, 1: E1Flat, 2: E2Flat, 3: E3Flat, 4: E4Flat, 5: E5Flat, 6: E6Flat, 7: E7Flat, 8: E8Flat, 9: E9Flat, 10: E10Flat},
	note.ENatural:     {-1: EN1Natural, 0: E0Natural, 1: E1Natural, 2: E2Natural, 3: E3Natural, 4: E4Natural, 5: E5Natural, 6: E6Natural, 7: E7Natural, 8: E8Natural, 9: E9Natural, 10: E10Natural},
	note.ESharp:       {-1: EN1Sharp, 0: E0Sharp, 1: E1Sharp, 2: E2Sharp, 3: E3Sharp, 4: E4Sharp, 5: E5Sharp, 6: E6Sharp, 7: E7Sharp, 8: E8Sharp, 9: E9Sharp, 10: E10Sharp},
	note.EDoubleSharp: {-1: EN1DoubleSharp, 0: E0DoubleSharp, 1: E1DoubleSharp, 2: E2DoubleSharp, 3: E3DoubleSharp, 4: E4DoubleSharp, 5: E5DoubleSharp, 6: E6DoubleSharp, 7: E7DoubleSharp, 8: E8DoubleSharp, 9: E9DoubleSharp, 10: E10DoubleSharp},
	note.FDoubleFlat:  {-1: FN1DoubleFlat, 0: F0DoubleFlat, 1: F1DoubleFlat, 2: F2DoubleFlat, 3: F3DoubleFlat, 4: F4DoubleFlat, 5: F5DoubleFlat, 6: F6DoubleFlat, 7: F7DoubleFlat, 8: F8DoubleFlat, 9: F9DoubleFlat, 10: F10DoubleFlat},
	note.FFlat:        {-1: FN1Flat, 0: F0Flat, 1: F1Flat, 2: F2Flat, 3: F3Flat, 4: F4Flat, 5: F5Flat, 6: F6Flat, 7: F7Flat, 8: F8Flat, 9: F9Flat, 10: F10Flat},
	note.FNatural:     {-1: FN1Natural, 0: F0Natural, 1: F1Natural, 2: F2Natural, 3: F3Natural, 4: F4Natural, 5: F5Natural, 6: F6Natural, 7: F7Natural, 8: F8Natural, 9: F9Natural, 10: F10Natural},
	note.FSharp:       {-1: FN1Sharp, 0: F0Sharp, 1: F1Sharp, 2: F2Sharp, 3: F3Sharp, 4: F4Sharp, 5: F5Sharp, 6: F6Sharp, 7: F7Sharp, 8: F8Sharp, 9: F9Sharp, 10: F10Sharp},
	note.FDoubleSharp: {-1: FN1DoubleSharp, 0: F0DoubleSharp, 1: F1DoubleSharp, 2: F2DoubleSharp, 3: F3DoubleSharp, 4: F4DoubleSharp, 5: F5DoubleSharp, 6: F6DoubleSharp, 7: F7DoubleSharp, 8: F8DoubleSharp, 9: F9DoubleSharp, 10: F10DoubleSharp},
	note.GDoubleFlat:  {-1: GN1DoubleFlat, 0: G0DoubleFlat, 1: G1DoubleFlat, 2: G2DoubleFlat, 3: G3DoubleFlat, 4: G4DoubleFlat, 5: G5DoubleFlat, 6: G6DoubleFlat, 7: G7DoubleFlat, 8: G8DoubleFlat, 9: G9DoubleFlat, 10: G10DoubleFlat},
	note.GFlat:        {-1: GN1Flat, 0: G0Flat, 1: G1Flat, 2: G2Flat, 3: G3Flat, 4: G4Flat, 5: G5Flat, 6: G6Flat, 7: G7Flat, 8: G8Flat, 9: G9Flat, 10: G10Flat},
	note.GNatural:     {-1: GN1Natural, 0: G0Natural, 1: G1Natural, 2: G2Natural, 3: G3Natural, 4: G4Natural, 5: G5Natural, 6: G6Natural, 7: G7Natural, 8: G8Natural, 9: G9Natural, 10: G10Natural},
	note.GSharp:       {-1: GN1Sharp, 0: G0Sharp, 1: G1Sharp, 2: G2Sharp, 3: G3Sharp, 4: G4Sharp, 5: G5Sharp, 6: G6Sharp, 7: G7Sharp, 8: G8Sharp, 9: G9Sharp, 10: G10Sharp},
	note.GDoubleSharp: {-1: GN1DoubleSharp, 0: G0DoubleSharp, 1: G1DoubleSharp, 2: G2DoubleSharp, 3: G3DoubleSharp, 4: G4DoubleSharp, 5: G5DoubleSharp, 6: G6DoubleSharp, 7: G7DoubleSharp, 8: G8DoubleSharp, 9: G9DoubleSharp, 10: G10DoubleSharp},
	note.ADoubleFlat:  {-1: AN1DoubleFlat, 0: A0DoubleFlat, 1: A1DoubleFlat, 2: A2DoubleFlat, 3: A3DoubleFlat, 4: A4DoubleFlat, 5: A5DoubleFlat, 6: A6DoubleFlat, 7: A7DoubleFlat, 8: A8DoubleFlat, 9: A9DoubleFlat, 10: A10DoubleFlat},
	note.AFlat:        {-1: AN1Flat, 0: A0Flat, 1: A1Flat, 2: A2Flat, 3: A3Flat, 4: A4Flat, 5: A5Flat, 6: A6Flat, 7: A7Flat, 8: A8Flat, 9: A9Flat, 10: A10Flat},
	note.ANatural:     {-1: AN1Natural, 0: A0Natural, 1: A1Natural, 2: A2Natural, 3: A3Natural, 4: A4Natural, 5: A5Natural, 6: A6Natural, 7: A7Natural, 8: A8Natural, 9: A9Natural, 10: A10Natural},
	note.ASharp:       {-1: AN1Sharp, 0: A0Sharp, 1: A1Sharp, 2: A2Sharp, 3: A3Sharp, 4: A4Sharp, 5: A5Sharp, 6: A6Sharp, 7: A7Sharp, 8: A8Sharp, 9: A9Sharp, 10: A10Sharp},
	note.ADoubleSharp: {-1: AN1DoubleSharp, 0: A0DoubleSharp, 1: A1DoubleSharp, 2: A2DoubleSharp, 3: A3DoubleSharp, 4: A4DoubleSharp, 5: A5DoubleSharp, 6: A6DoubleSharp, 7: A7DoubleSharp, 8: A8DoubleSharp, 9: A9DoubleSharp, 10: A10DoubleSharp},
	note.BDoubleFlat:  {-1: BN1DoubleFlat, 0: B0DoubleFlat, 1: B1DoubleFlat, 2: B2DoubleFlat, 3: B3DoubleFlat, 4: B4DoubleFlat, 5: B5DoubleFlat, 6: B6DoubleFlat, 7: B7DoubleFlat, 8: B8DoubleFlat, 9: B9DoubleFlat, 10: B10DoubleFlat},
	note.BFlat:        {-1: BN1Flat, 0: B0Flat, 1: B1Flat, 2: B2Flat, 3: B3Flat, 4: B4Flat, 5: B5Flat, 6: B6Flat, 7: B7Flat, 8: B8Flat, 9: B9Flat, 10: B10Flat},
	note.BNatural:     {-1: BN1Natural, 0: B0Natural, 1: B1Natural, 2: B2Natural, 3: B3Natural, 4: B4Natural, 5: B5Natural, 6: B6Natural, 7: B7Natural, 8: B8Natural, 9: B9Natural, 10: B10Natural},
	note.BSharp:       {-1: BN1Sharp, 0: B0Sharp, 1: B1Sharp, 2: B2Sharp, 3: B3Sharp, 4: B4Sharp, 5: B5Sharp, 6: B6Sharp, 7: B7Sharp, 8: B8Sharp, 9: B9Sharp},
	note.BDoubleSharp: {-1: BN1DoubleSharp, 0: B0DoubleSharp, 1: B1DoubleSharp, 2: B2DoubleSharp, 3: B3DoubleSharp, 4: B4DoubleSharp, 5: B5DoubleSharp, 6: B6DoubleSharp, 7: B7DoubleSharp, 8: B8DoubleSharp, 9: B9DoubleSharp},
}

// Make returns pitch based on given note and octave
func Make(givenNote note.Note, givenOctave int) Pitch {
	switch {
	case givenOctave < -1,
		givenOctave > 10:
		panic("invalid octave number")
	case givenOctave == -1 && givenNote == note.CFlat,
		givenOctave == -1 && givenNote == note.CDoubleFlat,
		givenOctave == -1 && givenNote == note.CTripleFlat,
		givenOctave == -1 && givenNote == note.DTripleFlat,
		givenOctave == 10 && givenNote == note.BSharp,
		givenOctave == 10 && givenNote == note.BDoubleSharp,
		givenOctave == 10 && givenNote == note.BTripleSharp:
		panic("invalid note")
	}

	// normalize triple accidental
	if givenNote.Accidental().Triple() {
		givenNote = givenNote.Normalize()
	}

	if v, found := pitchBuilder[givenNote][givenOctave]; found {
		return v
	}

	panic(fmt.Errorf("failed to build pitch for %s octave %d", givenNote, givenOctave))
}
