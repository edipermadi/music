package soundfont_test

import (
	"errors"
	"os"
	"testing"

	"github.com/edipermadi/music/pkg/soundfont"
	"github.com/stretchr/testify/require"
)

func TestSoundFont_Unmarshall(t *testing.T) {
	soundFontFilePath := "../../resources/FluidR3_GM.sf2"

	if _, err := os.Stat(soundFontFilePath); errors.Is(err, os.ErrNotExist) {
		t.Skip()
		return
	}

	file, err := os.Open(soundFontFilePath)
	require.NoError(t, err)

	defer func() { _ = file.Close() }()

	var sf soundfont.SoundFont
	require.NoError(t, sf.Unmarshall(file))
	t.Logf("soundfile = %#v", sf)
}
