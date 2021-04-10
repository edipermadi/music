package note_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/mode"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNotes_CircleOfFifth(t *testing.T) {
	circleOfFifth := mode.CNaturalIonian.Notes().CircleOfFifth()
	require.True(t, circleOfFifth[note.CNatural])
	require.False(t, circleOfFifth[note.CSharp])
	require.True(t, circleOfFifth[note.DNatural])
	require.False(t, circleOfFifth[note.DSharp])
	require.True(t, circleOfFifth[note.ENatural])
	require.True(t, circleOfFifth[note.FNatural])
	require.False(t, circleOfFifth[note.FSharp])
	require.True(t, circleOfFifth[note.GNatural])
	require.False(t, circleOfFifth[note.GSharp])
	require.True(t, circleOfFifth[note.ANatural])
	require.False(t, circleOfFifth[note.ASharp])
	require.True(t, circleOfFifth[note.BNatural])
}

func TestCircleOfFifth_HasLeftSibling(t *testing.T) {
	circleOfFifth := mode.CNaturalIonian.Notes().CircleOfFifth()
	require.True(t, circleOfFifth.HasLeftSibling(note.CNatural))
	require.True(t, circleOfFifth.HasLeftSibling(note.DNatural))
	require.True(t, circleOfFifth.HasLeftSibling(note.ENatural))
	require.False(t, circleOfFifth.HasLeftSibling(note.FNatural))
	require.True(t, circleOfFifth.HasLeftSibling(note.GNatural))
	require.True(t, circleOfFifth.HasLeftSibling(note.ANatural))
	require.True(t, circleOfFifth.HasLeftSibling(note.BNatural))
}

func TestCircleOfFifth_HasRightSibling(t *testing.T) {
	circleOfFifth := mode.CNaturalIonian.Notes().CircleOfFifth()
	require.True(t, circleOfFifth.HasRightSibling(note.CNatural))
	require.True(t, circleOfFifth.HasRightSibling(note.DNatural))
	require.True(t, circleOfFifth.HasRightSibling(note.ENatural))
	require.True(t, circleOfFifth.HasRightSibling(note.FNatural))
	require.True(t, circleOfFifth.HasRightSibling(note.GNatural))
	require.True(t, circleOfFifth.HasRightSibling(note.ANatural))
	require.False(t, circleOfFifth.HasRightSibling(note.BNatural))
}
