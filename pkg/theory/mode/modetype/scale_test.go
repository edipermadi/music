package modetype_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/mode/modetype"
	"github.com/edipermadi/music/pkg/theory/scale"
	"github.com/stretchr/testify/require"
)

func TestMode_Scale(t *testing.T) {
	for _, givenType := range modetype.AllTypes() {
		s, n := givenType.Scale()
		require.NotEmpty(t, n)
		require.NotEqual(t, scale.Invalid, s)
		t.Logf("%s is a mode of %s scale", givenType, s)
	}
}
