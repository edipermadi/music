package pitch_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestPitch_String(t *testing.T) {
	for _, p := range pitch.AllPitches() {
		require.NotEmpty(t, p.String())
	}
}

func TestPitch_PianoPitches(t *testing.T) {
	type testCase struct {
		Keys        int
		ExpectedMin int
		ExpectedMax int
	}

	testCases := []testCase{
		{
			Keys:        25,
			ExpectedMin: pitch.C4Natural.MIDINoteNumber(),
			ExpectedMax: pitch.C6Natural.MIDINoteNumber(),
		},
		{
			Keys:        88,
			ExpectedMin: pitch.A0Natural.MIDINoteNumber(),
			ExpectedMax: pitch.C8Natural.MIDINoteNumber(),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%dKeys", tc.Keys), func(t *testing.T) {
			allPitches := pitch.PianoPitches(tc.Keys)
			for _, givenPitch := range allPitches {
				require.True(t, givenPitch.MIDINoteNumber() >= tc.ExpectedMin)
				require.True(t, givenPitch.MIDINoteNumber() <= tc.ExpectedMax)
			}
		})
	}
}
