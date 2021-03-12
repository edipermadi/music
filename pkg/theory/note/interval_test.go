package note_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/interval"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNote_Interval(t *testing.T) {
	type testCase struct {
		Source      note.Note
		Destination note.Note
		Interval    interval.Interval
	}

	testCases := []testCase{
		{note.CNatural, note.CNatural, interval.PerfectUnison},
		{note.CNatural, note.CSharp, interval.MinorSecond},
		{note.CNatural, note.DNatural, interval.MajorSecond},
		{note.CNatural, note.DSharp, interval.MinorThird},
		{note.CNatural, note.ENatural, interval.MajorThird},
		{note.CNatural, note.FNatural, interval.PerfectFourth},
		{note.CNatural, note.FSharp, interval.AugmentedFourth},
		{note.CNatural, note.GNatural, interval.PerfectFifth},
		{note.CNatural, note.GSharp, interval.MinorSixth},
		{note.CNatural, note.ANatural, interval.MajorSixth},
		{note.CNatural, note.ASharp, interval.MinorSeventh},
		{note.CNatural, note.BNatural, interval.MajorSeventh},
		{note.CNatural, note.CNatural, interval.PerfectOctave},
	}

	for _, tc := range testCases {
		intervals, err := tc.Source.Interval(tc.Destination)
		require.NoError(t, err)
		require.Contains(t, intervals, tc.Interval)
	}
}
