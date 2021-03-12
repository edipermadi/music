package interval_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/interval"
	"github.com/edipermadi/music/pkg/theory/limit"
	"github.com/stretchr/testify/require"
)

func Test_FromSemitones(t *testing.T) {
	for i := 0; i <= limit.SemitonesInOneOctave; i++ {
		intervals, err := interval.FromSemitones(i)
		require.NoError(t, err)
		require.NotEmpty(t, intervals)
		t.Logf("%d semitones corresponds to %#v", i, intervals)
	}
}
