package signature

import (
	"fmt"

	"github.com/edipermadi/music/pkg/theory/note"
)

var mapSignatureToNotes = map[Signature]note.Notes{
	CNaturalMajor: {note.CNatural, note.DNatural, note.ENatural, note.FNatural, note.GNatural, note.ANatural, note.BNatural},
	GNaturalMajor: {note.GNatural, note.ANatural, note.BNatural, note.CNatural, note.DNatural, note.ENatural, note.FSharp},
	DNaturalMajor: {note.DNatural, note.ENatural, note.FSharp, note.GNatural, note.ANatural, note.BNatural, note.CSharp},
	ANaturalMajor: {note.ANatural, note.BNatural, note.CSharp, note.DNatural, note.ENatural, note.FSharp, note.GSharp},
	ENaturalMajor: {note.ENatural, note.FSharp, note.GSharp, note.ANatural, note.BNatural, note.CSharp, note.DSharp},
	BNaturalMajor: {note.BNatural, note.CSharp, note.DSharp, note.ENatural, note.FSharp, note.GSharp, note.ASharp},
	FSharpMajor:   {note.FSharp, note.GSharp, note.ASharp, note.BNatural, note.CSharp, note.DSharp, note.ESharp},
	CSharpMajor:   {note.CSharp, note.DSharp, note.ESharp, note.FSharp, note.GSharp, note.ASharp, note.BSharp},

	FNaturalMajor: {note.FNatural, note.GNatural, note.ANatural, note.BFlat, note.CNatural, note.DNatural, note.ENatural},
	BFlatMajor:    {note.BFlat, note.CNatural, note.DNatural, note.EFlat, note.FNatural, note.GNatural, note.ANatural},
	EFlatMajor:    {note.EFlat, note.FNatural, note.GNatural, note.AFlat, note.BFlat, note.CNatural, note.DNatural},
	AFlatMajor:    {note.AFlat, note.BFlat, note.CNatural, note.DFlat, note.EFlat, note.FNatural, note.GNatural},
	DFlatMajor:    {note.DFlat, note.EFlat, note.FNatural, note.GFlat, note.AFlat, note.BFlat, note.CNatural},
	GFlatMajor:    {note.GFlat, note.AFlat, note.BFlat, note.CFlat, note.DFlat, note.EFlat, note.FNatural},
}

// Notes return signature notes
func (s Signature) Notes() note.Notes {
	if v, found := mapSignatureToNotes[s]; found {
		return v
	}

	panic(fmt.Errorf("no such notes for signature %s", s.String()))
}
