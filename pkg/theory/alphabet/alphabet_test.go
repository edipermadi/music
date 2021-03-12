package alphabet_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/alphabet"
	"github.com/stretchr/testify/require"
)

func TestAlphabet_String(t *testing.T) {
	for _, a := range alphabet.Alphabets() {
		require.NotEmpty(t, a.String())
	}
}

func TestAlphabet_Alphabets(t *testing.T) {
	alphabets := alphabet.Alphabets()
	require.NotEmpty(t, alphabets)
}
