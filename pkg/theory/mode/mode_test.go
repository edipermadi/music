package mode_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/mode"
	"github.com/edipermadi/music/pkg/theory/mode/modetype"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestMode_AllModes(t *testing.T) {
	require.NotEmpty(t, mode.AllModes())
}

func TestMode_Notes(t *testing.T) {
	type testCase struct {
		GivenMode     mode.Mode
		ExpectedNotes note.Notes
	}

	testCases := []testCase{
		{
			GivenMode:     mode.New(note.CNatural, modetype.Minoric),
			ExpectedNotes: note.Notes{note.CNatural, note.ENatural, note.GSharp, note.CNatural},
		},
		{
			GivenMode:     mode.New(note.CNatural, modetype.Aerathic),
			ExpectedNotes: note.Notes{note.CNatural, note.CSharp, note.FNatural, note.GSharp, note.CNatural},
		},
		{
			GivenMode:     mode.New(note.CNatural, modetype.Aeolian),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.EFlat, note.FNatural, note.GNatural, note.AFlat, note.BFlat, note.CNatural},
		},
		{
			GivenMode:     mode.New(note.DSharp, modetype.Aeolian),
			ExpectedNotes: note.Notes{note.DSharp, note.ESharp, note.FSharp, note.GSharp, note.ASharp, note.BNatural, note.CSharp, note.DSharp},
		},
		{
			GivenMode:     mode.New(note.CNatural, modetype.Ionythian),
			ExpectedNotes: note.Notes{note.CNatural, note.DDoubleSharp, note.ESharp, note.FDoubleSharp, note.GDoubleSharp, note.ASharp, note.BNatural, note.CNatural},
		},
		{
			GivenMode:     mode.New(note.CNatural, modetype.Rocrimic),
			ExpectedNotes: note.Notes{note.CNatural, note.DNatural, note.EFlat, note.FNatural, note.GNatural, note.ASharp, note.CNatural},
		},
		{
			GivenMode:     mode.New(note.CNatural, modetype.Eporimic),
			ExpectedNotes: note.Notes{note.CNatural, note.DFlat, note.EFlat, note.FNatural, note.GSharp, note.ASharp, note.CNatural},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Mode%sNotes", tc.GivenMode), func(t *testing.T) {
			require.Equal(t, tc.ExpectedNotes, tc.GivenMode.Notes())
		})
	}
}

func TestMode_Number(t *testing.T) {
	type testCase struct {
		GivenMode      mode.Mode
		ExpectedNumber int
	}

	testCases := []testCase{
		{
			// https://ianring.com/musictheory/scales/2741
			GivenMode:      mode.New(note.CNatural, modetype.Ionian),
			ExpectedNumber: 2741,
		},
		{
			// https://ianring.com/musictheory/scales/1709
			GivenMode:      mode.New(note.CNatural, modetype.Dorian),
			ExpectedNumber: 1709,
		},
		{
			// https://ianring.com/musictheory/scales/1451
			GivenMode:      mode.New(note.CNatural, modetype.Phrygian),
			ExpectedNumber: 1451,
		},
		{
			// https://ianring.com/musictheory/scales/2773
			GivenMode:      mode.New(note.CNatural, modetype.Lydian),
			ExpectedNumber: 2773,
		},
		{
			// https://ianring.com/musictheory/scales/1717
			GivenMode:      mode.New(note.CNatural, modetype.Mixolydian),
			ExpectedNumber: 1717,
		},
		{
			// https://ianring.com/musictheory/scales/1453
			GivenMode:      mode.New(note.CNatural, modetype.Aeolian),
			ExpectedNumber: 1453,
		},
		{
			// https://ianring.com/musictheory/scales/1387
			GivenMode:      mode.New(note.CNatural, modetype.Locrian),
			ExpectedNumber: 1387,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Mode%sNumber", tc.GivenMode), func(t *testing.T) {
			require.Equal(t, tc.ExpectedNumber, tc.GivenMode.Number())
		})

	}
}

func TestMode_Luminosity(t *testing.T) {
	// flat
	{
		require.Equal(t, 7, mode.DFlatLydian.Luminosity())
		require.Equal(t, 6, mode.DFlatIonian.Luminosity())
		require.Equal(t, 5, mode.DFlatMixolydian.Luminosity())
		require.Equal(t, 4, mode.DFlatDorian.Luminosity())
		require.Equal(t, 3, mode.DFlatAeolian.Luminosity())
		require.Equal(t, 2, mode.DFlatPhrygian.Luminosity())
		require.Equal(t, 1, mode.DFlatLocrian.Luminosity())
	}
	
	// natural
	{
		require.Equal(t, 7, mode.CNaturalLydian.Luminosity())
		require.Equal(t, 6, mode.CNaturalIonian.Luminosity())
		require.Equal(t, 5, mode.CNaturalMixolydian.Luminosity())
		require.Equal(t, 4, mode.CNaturalDorian.Luminosity())
		require.Equal(t, 3, mode.CNaturalAeolian.Luminosity())
		require.Equal(t, 2, mode.CNaturalPhrygian.Luminosity())
		require.Equal(t, 1, mode.CNaturalLocrian.Luminosity())
	}

	// sharp
	{
		require.Equal(t, 7, mode.CSharpLydian.Luminosity())
		require.Equal(t, 6, mode.CSharpIonian.Luminosity())
		require.Equal(t, 5, mode.CSharpMixolydian.Luminosity())
		require.Equal(t, 4, mode.CSharpDorian.Luminosity())
		require.Equal(t, 3, mode.CSharpAeolian.Luminosity())
		require.Equal(t, 2, mode.CSharpPhrygian.Luminosity())
		require.Equal(t, 1, mode.CSharpLocrian.Luminosity())
	}
}
