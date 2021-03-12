package chordtype_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/stretchr/testify/require"
)

func TestChord_Intervals(t *testing.T) {
	for _, chordType := range chordtype.AllTypes() {
		require.NotEmpty(t, chordType.Intervals())
	}
}
