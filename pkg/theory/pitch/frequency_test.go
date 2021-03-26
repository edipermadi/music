package pitch_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestPitch_Frequency(t *testing.T) {
	require.Equal(t, 523.251131, pitch.C5Natural.Frequency())
	require.Equal(t, 261.625565, pitch.C4Natural.Frequency())
}
