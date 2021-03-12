package pitch

import (
	"sort"

	"github.com/edipermadi/music/pkg/theory/limit"
	"github.com/edipermadi/music/pkg/theory/note/accidental"
)

// Pitch represents a musical pitch
type Pitch int

// Enumeration of pitches
const (
	Invalid Pitch = iota

	CN1Natural     Pitch = iota
	CN1Sharp       Pitch = iota
	CN1DoubleSharp Pitch = iota

	DN1DoubleFlat  Pitch = iota
	DN1Flat        Pitch = iota
	DN1Natural     Pitch = iota
	DN1Sharp       Pitch = iota
	DN1DoubleSharp Pitch = iota

	EN1DoubleFlat  Pitch = iota
	EN1Flat        Pitch = iota
	EN1Natural     Pitch = iota
	EN1Sharp       Pitch = iota
	EN1DoubleSharp Pitch = iota

	FN1DoubleFlat  Pitch = iota
	FN1Flat        Pitch = iota
	FN1Natural     Pitch = iota
	FN1Sharp       Pitch = iota
	FN1DoubleSharp Pitch = iota

	GN1DoubleFlat  Pitch = iota
	GN1Flat        Pitch = iota
	GN1Natural     Pitch = iota
	GN1Sharp       Pitch = iota
	GN1DoubleSharp Pitch = iota

	AN1DoubleFlat  Pitch = iota
	AN1Flat        Pitch = iota
	AN1Natural     Pitch = iota
	AN1Sharp       Pitch = iota
	AN1DoubleSharp Pitch = iota

	BN1DoubleFlat  Pitch = iota
	BN1Flat        Pitch = iota
	BN1Natural     Pitch = iota
	BN1Sharp       Pitch = iota
	BN1DoubleSharp Pitch = iota

	C0DoubleFlat  Pitch = iota
	C0Flat        Pitch = iota
	C0Natural     Pitch = iota
	C0Sharp       Pitch = iota
	C0DoubleSharp Pitch = iota

	D0DoubleFlat  Pitch = iota
	D0Flat        Pitch = iota
	D0Natural     Pitch = iota
	D0Sharp       Pitch = iota
	D0DoubleSharp Pitch = iota

	E0DoubleFlat  Pitch = iota
	E0Flat        Pitch = iota
	E0Natural     Pitch = iota
	E0Sharp       Pitch = iota
	E0DoubleSharp Pitch = iota

	F0DoubleFlat  Pitch = iota
	F0Flat        Pitch = iota
	F0Natural     Pitch = iota
	F0Sharp       Pitch = iota
	F0DoubleSharp Pitch = iota

	G0DoubleFlat  Pitch = iota
	G0Flat        Pitch = iota
	G0Natural     Pitch = iota
	G0Sharp       Pitch = iota
	G0DoubleSharp Pitch = iota

	A0DoubleFlat  Pitch = iota
	A0Flat        Pitch = iota
	A0Natural     Pitch = iota
	A0Sharp       Pitch = iota
	A0DoubleSharp Pitch = iota

	B0DoubleFlat  Pitch = iota
	B0Flat        Pitch = iota
	B0Natural     Pitch = iota
	B0Sharp       Pitch = iota
	B0DoubleSharp Pitch = iota

	C1DoubleFlat  Pitch = iota
	C1Flat        Pitch = iota
	C1Natural     Pitch = iota
	C1Sharp       Pitch = iota
	C1DoubleSharp Pitch = iota

	D1DoubleFlat  Pitch = iota
	D1Flat        Pitch = iota
	D1Natural     Pitch = iota
	D1Sharp       Pitch = iota
	D1DoubleSharp Pitch = iota

	E1DoubleFlat  Pitch = iota
	E1Flat        Pitch = iota
	E1Natural     Pitch = iota
	E1Sharp       Pitch = iota
	E1DoubleSharp Pitch = iota

	F1DoubleFlat  Pitch = iota
	F1Flat        Pitch = iota
	F1Natural     Pitch = iota
	F1Sharp       Pitch = iota
	F1DoubleSharp Pitch = iota

	G1DoubleFlat  Pitch = iota
	G1Flat        Pitch = iota
	G1Natural     Pitch = iota
	G1Sharp       Pitch = iota
	G1DoubleSharp Pitch = iota

	A1DoubleFlat  Pitch = iota
	A1Flat        Pitch = iota
	A1Natural     Pitch = iota
	A1Sharp       Pitch = iota
	A1DoubleSharp Pitch = iota

	B1DoubleFlat  Pitch = iota
	B1Flat        Pitch = iota
	B1Natural     Pitch = iota
	B1Sharp       Pitch = iota
	B1DoubleSharp Pitch = iota

	C2DoubleFlat  Pitch = iota
	C2Flat        Pitch = iota
	C2Natural     Pitch = iota
	C2Sharp       Pitch = iota
	C2DoubleSharp Pitch = iota

	D2DoubleFlat  Pitch = iota
	D2Flat        Pitch = iota
	D2Natural     Pitch = iota
	D2Sharp       Pitch = iota
	D2DoubleSharp Pitch = iota

	E2DoubleFlat  Pitch = iota
	E2Flat        Pitch = iota
	E2Natural     Pitch = iota
	E2Sharp       Pitch = iota
	E2DoubleSharp Pitch = iota

	F2DoubleFlat  Pitch = iota
	F2Flat        Pitch = iota
	F2Natural     Pitch = iota
	F2Sharp       Pitch = iota
	F2DoubleSharp Pitch = iota

	G2DoubleFlat  Pitch = iota
	G2Flat        Pitch = iota
	G2Natural     Pitch = iota
	G2Sharp       Pitch = iota
	G2DoubleSharp Pitch = iota

	A2DoubleFlat  Pitch = iota
	A2Flat        Pitch = iota
	A2Natural     Pitch = iota
	A2Sharp       Pitch = iota
	A2DoubleSharp Pitch = iota

	B2DoubleFlat  Pitch = iota
	B2Flat        Pitch = iota
	B2Natural     Pitch = iota
	B2Sharp       Pitch = iota
	B2DoubleSharp Pitch = iota

	C3DoubleFlat  Pitch = iota
	C3Flat        Pitch = iota
	C3Natural     Pitch = iota
	C3Sharp       Pitch = iota
	C3DoubleSharp Pitch = iota

	D3DoubleFlat  Pitch = iota
	D3Flat        Pitch = iota
	D3Natural     Pitch = iota
	D3Sharp       Pitch = iota
	D3DoubleSharp Pitch = iota

	E3DoubleFlat  Pitch = iota
	E3Flat        Pitch = iota
	E3Natural     Pitch = iota
	E3Sharp       Pitch = iota
	E3DoubleSharp Pitch = iota

	F3DoubleFlat  Pitch = iota
	F3Flat        Pitch = iota
	F3Natural     Pitch = iota
	F3Sharp       Pitch = iota
	F3DoubleSharp Pitch = iota

	G3DoubleFlat  Pitch = iota
	G3Flat        Pitch = iota
	G3Natural     Pitch = iota
	G3Sharp       Pitch = iota
	G3DoubleSharp Pitch = iota

	A3DoubleFlat  Pitch = iota
	A3Flat        Pitch = iota
	A3Natural     Pitch = iota
	A3Sharp       Pitch = iota
	A3DoubleSharp Pitch = iota

	B3DoubleFlat  Pitch = iota
	B3Flat        Pitch = iota
	B3Natural     Pitch = iota
	B3Sharp       Pitch = iota
	B3DoubleSharp Pitch = iota

	C4DoubleFlat  Pitch = iota
	C4Flat        Pitch = iota
	C4Natural     Pitch = iota
	C4Sharp       Pitch = iota
	C4DoubleSharp Pitch = iota

	D4DoubleFlat  Pitch = iota
	D4Flat        Pitch = iota
	D4Natural     Pitch = iota
	D4Sharp       Pitch = iota
	D4DoubleSharp Pitch = iota

	E4DoubleFlat  Pitch = iota
	E4Flat        Pitch = iota
	E4Natural     Pitch = iota
	E4Sharp       Pitch = iota
	E4DoubleSharp Pitch = iota

	F4DoubleFlat  Pitch = iota
	F4Flat        Pitch = iota
	F4Natural     Pitch = iota
	F4Sharp       Pitch = iota
	F4DoubleSharp Pitch = iota

	G4DoubleFlat  Pitch = iota
	G4Flat        Pitch = iota
	G4Natural     Pitch = iota
	G4Sharp       Pitch = iota
	G4DoubleSharp Pitch = iota

	A4DoubleFlat  Pitch = iota
	A4Flat        Pitch = iota
	A4Natural     Pitch = iota
	A4Sharp       Pitch = iota
	A4DoubleSharp Pitch = iota

	B4DoubleFlat  Pitch = iota
	B4Flat        Pitch = iota
	B4Natural     Pitch = iota
	B4Sharp       Pitch = iota
	B4DoubleSharp Pitch = iota

	C5DoubleFlat  Pitch = iota
	C5Flat        Pitch = iota
	C5Natural     Pitch = iota
	C5Sharp       Pitch = iota
	C5DoubleSharp Pitch = iota

	D5DoubleFlat  Pitch = iota
	D5Flat        Pitch = iota
	D5Natural     Pitch = iota
	D5Sharp       Pitch = iota
	D5DoubleSharp Pitch = iota

	E5DoubleFlat  Pitch = iota
	E5Flat        Pitch = iota
	E5Natural     Pitch = iota
	E5Sharp       Pitch = iota
	E5DoubleSharp Pitch = iota

	F5DoubleFlat  Pitch = iota
	F5Flat        Pitch = iota
	F5Natural     Pitch = iota
	F5Sharp       Pitch = iota
	F5DoubleSharp Pitch = iota

	G5DoubleFlat  Pitch = iota
	G5Flat        Pitch = iota
	G5Natural     Pitch = iota
	G5Sharp       Pitch = iota
	G5DoubleSharp Pitch = iota

	A5DoubleFlat  Pitch = iota
	A5Flat        Pitch = iota
	A5Natural     Pitch = iota
	A5Sharp       Pitch = iota
	A5DoubleSharp Pitch = iota

	B5DoubleFlat  Pitch = iota
	B5Flat        Pitch = iota
	B5Natural     Pitch = iota
	B5Sharp       Pitch = iota
	B5DoubleSharp Pitch = iota

	C6DoubleFlat  Pitch = iota
	C6Flat        Pitch = iota
	C6Natural     Pitch = iota
	C6Sharp       Pitch = iota
	C6DoubleSharp Pitch = iota

	D6DoubleFlat  Pitch = iota
	D6Flat        Pitch = iota
	D6Natural     Pitch = iota
	D6Sharp       Pitch = iota
	D6DoubleSharp Pitch = iota

	E6DoubleFlat  Pitch = iota
	E6Flat        Pitch = iota
	E6Natural     Pitch = iota
	E6Sharp       Pitch = iota
	E6DoubleSharp Pitch = iota

	F6DoubleFlat  Pitch = iota
	F6Flat        Pitch = iota
	F6Natural     Pitch = iota
	F6Sharp       Pitch = iota
	F6DoubleSharp Pitch = iota

	G6DoubleFlat  Pitch = iota
	G6Flat        Pitch = iota
	G6Natural     Pitch = iota
	G6Sharp       Pitch = iota
	G6DoubleSharp Pitch = iota

	A6DoubleFlat  Pitch = iota
	A6Flat        Pitch = iota
	A6Natural     Pitch = iota
	A6Sharp       Pitch = iota
	A6DoubleSharp Pitch = iota

	B6DoubleFlat  Pitch = iota
	B6Flat        Pitch = iota
	B6Natural     Pitch = iota
	B6Sharp       Pitch = iota
	B6DoubleSharp Pitch = iota

	C7DoubleFlat  Pitch = iota
	C7Flat        Pitch = iota
	C7Natural     Pitch = iota
	C7Sharp       Pitch = iota
	C7DoubleSharp Pitch = iota

	D7DoubleFlat  Pitch = iota
	D7Flat        Pitch = iota
	D7Natural     Pitch = iota
	D7Sharp       Pitch = iota
	D7DoubleSharp Pitch = iota

	E7DoubleFlat  Pitch = iota
	E7Flat        Pitch = iota
	E7Natural     Pitch = iota
	E7Sharp       Pitch = iota
	E7DoubleSharp Pitch = iota

	F7DoubleFlat  Pitch = iota
	F7Flat        Pitch = iota
	F7Natural     Pitch = iota
	F7Sharp       Pitch = iota
	F7DoubleSharp Pitch = iota

	G7DoubleFlat  Pitch = iota
	G7Flat        Pitch = iota
	G7Natural     Pitch = iota
	G7Sharp       Pitch = iota
	G7DoubleSharp Pitch = iota

	A7DoubleFlat  Pitch = iota
	A7Flat        Pitch = iota
	A7Natural     Pitch = iota
	A7Sharp       Pitch = iota
	A7DoubleSharp Pitch = iota

	B7DoubleFlat  Pitch = iota
	B7Flat        Pitch = iota
	B7Natural     Pitch = iota
	B7Sharp       Pitch = iota
	B7DoubleSharp Pitch = iota

	C8DoubleFlat  Pitch = iota
	C8Flat        Pitch = iota
	C8Natural     Pitch = iota
	C8Sharp       Pitch = iota
	C8DoubleSharp Pitch = iota

	D8DoubleFlat  Pitch = iota
	D8Flat        Pitch = iota
	D8Natural     Pitch = iota
	D8Sharp       Pitch = iota
	D8DoubleSharp Pitch = iota

	E8DoubleFlat  Pitch = iota
	E8Flat        Pitch = iota
	E8Natural     Pitch = iota
	E8Sharp       Pitch = iota
	E8DoubleSharp Pitch = iota

	F8DoubleFlat  Pitch = iota
	F8Flat        Pitch = iota
	F8Natural     Pitch = iota
	F8Sharp       Pitch = iota
	F8DoubleSharp Pitch = iota

	G8DoubleFlat  Pitch = iota
	G8Flat        Pitch = iota
	G8Natural     Pitch = iota
	G8Sharp       Pitch = iota
	G8DoubleSharp Pitch = iota

	A8DoubleFlat  Pitch = iota
	A8Flat        Pitch = iota
	A8Natural     Pitch = iota
	A8Sharp       Pitch = iota
	A8DoubleSharp Pitch = iota

	B8DoubleFlat  Pitch = iota
	B8Flat        Pitch = iota
	B8Natural     Pitch = iota
	B8Sharp       Pitch = iota
	B8DoubleSharp Pitch = iota

	C9DoubleFlat  Pitch = iota
	C9Flat        Pitch = iota
	C9Natural     Pitch = iota
	C9Sharp       Pitch = iota
	C9DoubleSharp Pitch = iota

	D9DoubleFlat  Pitch = iota
	D9Flat        Pitch = iota
	D9Natural     Pitch = iota
	D9Sharp       Pitch = iota
	D9DoubleSharp Pitch = iota

	E9DoubleFlat  Pitch = iota
	E9Flat        Pitch = iota
	E9Natural     Pitch = iota
	E9Sharp       Pitch = iota
	E9DoubleSharp Pitch = iota

	F9DoubleFlat  Pitch = iota
	F9Flat        Pitch = iota
	F9Natural     Pitch = iota
	F9Sharp       Pitch = iota
	F9DoubleSharp Pitch = iota

	G9DoubleFlat  Pitch = iota
	G9Flat        Pitch = iota
	G9Natural     Pitch = iota
	G9Sharp       Pitch = iota
	G9DoubleSharp Pitch = iota

	A9DoubleFlat  Pitch = iota
	A9Flat        Pitch = iota
	A9Natural     Pitch = iota
	A9Sharp       Pitch = iota
	A9DoubleSharp Pitch = iota

	B9DoubleFlat  Pitch = iota
	B9Flat        Pitch = iota
	B9Natural     Pitch = iota
	B9Sharp       Pitch = iota
	B9DoubleSharp Pitch = iota

	C10DoubleFlat  Pitch = iota
	C10Flat        Pitch = iota
	C10Natural     Pitch = iota
	C10Sharp       Pitch = iota
	C10DoubleSharp Pitch = iota

	D10DoubleFlat  Pitch = iota
	D10Flat        Pitch = iota
	D10Natural     Pitch = iota
	D10Sharp       Pitch = iota
	D10DoubleSharp Pitch = iota

	E10DoubleFlat  Pitch = iota
	E10Flat        Pitch = iota
	E10Natural     Pitch = iota
	E10Sharp       Pitch = iota
	E10DoubleSharp Pitch = iota

	F10DoubleFlat  Pitch = iota
	F10Flat        Pitch = iota
	F10Natural     Pitch = iota
	F10Sharp       Pitch = iota
	F10DoubleSharp Pitch = iota

	G10DoubleFlat  Pitch = iota
	G10Flat        Pitch = iota
	G10Natural     Pitch = iota
	G10Sharp       Pitch = iota
	G10DoubleSharp Pitch = iota

	A10DoubleFlat  Pitch = iota
	A10Flat        Pitch = iota
	A10Natural     Pitch = iota
	A10Sharp       Pitch = iota
	A10DoubleSharp Pitch = iota

	B10DoubleFlat Pitch = iota
	B10Flat       Pitch = iota
	B10Natural    Pitch = iota
)

