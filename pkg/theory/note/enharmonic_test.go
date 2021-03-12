package note_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNote_Enharmonic(t *testing.T) {
	for octave := 0; octave < 11; octave++ {
		for _, n1 := range note.AllNotes() {
			enharmonics := n1.Enharmonic()
			for _, enharmonic := range enharmonics {
				n2 := enharmonic.Note
				octaveShift := enharmonic.OctaveShift
				require.Equal(t, n1.MIDINoteNumber(octave), n2.MIDINoteNumber(octave+octaveShift), fmt.Sprintf("failed at note %s with enharmonic %s (%d)", n1, n2, octaveShift))
			}
		}
	}
}
