package scale_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/scale"
	"github.com/stretchr/testify/require"
)

func TestScale_Perfection(t *testing.T) {
	perfect, imperfect, prefectionProfile := scale.Lydian.Perfection()
	require.Equal(t, 6, perfect)
	require.Equal(t, 1, imperfect)
	t.Logf("perfection profile = %v", prefectionProfile)
}
