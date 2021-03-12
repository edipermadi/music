package illustration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/edipermadi/music/pkg/illustration"
	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestMarshallToMIDI(t *testing.T) {
	testOutDir := os.Getenv("TEST_OUT_DIR")
	require.NotEmpty(t, testOutDir)

	file, err := os.Create(fmt.Sprintf("%s/output.mid", testOutDir))
	require.NoError(t, err)
	defer func() { _ = file.Close() }()

	midiFile, err := illustration.SaveAsMIDI("CNaturalMajor", pitch.C4Natural, pitch.E4Natural, pitch.G4Natural)
	require.NoError(t, err)
	require.NoError(t, midiFile.Marshall(file))
}
