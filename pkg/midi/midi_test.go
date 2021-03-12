package midi_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/edipermadi/music/pkg/midi"
	"github.com/stretchr/testify/require"
)

func TestFile_Load(t *testing.T) {
	midiFilePath := "../../resources/twinkle-twinkle-little-star.mid"

	if _, err := os.Stat(midiFilePath); errors.Is(err, os.ErrNotExist) {
		t.Skip()
		return
	}

	file, err := os.Open(midiFilePath)
	require.NoError(t, err)

	source, err := ioutil.ReadAll(file)
	require.NoError(t, file.Close())

	var midiFile midi.File
	require.NoError(t, midiFile.Unmarshall(bytes.NewReader(source)))
	t.Logf("midi-file = %#v", midiFile)

	var buff bytes.Buffer
	require.NoError(t, midiFile.Marshall(&buff))

	destination := buff.Bytes()
	require.Equal(t, source, destination)

}
