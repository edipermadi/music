package note_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestNote_String(t *testing.T) {
	for _, n := range note.AllNotes() {
		require.NotEmpty(t, n.String())
	}
}

func TestNote_Notes(t *testing.T) {
	require.NotEmpty(t, note.AllNotes())
}

func TestNote_Name(t *testing.T) {
	for _, givenNote := range note.AllNotes() {
		require.NotEmpty(t, givenNote.Name(), fmt.Sprintf("failed on %s", givenNote))
	}
}
