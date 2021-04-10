package illustration_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/edipermadi/music/pkg/illustration"
	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestKeyboard_Draw(t *testing.T) {
	testOutDir := os.Getenv("TEST_OUT_DIR")
	if testOutDir == "" {
		t.Skip("TEST_OUT_DIR is empty")
	}

	template := fmt.Sprintf("test_%s", time.Now().Format("20060102_150405"))
	for _, givenPitch := range pitch.SimplePianoPitches(37) {
		t.Run(givenPitch.String(), func(t *testing.T) {
			filename := fmt.Sprintf("%s/%s_%s.png", testOutDir, template, givenPitch.String())
			file, err := os.Create(filename)
			require.NoError(t, err)

			defer func() { _ = file.Close() }()

			kbd := illustration.NewKeyboard(file)
			require.NoError(t, kbd.Draw(givenPitch))
		})
	}
}
