package note_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/stretchr/testify/require"
)

func TestAlphabet_String(t *testing.T) {
	for _, n := range note.AllNotes() {
		require.Equal(t, string(n.String()[0]), n.Alphabet().String())
	}
}