var pitchStringifieds = map[Pitch]string{
	Invalid:    "Invalid",
	CN1Natural: "CN1Natural", CN1Sharp: "CN1Sharp", CN1DoubleSharp: "CN1DoubleSharp",
	DN1DoubleFlat: "DN1DoubleFlat", DN1Flat: "DN1Flat", DN1Natural: "DN1Natural", DN1Sharp: "DN1Sharp", DN1DoubleSharp: "DN1DoubleSharp",
	EN1DoubleFlat: "EN1DoubleFlat", EN1Flat: "EN1Flat", EN1Natural: "EN1Natural", EN1Sharp: "EN1Sharp", EN1DoubleSharp: "EN1DoubleSharp",
	FN1DoubleFlat: "FN1DoubleFlat", FN1Flat: "FN1Flat", FN1Natural: "FN1Natural", FN1Sharp: "FN1Sharp", FN1DoubleSharp: "FN1DoubleSharp",
	GN1DoubleFlat: "GN1DoubleFlat", GN1Flat: "GN1Flat", GN1Natural: "GN1Natural", GN1Sharp: "GN1Sharp", GN1DoubleSharp: "GN1DoubleSharp",
	AN1DoubleFlat: "AN1DoubleFlat", AN1Flat: "AN1Flat", AN1Natural: "AN1Natural", AN1Sharp: "AN1Sharp", AN1DoubleSharp: "AN1DoubleSharp",
	BN1DoubleFlat: "BN1DoubleFlat", BN1Flat: "BN1Flat", BN1Natural: "BN1Natural", BN1Sharp: "BN1Sharp", BN1DoubleSharp: "BN1DoubleSharp",

	C0DoubleFlat: "C0DoubleFlat", C0Flat: "C0Flat", C0Natural: "C0Natural", C0Sharp: "C0Sharp", C0DoubleSharp: "C0DoubleSharp",
	D0DoubleFlat: "D0DoubleFlat", D0Flat: "D0Flat", D0Natural: "D0Natural", D0Sharp: "D0Sharp", D0DoubleSharp: "D0DoubleSharp",
	E0DoubleFlat: "E0DoubleFlat", E0Flat: "E0Flat", E0Natural: "E0Natural", E0Sharp: "E0Sharp", E0DoubleSharp: "E0DoubleSharp",
	F0DoubleFlat: "F0DoubleFlat", F0Flat: "F0Flat", F0Natural: "F0Natural", F0Sharp: "F0Sharp", F0DoubleSharp: "F0DoubleSharp",
	G0DoubleFlat: "G0DoubleFlat", G0Flat: "G0Flat", G0Natural: "G0Natural", G0Sharp: "G0Sharp", G0DoubleSharp: "G0DoubleSharp",
	A0DoubleFlat: "A0DoubleFlat", A0Flat: "A0Flat", A0Natural: "A0Natural", A0Sharp: "A0Sharp", A0DoubleSharp: "A0DoubleSharp",
	B0DoubleFlat: "B0DoubleFlat", B0Flat: "B0Flat", B0Natural: "B0Natural", B0Sharp: "B0Sharp", B0DoubleSharp: "B0DoubleSharp",

	C1DoubleFlat: "C1DoubleFlat", C1Flat: "C1Flat", C1Natural: "C1Natural", C1Sharp: "C1Sharp", C1DoubleSharp: "C1DoubleSharp",
	D1DoubleFlat: "D1DoubleFlat", D1Flat: "D1Flat", D1Natural: "D1Natural", D1Sharp: "D1Sharp", D1DoubleSharp: "D1DoubleSharp",
	E1DoubleFlat: "E1DoubleFlat", E1Flat: "E1Flat", E1Natural: "E1Natural", E1Sharp: "E1Sharp", E1DoubleSharp: "E1DoubleSharp",
	F1DoubleFlat: "F1DoubleFlat", F1Flat: "F1Flat", F1Natural: "F1Natural", F1Sharp: "F1Sharp", F1DoubleSharp: "F1DoubleSharp",
	G1DoubleFlat: "G1DoubleFlat", G1Flat: "G1Flat", G1Natural: "G1Natural", G1Sharp: "G1Sharp", G1DoubleSharp: "G1DoubleSharp",
	A1DoubleFlat: "A1DoubleFlat", A1Flat: "A1Flat", A1Natural: "A1Natural", A1Sharp: "A1Sharp", A1DoubleSharp: "A1DoubleSharp",
	B1DoubleFlat: "B1DoubleFlat", B1Flat: "B1Flat", B1Natural: "B1Natural", B1Sharp: "B1Sharp", B1DoubleSharp: "B1DoubleSharp",

	C2DoubleFlat: "C2DoubleFlat", C2Flat: "C2Flat", C2Natural: "C2Natural", C2Sharp: "C2Sharp", C2DoubleSharp: "C2DoubleSharp",
	D2DoubleFlat: "D2DoubleFlat", D2Flat: "D2Flat", D2Natural: "D2Natural", D2Sharp: "D2Sharp", D2DoubleSharp: "D2DoubleSharp",
	E2DoubleFlat: "E2DoubleFlat", E2Flat: "E2Flat", E2Natural: "E2Natural", E2Sharp: "E2Sharp", E2DoubleSharp: "E2DoubleSharp",
	F2DoubleFlat: "F2DoubleFlat", F2Flat: "F2Flat", F2Natural: "F2Natural", F2Sharp: "F2Sharp", F2DoubleSharp: "F2DoubleSharp",
	G2DoubleFlat: "G2DoubleFlat", G2Flat: "G2Flat", G2Natural: "G2Natural", G2Sharp: "G2Sharp", G2DoubleSharp: "G2DoubleSharp",
	A2DoubleFlat: "A2DoubleFlat", A2Flat: "A2Flat", A2Natural: "A2Natural", A2Sharp: "A2Sharp", A2DoubleSharp: "A2DoubleSharp",
	B2DoubleFlat: "B2DoubleFlat", B2Flat: "B2Flat", B2Natural: "B2Natural", B2Sharp: "B2Sharp", B2DoubleSharp: "B2DoubleSharp",

	C3DoubleFlat: "C3DoubleFlat", C3Flat: "C3Flat", C3Natural: "C3Natural", C3Sharp: "C3Sharp", C3DoubleSharp: "C3DoubleSharp",
	D3DoubleFlat: "D3DoubleFlat", D3Flat: "D3Flat", D3Natural: "D3Natural", D3Sharp: "D3Sharp", D3DoubleSharp: "D3DoubleSharp",
	E3DoubleFlat: "E3DoubleFlat", E3Flat: "E3Flat", E3Natural: "E3Natural", E3Sharp: "E3Sharp", E3DoubleSharp: "E3DoubleSharp",
	F3DoubleFlat: "F3DoubleFlat", F3Flat: "F3Flat", F3Natural: "F3Natural", F3Sharp: "F3Sharp", F3DoubleSharp: "F3DoubleSharp",
	G3DoubleFlat: "G3DoubleFlat", G3Flat: "G3Flat", G3Natural: "G3Natural", G3Sharp: "G3Sharp", G3DoubleSharp: "G3DoubleSharp",
	A3DoubleFlat: "A3DoubleFlat", A3Flat: "A3Flat", A3Natural: "A3Natural", A3Sharp: "A3Sharp", A3DoubleSharp: "A3DoubleSharp",
	B3DoubleFlat: "B3DoubleFlat", B3Flat: "B3Flat", B3Natural: "B3Natural", B3Sharp: "B3Sharp", B3DoubleSharp: "B3DoubleSharp",

	C4DoubleFlat: "C4DoubleFlat", C4Flat: "C4Flat", C4Natural: "C4Natural", C4Sharp: "C4Sharp", C4DoubleSharp: "C4DoubleSharp",
	D4DoubleFlat: "D4DoubleFlat", D4Flat: "D4Flat", D4Natural: "D4Natural", D4Sharp: "D4Sharp", D4DoubleSharp: "D4DoubleSharp",
	E4DoubleFlat: "E4DoubleFlat", E4Flat: "E4Flat", E4Natural: "E4Natural", E4Sharp: "E4Sharp", E4DoubleSharp: "E4DoubleSharp",
	F4DoubleFlat: "F4DoubleFlat", F4Flat: "F4Flat", F4Natural: "F4Natural", F4Sharp: "F4Sharp", F4DoubleSharp: "F4DoubleSharp",
	G4DoubleFlat: "G4DoubleFlat", G4Flat: "G4Flat", G4Natural: "G4Natural", G4Sharp: "G4Sharp", G4DoubleSharp: "G4DoubleSharp",
	A4DoubleFlat: "A4DoubleFlat", A4Flat: "A4Flat", A4Natural: "A4Natural", A4Sharp: "A4Sharp", A4DoubleSharp: "A4DoubleSharp",
	B4DoubleFlat: "B4DoubleFlat", B4Flat: "B4Flat", B4Natural: "B4Natural", B4Sharp: "B4Sharp", B4DoubleSharp: "B4DoubleSharp",

	C5DoubleFlat: "C5DoubleFlat", C5Flat: "C5Flat", C5Natural: "C5Natural", C5Sharp: "C5Sharp", C5DoubleSharp: "C5DoubleSharp",
	D5DoubleFlat: "D5DoubleFlat", D5Flat: "D5Flat", D5Natural: "D5Natural", D5Sharp: "D5Sharp", D5DoubleSharp: "D5DoubleSharp",
	E5DoubleFlat: "E5DoubleFlat", E5Flat: "E5Flat", E5Natural: "E5Natural", E5Sharp: "E5Sharp", E5DoubleSharp: "E5DoubleSharp",
	F5DoubleFlat: "F5DoubleFlat", F5Flat: "F5Flat", F5Natural: "F5Natural", F5Sharp: "F5Sharp", F5DoubleSharp: "F5DoubleSharp",
	G5DoubleFlat: "G5DoubleFlat", G5Flat: "G5Flat", G5Natural: "G5Natural", G5Sharp: "G5Sharp", G5DoubleSharp: "G5DoubleSharp",
	A5DoubleFlat: "A5DoubleFlat", A5Flat: "A5Flat", A5Natural: "A5Natural", A5Sharp: "A5Sharp", A5DoubleSharp: "A5DoubleSharp",
	B5DoubleFlat: "B5DoubleFlat", B5Flat: "B5Flat", B5Natural: "B5Natural", B5Sharp: "B5Sharp", B5DoubleSharp: "B5DoubleSharp",

	C6DoubleFlat: "C6DoubleFlat", C6Flat: "C6Flat", C6Natural: "C6Natural", C6Sharp: "C6Sharp", C6DoubleSharp: "C6DoubleSharp",
	D6DoubleFlat: "D6DoubleFlat", D6Flat: "D6Flat", D6Natural: "D6Natural", D6Sharp: "D6Sharp", D6DoubleSharp: "D6DoubleSharp",
	E6DoubleFlat: "E6DoubleFlat", E6Flat: "E6Flat", E6Natural: "E6Natural", E6Sharp: "E6Sharp", E6DoubleSharp: "E6DoubleSharp",
	F6DoubleFlat: "F6DoubleFlat", F6Flat: "F6Flat", F6Natural: "F6Natural", F6Sharp: "F6Sharp", F6DoubleSharp: "F6DoubleSharp",
	G6DoubleFlat: "G6DoubleFlat", G6Flat: "G6Flat", G6Natural: "G6Natural", G6Sharp: "G6Sharp", G6DoubleSharp: "G6DoubleSharp",
	A6DoubleFlat: "A6DoubleFlat", A6Flat: "A6Flat", A6Natural: "A6Natural", A6Sharp: "A6Sharp", A6DoubleSharp: "A6DoubleSharp",
	B6DoubleFlat: "B6DoubleFlat", B6Flat: "B6Flat", B6Natural: "B6Natural", B6Sharp: "B6Sharp", B6DoubleSharp: "B6DoubleSharp",

	C7DoubleFlat: "C7DoubleFlat", C7Flat: "C7Flat", C7Natural: "C7Natural", C7Sharp: "C7Sharp", C7DoubleSharp: "C7DoubleSharp",
	D7DoubleFlat: "D7DoubleFlat", D7Flat: "D7Flat", D7Natural: "D7Natural", D7Sharp: "D7Sharp", D7DoubleSharp: "D7DoubleSharp",
	E7DoubleFlat: "E7DoubleFlat", E7Flat: "E7Flat", E7Natural: "E7Natural", E7Sharp: "E7Sharp", E7DoubleSharp: "E7DoubleSharp",
	F7DoubleFlat: "F7DoubleFlat", F7Flat: "F7Flat", F7Natural: "F7Natural", F7Sharp: "F7Sharp", F7DoubleSharp: "F7DoubleSharp",
	G7DoubleFlat: "G7DoubleFlat", G7Flat: "G7Flat", G7Natural: "G7Natural", G7Sharp: "G7Sharp", G7DoubleSharp: "G7DoubleSharp",
	A7DoubleFlat: "A7DoubleFlat", A7Flat: "A7Flat", A7Natural: "A7Natural", A7Sharp: "A7Sharp", A7DoubleSharp: "A7DoubleSharp",
	B7DoubleFlat: "B7DoubleFlat", B7Flat: "B7Flat", B7Natural: "B7Natural", B7Sharp: "B7Sharp", B7DoubleSharp: "B7DoubleSharp",

	C8DoubleFlat: "C8DoubleFlat", C8Flat: "C8Flat", C8Natural: "C8Natural", C8Sharp: "C8Sharp", C8DoubleSharp: "C8DoubleSharp",
	D8DoubleFlat: "D8DoubleFlat", D8Flat: "D8Flat", D8Natural: "D8Natural", D8Sharp: "D8Sharp", D8DoubleSharp: "D8DoubleSharp",
	E8DoubleFlat: "E8DoubleFlat", E8Flat: "E8Flat", E8Natural: "E8Natural", E8Sharp: "E8Sharp", E8DoubleSharp: "E8DoubleSharp",
	F8DoubleFlat: "F8DoubleFlat", F8Flat: "F8Flat", F8Natural: "F8Natural", F8Sharp: "F8Sharp", F8DoubleSharp: "F8DoubleSharp",
	G8DoubleFlat: "G8DoubleFlat", G8Flat: "G8Flat", G8Natural: "G8Natural", G8Sharp: "G8Sharp", G8DoubleSharp: "G8DoubleSharp",
	A8DoubleFlat: "A8DoubleFlat", A8Flat: "A8Flat", A8Natural: "A8Natural", A8Sharp: "A8Sharp", A8DoubleSharp: "A8DoubleSharp",
	B8DoubleFlat: "B8DoubleFlat", B8Flat: "B8Flat", B8Natural: "B8Natural", B8Sharp: "B8Sharp", B8DoubleSharp: "B8DoubleSharp",

	C9DoubleFlat: "C9DoubleFlat", C9Flat: "C9Flat", C9Natural: "C9Natural", C9Sharp: "C9Sharp", C9DoubleSharp: "C9DoubleSharp",
	D9DoubleFlat: "D9DoubleFlat", D9Flat: "D9Flat", D9Natural: "D9Natural", D9Sharp: "D9Sharp", D9DoubleSharp: "D9DoubleSharp",
	E9DoubleFlat: "E9DoubleFlat", E9Flat: "E9Flat", E9Natural: "E9Natural", E9Sharp: "E9Sharp", E9DoubleSharp: "E9DoubleSharp",
	F9DoubleFlat: "F9DoubleFlat", F9Flat: "F9Flat", F9Natural: "F9Natural", F9Sharp: "F9Sharp", F9DoubleSharp: "F9DoubleSharp",
	G9DoubleFlat: "G9DoubleFlat", G9Flat: "G9Flat", G9Natural: "G9Natural", G9Sharp: "G9Sharp", G9DoubleSharp: "G9DoubleSharp",
	A9DoubleFlat: "A9DoubleFlat", A9Flat: "A9Flat", A9Natural: "A9Natural", A9Sharp: "A9Sharp", A9DoubleSharp: "A9DoubleSharp",
	B9DoubleFlat: "B9DoubleFlat", B9Flat: "B9Flat", B9Natural: "B9Natural", B9Sharp: "B9Sharp", B9DoubleSharp: "B9DoubleSharp",

	C10DoubleFlat: "C10DoubleFlat", C10Flat: "C10Flat", C10Natural: "C10Natural", C10Sharp: "C10Sharp", C10DoubleSharp: "C10DoubleSharp",
	D10DoubleFlat: "D10DoubleFlat", D10Flat: "D10Flat", D10Natural: "D10Natural", D10Sharp: "D10Sharp", D10DoubleSharp: "D10DoubleSharp",
	E10DoubleFlat: "E10DoubleFlat", E10Flat: "E10Flat", E10Natural: "E10Natural", E10Sharp: "E10Sharp", E10DoubleSharp: "E10DoubleSharp",
	F10DoubleFlat: "F10DoubleFlat", F10Flat: "F10Flat", F10Natural: "F10Natural", F10Sharp: "F10Sharp", F10DoubleSharp: "F10DoubleSharp",
	G10DoubleFlat: "G10DoubleFlat", G10Flat: "G10Flat", G10Natural: "G10Natural", G10Sharp: "G10Sharp", G10DoubleSharp: "G10DoubleSharp",
	A10DoubleFlat: "A10DoubleFlat", A10Flat: "A10Flat", A10Natural: "A10Natural", A10Sharp: "A10Sharp", A10DoubleSharp: "A10DoubleSharp",
	B10DoubleFlat: "B10DoubleFlat", B10Flat: "B10Flat", B10Natural: "B10Natural",
}

