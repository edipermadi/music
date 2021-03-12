package modetype_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/mode/modetype"
	"github.com/stretchr/testify/require"
)

func TestMode_Brighten(t *testing.T) {
	require.Equal(t, modetype.Phrygian, modetype.Locrian.Brighten())
	require.Equal(t, modetype.Aeolian, modetype.Phrygian.Brighten())
	require.Equal(t, modetype.Dorian, modetype.Aeolian.Brighten())
	require.Equal(t, modetype.Mixolydian, modetype.Dorian.Brighten())
	require.Equal(t, modetype.Ionian, modetype.Mixolydian.Brighten())
	require.Equal(t, modetype.Lydian, modetype.Ionian.Brighten())
	require.Equal(t, modetype.Lydian, modetype.Lydian.Brighten())
}

func TestMode_Darken(t *testing.T) {
	require.Equal(t, modetype.Ionian, modetype.Lydian.Darken())
	require.Equal(t, modetype.Mixolydian, modetype.Ionian.Darken())
	require.Equal(t, modetype.Dorian, modetype.Mixolydian.Darken())
	require.Equal(t, modetype.Aeolian, modetype.Dorian.Darken())
	require.Equal(t, modetype.Phrygian, modetype.Aeolian.Darken())
	require.Equal(t, modetype.Locrian, modetype.Phrygian.Darken())
	require.Equal(t, modetype.Locrian, modetype.Locrian.Darken())
}
