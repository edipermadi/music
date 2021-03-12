package pitch_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestPitch_MIDINoteNumber(t *testing.T) {
	require.Equal(t, 0, pitch.CN1Natural.MIDINoteNumber())
	require.Equal(t, 127, pitch.G9Natural.MIDINoteNumber())
	require.Equal(t, -1, pitch.G9Sharp.MIDINoteNumber())
}