// AllPitches return all pitches
func AllPitches() []Pitch {
	pitches := make([]Pitch, 0)
	for v := range pitchStringifieds {
		if v != Invalid {
			pitches = append(pitches, v)
		}
	}

	sort.SliceStable(pitches, func(i, j int) bool {
		return pitches[i] < pitches[j]
	})
	return pitches
}

// PianoPitches returns piano pitches
func PianoPitches(keys int) []Pitch {
	pitches := make([]Pitch, 0)
	for _, v := range AllPitches() {
		switch {
		case v.MIDINoteNumber() < limit.MIDINoteNumberMin,
			v.MIDINoteNumber() > limit.MIDINoteNumberMax:
			continue
		case keys == 88 && (v.MIDINoteNumber() < A0Natural.MIDINoteNumber() ||
			v.MIDINoteNumber() > C8Natural.MIDINoteNumber()):
			continue
		case keys == 37 && (v.MIDINoteNumber() < C4Natural.MIDINoteNumber() ||
			v.MIDINoteNumber() > C7Natural.MIDINoteNumber()):
			continue
		case keys == 25 && (v.MIDINoteNumber() < C4Natural.MIDINoteNumber() ||
			v.MIDINoteNumber() > C6Natural.MIDINoteNumber()):
			continue
		}

		pitches = append(pitches, v)
	}

	sort.SliceStable(pitches, func(i, j int) bool {
		return pitches[i] < pitches[j]
	})
	return pitches
}

// SimplePianoPitches returns simple piano pitches
func SimplePianoPitches(keys int) []Pitch {
	pitches := make([]Pitch, 0)
	uniqueMap := make(map[Pitch]struct{})
	for _, v := range PianoPitches(keys) {
		switch v.Note().Accidental() {
		case accidental.TripleFlat,
			accidental.TripleSharp,
			accidental.DoubleFlat,
			accidental.DoubleSharp:
			continue
		}

		normalized := v.Normalize()
		if _, found := uniqueMap[normalized]; !found {
			uniqueMap[normalized] = struct{}{}
			pitches = append(pitches, normalized)
		}
	}

	sort.SliceStable(pitches, func(i, j int) bool {
		return pitches[i] < pitches[j]
	})
	return pitches
}

// String returns stringified pitch name
func (p Pitch) String() string {
	return pitchStringifieds[p]
}
