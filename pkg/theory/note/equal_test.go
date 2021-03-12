package note_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNote_Equal(t *testing.T) {
	require.True(t, note.CFlat.Equal(note.BNatural))
	require.True(t, note.CNatural.Equal(note.BSharp))
	require.True(t, note.CSharp.Equal(note.DFlat))

	require.True(t, note.DFlat.Equal(note.CSharp))
	require.True(t, note.DSharp.Equal(note.EFlat))

	require.True(t, note.EFlat.Equal(note.DSharp))
	require.True(t, note.ENatural.Equal(note.FFlat))
	require.True(t, note.ESharp.Equal(note.FNatural))

	require.True(t, note.FFlat.Equal(note.ENatural))
	require.True(t, note.FNatural.Equal(note.ESharp))
	require.True(t, note.FSharp.Equal(note.GFlat))

	require.True(t, note.GFlat.Equal(note.FSharp))
	require.True(t, note.GSharp.Equal(note.AFlat))

	require.True(t, note.AFlat.Equal(note.GSharp))
	require.True(t, note.ASharp.Equal(note.BFlat))
}
