package note_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/limit"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNote_MIDINoteNumber(t *testing.T) {
	expectedMIDINoteNumbers := map[note.Note][]int{
		note.CDoubleFlat:  {-1, 10, 22, 34, 46, 58, 70, 82, 94, 106, 118, -1},
		note.CFlat:        {-1, 11, 23, 35, 47, 59, 71, 83, 95, 107, 119, -1},
		note.CNatural:     {0, 12, 24, 36, 48, 60, 72, 84, 96, 108, 120, -1},
		note.CSharp:       {1, 13, 25, 37, 49, 61, 73, 85, 97, 109, 121, -1},
		note.CDoubleSharp: {2, 14, 26, 38, 50, 62, 74, 86, 98, 110, 122, -1},

		note.DDoubleFlat:  {0, 12, 24, 36, 48, 60, 72, 84, 96, 108, 120, -1},
		note.DFlat:        {1, 13, 25, 37, 49, 61, 73, 85, 97, 109, 121, -1},
		note.DNatural:     {2, 14, 26, 38, 50, 62, 74, 86, 98, 110, 122, -1},
		note.DSharp:       {3, 15, 27, 39, 51, 63, 75, 87, 99, 111, 123, -1},
		note.DDoubleSharp: {4, 16, 28, 40, 52, 64, 76, 88, 100, 112, 124, -1},

		note.EDoubleFlat:  {2, 14, 26, 38, 50, 62, 74, 86, 98, 110, 122, -1},
		note.EFlat:        {3, 15, 27, 39, 51, 63, 75, 87, 99, 111, 123, -1},
		note.ENatural:     {4, 16, 28, 40, 52, 64, 76, 88, 100, 112, 124, -1},
		note.ESharp:       {5, 17, 29, 41, 53, 65, 77, 89, 101, 113, 125, -1},
		note.EDoubleSharp: {6, 18, 30, 42, 54, 66, 78, 90, 102, 114, 126, -1},

		note.FDoubleFlat:  {3, 15, 27, 39, 51, 63, 75, 87, 99, 111, 123, -1},
		note.FFlat:        {4, 16, 28, 40, 52, 64, 76, 88, 100, 112, 124, -1},
		note.FNatural:     {5, 17, 29, 41, 53, 65, 77, 89, 101, 113, 125, -1},
		note.FSharp:       {6, 18, 30, 42, 54, 66, 78, 90, 102, 114, 126, -1},
		note.FDoubleSharp: {7, 19, 31, 43, 55, 67, 79, 91, 103, 115, 127, -1},

		note.GDoubleFlat:  {5, 17, 29, 41, 53, 65, 77, 89, 101, 113, 125, -1},
		note.GFlat:        {6, 18, 30, 42, 54, 66, 78, 90, 102, 114, 126, -1},
		note.GNatural:     {7, 19, 31, 43, 55, 67, 79, 91, 103, 115, 127, -1},
		note.GSharp:       {8, 20, 32, 44, 56, 68, 80, 92, 104, 116, -1, -1},
		note.GDoubleSharp: {9, 21, 33, 45, 57, 69, 81, 93, 105, 117, -1, -1},

		note.ADoubleFlat:  {7, 19, 31, 43, 55, 67, 79, 91, 103, 115, 127, -1},
		note.AFlat:        {8, 20, 32, 44, 56, 68, 80, 92, 104, 116, -1, -1},
		note.ANatural:     {9, 21, 33, 45, 57, 69, 81, 93, 105, 117, -1, -1},
		note.ASharp:       {10, 22, 34, 46, 58, 70, 82, 94, 106, 118, -1, -1},
		note.ADoubleSharp: {11, 23, 35, 47, 59, 71, 83, 95, 107, 119, -1, -1},

		note.BDoubleFlat:  {9, 21, 33, 45, 57, 69, 81, 93, 105, 117, -1, -1},
		note.BFlat:        {10, 22, 34, 46, 58, 70, 82, 94, 106, 118, -1, -1},
		note.BNatural:     {11, 23, 35, 47, 59, 71, 83, 95, 107, 119, -1, -1},
		note.BSharp:       {12, 24, 36, 48, 60, 72, 84, 96, 108, 120, -1, -1},
		note.BDoubleSharp: {13, 25, 37, 49, 61, 73, 85, 97, 109, 121, -1, -1},
	}

	for n, values := range expectedMIDINoteNumbers {
		for octave := -1; octave < 11; octave++ {
			require.Equal(t, values[octave+1], n.MIDINoteNumber(octave), fmt.Errorf("failed at note %s octave %d", n, octave))
		}
	}

}

func TestNote_FromMIDINoteNumber(t *testing.T) {
	for i := limit.MIDINoteNumberMin; i <= limit.MIDINoteNumberMax; i++ {
		notes, err := note.FromMIDINoteNumber(i)
		require.NoError(t, err)
		t.Logf("midi note %d => %#v", i, notes)
		switch i % limit.SemitonesInOneOctave {
		case 0:
			require.Contains(t, notes, note.CNatural)
		case 1:
			require.Contains(t, notes, note.CSharp)
		case 2:
			require.Contains(t, notes, note.DNatural)
		case 3:
			require.Contains(t, notes, note.DSharp)
		case 4:
			require.Contains(t, notes, note.ENatural)
		case 5:
			require.Contains(t, notes, note.FNatural)
		case 6:
			require.Contains(t, notes, note.FSharp)
		case 7:
			require.Contains(t, notes, note.GNatural)
		case 8:
			require.Contains(t, notes, note.GSharp)
		case 9:
			require.Contains(t, notes, note.ANatural)
		case 10:
			require.Contains(t, notes, note.ASharp)
		case 11:
			require.Contains(t, notes, note.BNatural)
		}
	}
}
