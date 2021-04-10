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

func TestNote_NextFifth(t *testing.T) {
	require.Equal(t, note.GNatural, note.CNatural.NextFifth())
	require.Equal(t, note.DNatural, note.GNatural.NextFifth())
	require.Equal(t, note.ANatural, note.DNatural.NextFifth())
	require.Equal(t, note.ENatural, note.ANatural.NextFifth())
	require.Equal(t, note.BNatural, note.ENatural.NextFifth())
	require.Equal(t, note.GFlat, note.BNatural.NextFifth())
	require.Equal(t, note.DFlat, note.GFlat.NextFifth())
	require.Equal(t, note.AFlat, note.DFlat.NextFifth())
	require.Equal(t, note.EFlat, note.AFlat.NextFifth())
	require.Equal(t, note.BFlat, note.EFlat.NextFifth())
	require.Equal(t, note.FNatural, note.BFlat.NextFifth())
	require.Equal(t, note.CNatural, note.FNatural.NextFifth())
}

func TestNote_PreviousFifth(t *testing.T) {
	require.Equal(t, note.FNatural, note.CNatural.PreviousFifth())
	require.Equal(t, note.CNatural, note.GNatural.PreviousFifth())
	require.Equal(t, note.GNatural, note.DNatural.PreviousFifth())
	require.Equal(t, note.DNatural, note.ANatural.PreviousFifth())
	require.Equal(t, note.ANatural, note.ENatural.PreviousFifth())
	require.Equal(t, note.ENatural, note.BNatural.PreviousFifth())
	require.Equal(t, note.BNatural, note.GFlat.PreviousFifth())
	require.Equal(t, note.GFlat, note.DFlat.PreviousFifth())
	require.Equal(t, note.DFlat, note.AFlat.PreviousFifth())
	require.Equal(t, note.AFlat, note.EFlat.PreviousFifth())
	require.Equal(t, note.EFlat, note.BFlat.PreviousFifth())
	require.Equal(t, note.BFlat, note.FNatural.PreviousFifth())
}
