package chord_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/chord"
	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestChord_AllChords(t *testing.T) {
	require.NotEmpty(t, chord.AllChords())
}

func TestChord_String(t *testing.T) {
	for _, c := range chord.AllChords() {
		require.NotEmpty(t, c.String())
	}
}

func TestChord_Name(t *testing.T) {
	for _, c := range chord.AllChords() {
		name := c.Name()
		require.NotEmpty(t, name)
		t.Logf("name of %s code is %s", c, name)
	}
}

func TestChord_Notes(t *testing.T) {
	type testCase struct {
		GivenChord    chord.Chord
		ExpectedNotes note.Notes
	}
	testCases := []testCase{
		// Power Chord
		{
			GivenChord:    chord.New(note.CNatural, chordtype.PowerChord),
			ExpectedNotes: note.Notes{note.CNatural, note.GNatural},
		},

		// Triad
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Diminished),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorDoubleFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Minor),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorAddNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorAddEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorAddSharpNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.DSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorAddFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.FNatural, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorAddSharpFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.FSharp, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.AFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Major),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorAddNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorAddEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorAddSharpNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.DSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorDoubleSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorAddFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.FNatural, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorAddSharpFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.FSharp, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Augmented),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedSecondDoubleFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedSecondFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedSecondFlatFifthAddSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GFlat, note.GSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedSecond),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedSecondSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedFourthFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedFourthSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.SuspendedFourthDoubleSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Phrygian),
			ExpectedNotes: note.Notes{note.CNatural, note.DFlat, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.PhrygianAddSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.DFlat, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Lydian),
			ExpectedNotes: note.Notes{note.CNatural, note.FSharp, note.GNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.LydianMajorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.FSharp, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Locrian),
			ExpectedNotes: note.Notes{note.CNatural, note.DFlat, note.GFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.Quartal),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.QuartalAugmented),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.BNatural},
		},

		// Sixth
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSixth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSixthAddNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.ANatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSixthAddFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.ANatural, note.DFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthAddNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.ANatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthAddFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.ANatural, note.DFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GFlat, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthSuspendedSecond),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthSuspendedSecondFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GFlat, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthSuspendedSecondDoubleFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSixthSuspendedFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GNatural, note.ANatural},
		},

		// Seventh
		{
			GivenChord:    chord.New(note.CNatural, chordtype.FullDiminishedSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GFlat, note.BDoubleFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.HalfDiminishedSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GFlat, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventhDoubleFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.FNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventhFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.DFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.DNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.DNatural, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventhAddEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventhAddThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventhAddSharpEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BFlat, note.FSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorSeventhSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GSharp, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorMajorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorMajorNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BNatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorMajorEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BNatural, note.DNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorMajorThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BNatural, note.DNatural, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorMajorSeventhAddEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MinorMajorSeventhAddThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GNatural, note.BNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSuspendedSecondFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural, note.BDoubleFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSuspendedSecond),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantNinthSuspendedSecond),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural, note.BFlat, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSuspendedFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantNinthSuspendedFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GNatural, note.BFlat, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GFlat, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantNinthSharpEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DNatural, note.FSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantNinthFlatThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DNatural, note.AFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DNatural, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantThirteenthFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DFlat, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhAddFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.FNatural, note.GNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhAddEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhAddThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhAddSharpFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.FSharp, note.GNatural, note.BFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhFlatFifthFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GFlat, note.BFlat, note.DFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSharpFifthFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GSharp, note.BFlat, note.DFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhFlatNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhFlatNinthFlatThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DFlat, note.AFlat},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSharpNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSharpNinthSharpEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.DSharp, note.FSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DominantSeventhSharpEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BFlat, note.FSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DiminishedMajorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GFlat, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.DiminishedMajorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.EFlat, note.GFlat, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhFlatFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GFlat, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhSuspendedSecond),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorNinthSuspendedSecond),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.GNatural, note.BNatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhSuspendedFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorNinthSuspendedFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GNatural, note.BNatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhSuspendedFourthSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.GSharp, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhSuspendedFourthDoubleSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.FNatural, note.ANatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorNinth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural, note.DNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural, note.DNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural, note.DNatural, note.FNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhAddEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural, note.FNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhAddThirteenth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural, note.ANatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhAddSharpEleventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GNatural, note.BNatural, note.FSharp},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhAddFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.FNatural, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhAddSharpFourth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.FSharp, note.GNatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.MajorSeventhDoubleSharpFifth),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.ANatural, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.AugmentedMajorSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GSharp, note.BNatural},
		},
		{
			GivenChord:    chord.New(note.CNatural, chordtype.AugmentedAugmentedSeventh),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GSharp, note.BSharp},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TestChord%sNotes", tc.GivenChord), func(t *testing.T) {
			computedNotes := tc.GivenChord.Notes()
			errMsg := fmt.Sprintf("failed on %s, expected = %#v, actual = %#v", tc.GivenChord.String(), tc.ExpectedNotes, computedNotes)
			require.Equal(t, tc.ExpectedNotes, computedNotes, errMsg)
		})
	}
}

func TestChord_Number(t *testing.T) {
	type testCase struct {
		GivenChord     chord.Chord
		ExpectedNumber int
	}

	testCases := []testCase{
		{
			GivenChord:     chord.New(note.CNatural, chordtype.Major),
			ExpectedNumber: 1 + 16 + 128,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TestChord%sNumber", tc.GivenChord), func(t *testing.T) {
			require.Equal(t, tc.ExpectedNumber, tc.GivenChord.Number())
		})
	}
}
