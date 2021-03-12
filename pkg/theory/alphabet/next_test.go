package alphabet_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/alphabet"
	"github.com/stretchr/testify/require"
)

func TestAlphabet_Next(t *testing.T) {
	for _, a := range alphabet.Alphabets() {
		require.NotEqual(t, a, a.Next())
	}
}
