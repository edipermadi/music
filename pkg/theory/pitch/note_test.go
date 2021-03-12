package pitch_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestPitch_Note(t *testing.T) {
	for _, p := range pitch.AllPitches() {
		require.NotEmpty(t, p.Note().String())
	}
}
