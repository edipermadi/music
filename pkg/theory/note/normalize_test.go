package note_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNote_Normalize(t *testing.T) {
	normalizedNotes := map[note.Note]note.Note{
		note.CFlat:        note.BNatural,
		note.CDoubleSharp: note.DNatural,
		note.DDoubleFlat:  note.CNatural,
		note.DDoubleSharp: note.ENatural,
		note.EDoubleFlat:  note.DNatural,
		note.ESharp:       note.FNatural,
		note.FFlat:        note.ENatural,
		note.FDoubleSharp: note.GNatural,
		note.GDoubleFlat:  note.FNatural,
		note.GDoubleSharp: note.ANatural,
		note.ADoubleFlat:  note.GNatural,
		note.ADoubleSharp: note.BNatural,
		note.BDoubleFlat:  note.ANatural,
		note.BSharp:       note.CNatural,
	}

	for src, dst := range normalizedNotes {
		require.Equal(t, dst, src.Normalize())
	}

}
