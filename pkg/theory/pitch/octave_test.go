package pitch_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestPitch_Octave(t *testing.T) {
	require.Equal(t, -1, pitch.CN1Natural.Octave())
	require.Equal(t, 10, pitch.C10Natural.Octave())
}
