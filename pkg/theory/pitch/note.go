package pitch

import (
	"fmt"

	"github.com/edipermadi/music/pkg/theory/note"
)

var pitchNotes = map[Pitch]note.Note{
	CN1Natural: note.CNatural, CN1Sharp: note.CSharp, CN1DoubleSharp: note.CDoubleFlat,
	DN1DoubleFlat: note.DDoubleFlat, DN1Flat: note.DFlat, DN1Natural: note.DNatural, DN1Sharp: note.DSharp, DN1DoubleSharp: note.DDoubleSharp,
	EN1DoubleFlat: note.EDoubleFlat, EN1Flat: note.EFlat, EN1Natural: note.ENatural, EN1Sharp: note.ESharp, EN1DoubleSharp: note.EDoubleSharp,
	FN1DoubleFlat: note.FDoubleFlat, FN1Flat: note.FFlat, FN1Natural: note.FNatural, FN1Sharp: note.FSharp, FN1DoubleSharp: note.FDoubleSharp,
	GN1DoubleFlat: note.GDoubleFlat, GN1Flat: note.GFlat, GN1Natural: note.GNatural, GN1Sharp: note.GSharp, GN1DoubleSharp: note.GDoubleSharp,
	AN1DoubleFlat: note.ADoubleFlat, AN1Flat: note.AFlat, AN1Natural: note.ANatural, AN1Sharp: note.ASharp, AN1DoubleSharp: note.ADoubleSharp,
	BN1DoubleFlat: note.BDoubleFlat, BN1Flat: note.BFlat, BN1Natural: note.BNatural, BN1Sharp: note.BSharp, BN1DoubleSharp: note.BDoubleSharp,

	C0DoubleFlat: note.CDoubleFlat, C0Flat: note.CFlat, C0Natural: note.CNatural, C0Sharp: note.CSharp, C0DoubleSharp: note.CDoubleSharp,
	D0DoubleFlat: note.DDoubleFlat, D0Flat: note.DFlat, D0Natural: note.DNatural, D0Sharp: note.DSharp, D0DoubleSharp: note.DDoubleSharp,
	E0DoubleFlat: note.EDoubleFlat, E0Flat: note.EFlat, E0Natural: note.ENatural, E0Sharp: note.ESharp, E0DoubleSharp: note.EDoubleSharp,
	F0DoubleFlat: note.FDoubleFlat, F0Flat: note.FFlat, F0Natural: note.FNatural, F0Sharp: note.FSharp, F0DoubleSharp: note.FDoubleSharp,
	G0DoubleFlat: note.GDoubleFlat, G0Flat: note.GFlat, G0Natural: note.GNatural, G0Sharp: note.GSharp, G0DoubleSharp: note.GDoubleSharp,
	A0DoubleFlat: note.ADoubleFlat, A0Flat: note.AFlat, A0Natural: note.ANatural, A0Sharp: note.ASharp, A0DoubleSharp: note.ADoubleSharp,
	B0DoubleFlat: note.BDoubleFlat, B0Flat: note.BFlat, B0Natural: note.BNatural, B0Sharp: note.BSharp, B0DoubleSharp: note.BDoubleSharp,

	C1DoubleFlat: note.CDoubleFlat, C1Flat: note.CFlat, C1Natural: note.CNatural, C1Sharp: note.CSharp, C1DoubleSharp: note.CDoubleSharp,
	D1DoubleFlat: note.DDoubleFlat, D1Flat: note.DFlat, D1Natural: note.DNatural, D1Sharp: note.DSharp, D1DoubleSharp: note.DDoubleSharp,
	E1DoubleFlat: note.EDoubleFlat, E1Flat: note.EFlat, E1Natural: note.ENatural, E1Sharp: note.ESharp, E1DoubleSharp: note.EDoubleSharp,
	F1DoubleFlat: note.FDoubleFlat, F1Flat: note.FFlat, F1Natural: note.FNatural, F1Sharp: note.FSharp, F1DoubleSharp: note.FDoubleSharp,
	G1DoubleFlat: note.GDoubleFlat, G1Flat: note.GFlat, G1Natural: note.GNatural, G1Sharp: note.GSharp, G1DoubleSharp: note.GDoubleSharp,
	A1DoubleFlat: note.ADoubleFlat, A1Flat: note.AFlat, A1Natural: note.ANatural, A1Sharp: note.ASharp, A1DoubleSharp: note.ADoubleSharp,
	B1DoubleFlat: note.BDoubleFlat, B1Flat: note.BFlat, B1Natural: note.BNatural, B1Sharp: note.BSharp, B1DoubleSharp: note.BDoubleSharp,

	C2DoubleFlat: note.CDoubleFlat, C2Flat: note.CFlat, C2Natural: note.CNatural, C2Sharp: note.CSharp, C2DoubleSharp: note.CDoubleSharp,
	D2DoubleFlat: note.DDoubleFlat, D2Flat: note.DFlat, D2Natural: note.DNatural, D2Sharp: note.DSharp, D2DoubleSharp: note.DDoubleSharp,
	E2DoubleFlat: note.EDoubleFlat, E2Flat: note.EFlat, E2Natural: note.ENatural, E2Sharp: note.ESharp, E2DoubleSharp: note.EDoubleSharp,
	F2DoubleFlat: note.FDoubleFlat, F2Flat: note.FFlat, F2Natural: note.FNatural, F2Sharp: note.FSharp, F2DoubleSharp: note.FDoubleSharp,
	G2DoubleFlat: note.GDoubleFlat, G2Flat: note.GFlat, G2Natural: note.GNatural, G2Sharp: note.GSharp, G2DoubleSharp: note.GDoubleSharp,
	A2DoubleFlat: note.ADoubleFlat, A2Flat: note.AFlat, A2Natural: note.ANatural, A2Sharp: note.ASharp, A2DoubleSharp: note.ADoubleSharp,
	B2DoubleFlat: note.BDoubleFlat, B2Flat: note.BFlat, B2Natural: note.BNatural, B2Sharp: note.BSharp, B2DoubleSharp: note.BDoubleSharp,

	C3DoubleFlat: note.CDoubleFlat, C3Flat: note.CFlat, C3Natural: note.CNatural, C3Sharp: note.CSharp, C3DoubleSharp: note.CDoubleSharp,
	D3DoubleFlat: note.DDoubleFlat, D3Flat: note.DFlat, D3Natural: note.DNatural, D3Sharp: note.DSharp, D3DoubleSharp: note.DDoubleSharp,
	E3DoubleFlat: note.EDoubleFlat, E3Flat: note.EFlat, E3Natural: note.ENatural, E3Sharp: note.ESharp, E3DoubleSharp: note.EDoubleSharp,
	F3DoubleFlat: note.FDoubleFlat, F3Flat: note.FFlat, F3Natural: note.FNatural, F3Sharp: note.FSharp, F3DoubleSharp: note.FDoubleSharp,
	G3DoubleFlat: note.GDoubleFlat, G3Flat: note.GFlat, G3Natural: note.GNatural, G3Sharp: note.GSharp, G3DoubleSharp: note.GDoubleSharp,
	A3DoubleFlat: note.ADoubleFlat, A3Flat: note.AFlat, A3Natural: note.ANatural, A3Sharp: note.ASharp, A3DoubleSharp: note.ADoubleSharp,
	B3DoubleFlat: note.BDoubleFlat, B3Flat: note.BFlat, B3Natural: note.BNatural, B3Sharp: note.BSharp, B3DoubleSharp: note.BDoubleSharp,

	C4DoubleFlat: note.CDoubleFlat, C4Flat: note.CFlat, C4Natural: note.CNatural, C4Sharp: note.CSharp, C4DoubleSharp: note.CDoubleSharp,
	D4DoubleFlat: note.DDoubleFlat, D4Flat: note.DFlat, D4Natural: note.DNatural, D4Sharp: note.DSharp, D4DoubleSharp: note.DDoubleSharp,
	E4DoubleFlat: note.EDoubleFlat, E4Flat: note.EFlat, E4Natural: note.ENatural, E4Sharp: note.ESharp, E4DoubleSharp: note.EDoubleSharp,
	F4DoubleFlat: note.FDoubleFlat, F4Flat: note.FFlat, F4Natural: note.FNatural, F4Sharp: note.FSharp, F4DoubleSharp: note.FDoubleSharp,
	G4DoubleFlat: note.GDoubleFlat, G4Flat: note.GFlat, G4Natural: note.GNatural, G4Sharp: note.GSharp, G4DoubleSharp: note.GDoubleSharp,
	A4DoubleFlat: note.ADoubleFlat, A4Flat: note.AFlat, A4Natural: note.ANatural, A4Sharp: note.ASharp, A4DoubleSharp: note.ADoubleSharp,
	B4DoubleFlat: note.BDoubleFlat, B4Flat: note.BFlat, B4Natural: note.BNatural, B4Sharp: note.BSharp, B4DoubleSharp: note.BDoubleSharp,

	C5DoubleFlat: note.CDoubleFlat, C5Flat: note.CFlat, C5Natural: note.CNatural, C5Sharp: note.CSharp, C5DoubleSharp: note.CDoubleFlat,
	D5DoubleFlat: note.DDoubleFlat, D5Flat: note.DFlat, D5Natural: note.DNatural, D5Sharp: note.DSharp, D5DoubleSharp: note.DDoubleFlat,
	E5DoubleFlat: note.EDoubleFlat, E5Flat: note.EFlat, E5Natural: note.ENatural, E5Sharp: note.ESharp, E5DoubleSharp: note.EDoubleFlat,
	F5DoubleFlat: note.FDoubleFlat, F5Flat: note.FFlat, F5Natural: note.FNatural, F5Sharp: note.FSharp, F5DoubleSharp: note.FDoubleFlat,
	G5DoubleFlat: note.GDoubleFlat, G5Flat: note.GFlat, G5Natural: note.GNatural, G5Sharp: note.GSharp, G5DoubleSharp: note.GDoubleFlat,
	A5DoubleFlat: note.ADoubleFlat, A5Flat: note.AFlat, A5Natural: note.ANatural, A5Sharp: note.ASharp, A5DoubleSharp: note.ADoubleFlat,
	B5DoubleFlat: note.BDoubleFlat, B5Flat: note.BFlat, B5Natural: note.BNatural, B5Sharp: note.BSharp, B5DoubleSharp: note.BDoubleFlat,

	C6DoubleFlat: note.CDoubleFlat, C6Flat: note.CFlat, C6Natural: note.CNatural, C6Sharp: note.CSharp, C6DoubleSharp: note.CDoubleFlat,
	D6DoubleFlat: note.DDoubleFlat, D6Flat: note.DFlat, D6Natural: note.DNatural, D6Sharp: note.DSharp, D6DoubleSharp: note.DDoubleFlat,
	E6DoubleFlat: note.EDoubleFlat, E6Flat: note.EFlat, E6Natural: note.ENatural, E6Sharp: note.ESharp, E6DoubleSharp: note.EDoubleFlat,
	F6DoubleFlat: note.FDoubleFlat, F6Flat: note.FFlat, F6Natural: note.FNatural, F6Sharp: note.FSharp, F6DoubleSharp: note.FDoubleFlat,
	G6DoubleFlat: note.GDoubleFlat, G6Flat: note.GFlat, G6Natural: note.GNatural, G6Sharp: note.GSharp, G6DoubleSharp: note.GDoubleFlat,
	A6DoubleFlat: note.ADoubleFlat, A6Flat: note.AFlat, A6Natural: note.ANatural, A6Sharp: note.ASharp, A6DoubleSharp: note.ADoubleFlat,
	B6DoubleFlat: note.BDoubleFlat, B6Flat: note.BFlat, B6Natural: note.BNatural, B6Sharp: note.BSharp, B6DoubleSharp: note.BDoubleFlat,

	C7DoubleFlat: note.CDoubleFlat, C7Flat: note.CFlat, C7Natural: note.CNatural, C7Sharp: note.CSharp, C7DoubleSharp: note.CDoubleFlat,
	D7DoubleFlat: note.DDoubleFlat, D7Flat: note.DFlat, D7Natural: note.DNatural, D7Sharp: note.DSharp, D7DoubleSharp: note.DDoubleFlat,
	E7DoubleFlat: note.EDoubleFlat, E7Flat: note.EFlat, E7Natural: note.ENatural, E7Sharp: note.ESharp, E7DoubleSharp: note.EDoubleFlat,
	F7DoubleFlat: note.FDoubleFlat, F7Flat: note.FFlat, F7Natural: note.FNatural, F7Sharp: note.FSharp, F7DoubleSharp: note.FDoubleFlat,
	G7DoubleFlat: note.GDoubleFlat, G7Flat: note.GFlat, G7Natural: note.GNatural, G7Sharp: note.GSharp, G7DoubleSharp: note.GDoubleFlat,
	A7DoubleFlat: note.ADoubleFlat, A7Flat: note.AFlat, A7Natural: note.ANatural, A7Sharp: note.ASharp, A7DoubleSharp: note.ADoubleFlat,
	B7DoubleFlat: note.BDoubleFlat, B7Flat: note.BFlat, B7Natural: note.BNatural, B7Sharp: note.BSharp, B7DoubleSharp: note.BDoubleFlat,

	C8DoubleFlat: note.CDoubleFlat, C8Flat: note.CFlat, C8Natural: note.CNatural, C8Sharp: note.CSharp, C8DoubleSharp: note.CDoubleFlat,
	D8DoubleFlat: note.DDoubleFlat, D8Flat: note.DFlat, D8Natural: note.DNatural, D8Sharp: note.DSharp, D8DoubleSharp: note.DDoubleFlat,
	E8DoubleFlat: note.EDoubleFlat, E8Flat: note.EFlat, E8Natural: note.ENatural, E8Sharp: note.ESharp, E8DoubleSharp: note.EDoubleFlat,
	F8DoubleFlat: note.FDoubleFlat, F8Flat: note.FFlat, F8Natural: note.FNatural, F8Sharp: note.FSharp, F8DoubleSharp: note.FDoubleFlat,
	G8DoubleFlat: note.GDoubleFlat, G8Flat: note.GFlat, G8Natural: note.GNatural, G8Sharp: note.GSharp, G8DoubleSharp: note.GDoubleFlat,
	A8DoubleFlat: note.ADoubleFlat, A8Flat: note.AFlat, A8Natural: note.ANatural, A8Sharp: note.ASharp, A8DoubleSharp: note.ADoubleFlat,
	B8DoubleFlat: note.BDoubleFlat, B8Flat: note.BFlat, B8Natural: note.BNatural, B8Sharp: note.BSharp, B8DoubleSharp: note.BDoubleFlat,

	C9DoubleFlat: note.CDoubleFlat, C9Flat: note.CFlat, C9Natural: note.CNatural, C9Sharp: note.CSharp, C9DoubleSharp: note.CDoubleFlat,
	D9DoubleFlat: note.DDoubleFlat, D9Flat: note.DFlat, D9Natural: note.DNatural, D9Sharp: note.DSharp, D9DoubleSharp: note.DDoubleFlat,
	E9DoubleFlat: note.EDoubleFlat, E9Flat: note.EFlat, E9Natural: note.ENatural, E9Sharp: note.ESharp, E9DoubleSharp: note.EDoubleFlat,
	F9DoubleFlat: note.FDoubleFlat, F9Flat: note.FFlat, F9Natural: note.FNatural, F9Sharp: note.FSharp, F9DoubleSharp: note.FDoubleFlat,
	G9DoubleFlat: note.GDoubleFlat, G9Flat: note.GFlat, G9Natural: note.GNatural, G9Sharp: note.GSharp, G9DoubleSharp: note.GDoubleFlat,
	A9DoubleFlat: note.ADoubleFlat, A9Flat: note.AFlat, A9Natural: note.ANatural, A9Sharp: note.ASharp, A9DoubleSharp: note.ADoubleFlat,
	B9DoubleFlat: note.BDoubleFlat, B9Flat: note.BFlat, B9Natural: note.BNatural, B9Sharp: note.BSharp, B9DoubleSharp: note.BDoubleFlat,

	C10DoubleFlat: note.CDoubleFlat, C10Flat: note.CFlat, C10Natural: note.CNatural, C10Sharp: note.CSharp, C10DoubleSharp: note.CDoubleFlat,
	D10DoubleFlat: note.DDoubleFlat, D10Flat: note.DFlat, D10Natural: note.DNatural, D10Sharp: note.DSharp, D10DoubleSharp: note.DDoubleFlat,
	E10DoubleFlat: note.EDoubleFlat, E10Flat: note.EFlat, E10Natural: note.ENatural, E10Sharp: note.ESharp, E10DoubleSharp: note.EDoubleFlat,
	F10DoubleFlat: note.FDoubleFlat, F10Flat: note.FFlat, F10Natural: note.FNatural, F10Sharp: note.FSharp, F10DoubleSharp: note.FDoubleFlat,
	G10DoubleFlat: note.GDoubleFlat, G10Flat: note.GFlat, G10Natural: note.GNatural, G10Sharp: note.GSharp, G10DoubleSharp: note.GDoubleFlat,
	A10DoubleFlat: note.ADoubleFlat, A10Flat: note.AFlat, A10Natural: note.ANatural, A10Sharp: note.ASharp, A10DoubleSharp: note.ADoubleFlat,
	B10DoubleFlat: note.BDoubleFlat, B10Flat: note.BFlat, B10Natural: note.BNatural,
}

// Note return the note of the pitch
func (p Pitch) Note() note.Note {
	v, found := pitchNotes[p]
	if !found {
		// should never happened
		panic(fmt.Errorf("no such note for pitch %s", p))
	}

	return v
}
