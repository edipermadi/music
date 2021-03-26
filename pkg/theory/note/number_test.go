package note_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
)

func TestNote_Number(t *testing.T) {
	type testCase struct {
		Note           note.Note
		ExpectedNumber int
	}

	testCases := []testCase{
		{Note: note.CTripleFlat, ExpectedNumber: 9},
		{Note: note.CDoubleFlat, ExpectedNumber: 10},
		{Note: note.CFlat, ExpectedNumber: 11},
		{Note: note.CNatural, ExpectedNumber: 0},
		{Note: note.CSharp, ExpectedNumber: 1},
		{Note: note.CDoubleSharp, ExpectedNumber: 2},
		{Note: note.CTripleSharp, ExpectedNumber: 3},

		{Note: note.DTripleFlat, ExpectedNumber: 11},
		{Note: note.DDoubleFlat, ExpectedNumber: 0},
		{Note: note.DFlat, ExpectedNumber: 1},
		{Note: note.DNatural, ExpectedNumber: 2},
		{Note: note.DSharp, ExpectedNumber: 3},
		{Note: note.DDoubleSharp, ExpectedNumber: 4},
		{Note: note.DTripleSharp, ExpectedNumber: 5},

		{Note: note.ETripleFlat, ExpectedNumber: 1},
		{Note: note.EDoubleFlat, ExpectedNumber: 2},
		{Note: note.EFlat, ExpectedNumber: 3},
		{Note: note.ENatural, ExpectedNumber: 4},
		{Note: note.ESharp, ExpectedNumber: 5},
		{Note: note.EDoubleSharp, ExpectedNumber: 6},
		{Note: note.ETripleSharp, ExpectedNumber: 7},

		{Note: note.FTripleFlat, ExpectedNumber: 2},
		{Note: note.FDoubleFlat, ExpectedNumber: 3},
		{Note: note.FFlat, ExpectedNumber: 4},
		{Note: note.FNatural, ExpectedNumber: 5},
		{Note: note.FSharp, ExpectedNumber: 6},
		{Note: note.FDoubleSharp, ExpectedNumber: 7},
		{Note: note.FTripleSharp, ExpectedNumber: 8},

		{Note: note.GTripleFlat, ExpectedNumber: 4},
		{Note: note.GDoubleFlat, ExpectedNumber: 5},
		{Note: note.GFlat, ExpectedNumber: 6},
		{Note: note.GNatural, ExpectedNumber: 7},
		{Note: note.GSharp, ExpectedNumber: 8},
		{Note: note.GDoubleSharp, ExpectedNumber: 9},
		{Note: note.GTripleSharp, ExpectedNumber: 10},

		{Note: note.ATripleFlat, ExpectedNumber: 6},
		{Note: note.ADoubleFlat, ExpectedNumber: 7},
		{Note: note.AFlat, ExpectedNumber: 8},
		{Note: note.ANatural, ExpectedNumber: 9},
		{Note: note.ASharp, ExpectedNumber: 10},
		{Note: note.ADoubleSharp, ExpectedNumber: 11},
		{Note: note.ATripleSharp, ExpectedNumber: 0},

		{Note: note.BTripleFlat, ExpectedNumber: 8},
		{Note: note.BDoubleFlat, ExpectedNumber: 9},
		{Note: note.BFlat, ExpectedNumber: 10},
		{Note: note.BNatural, ExpectedNumber: 11},
		{Note: note.BSharp, ExpectedNumber: 0},
		{Note: note.BDoubleSharp, ExpectedNumber: 1},
		{Note: note.BTripleSharp, ExpectedNumber: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.Note.String(), func(t *testing.T) {
			require.Equal(t, tc.ExpectedNumber, tc.Note.Number())
		})
	}
}
