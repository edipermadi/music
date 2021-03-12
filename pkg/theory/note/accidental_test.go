package note_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/edipermadi/music/pkg/theory/note/accidental"
	"github.com/stretchr/testify/require"
)

func TestNote_Accidental(t *testing.T) {
	require.Equal(t, accidental.TripleFlat, note.CTripleFlat.Accidental())
	require.Equal(t, accidental.DoubleFlat, note.CDoubleFlat.Accidental())
	require.Equal(t, accidental.Flat, note.CFlat.Accidental())
	require.Equal(t, accidental.Natural, note.CNatural.Accidental())
	require.Equal(t, accidental.Sharp, note.CSharp.Accidental())
	require.Equal(t, accidental.DoubleSharp, note.CDoubleSharp.Accidental())
	require.Equal(t, accidental.TripleSharp, note.CTripleSharp.Accidental())
}

func TestAccidental_Name(t *testing.T) {
	for _, givenAccidental := range accidental.AllAccidentals() {
		if givenAccidental.Natural() {
			require.Empty(t, givenAccidental.Name(), fmt.Sprintf("failed on %s", givenAccidental))
		} else {
			require.NotEmpty(t, givenAccidental.Name(), fmt.Sprintf("failed on %s", givenAccidental))
		}

	}
}
