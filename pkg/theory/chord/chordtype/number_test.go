package chordtype_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/stretchr/testify/require"
)

func TestType_Number(t *testing.T) {
	chordTypeNumbers := map[chordtype.Type]int{
		// Triad
		chordtype.Diminished:      73,  // (1 << 0) + (1 << 3) + (1 << 6)
		chordtype.Minor:           137, // (1 << 0) + (1 << 3) + (1 << 7)
		chordtype.Major:           145, // (1 << 0) + (1 << 4) + (1 << 7)
		chordtype.Augmented:       273, // (1 << 0) + (1 << 4) + (1 << 8)
		chordtype.SuspendedSecond: 133, // (1 << 0) + (1 << 2) + (1 << 7)
		chordtype.SuspendedFourth: 161, // (1 << 0) + (1 << 5) + (1 << 7)
		chordtype.MajorFlatFifth:  81,  // (1 << 0) + (1 << 4) + (1 << 6)

		// Seventh
		chordtype.FullDiminishedSeventh:          585,  // (1 << 0) + (1 << 3) + (1 << 6) + (1 << 9)
		chordtype.HalfDiminishedSeventh:          1097, // (1 << 0) + (1 << 3) + (1 << 6) + (1 << 10)
		chordtype.MinorSeventh:                   1161, // (1 << 0) + (1 << 3) + (1 << 7) + (1 << 10)
		chordtype.MinorMajorSeventh:              2185, // (1 << 0) + (1 << 3) + (1 << 7) + (1 << 11)
		chordtype.DominantSeventh:                1169, // (1 << 0) + (1 << 4) + (1 << 7) + (1 << 10)
		chordtype.MajorSeventh:                   2193, // (1 << 0) + (1 << 4) + (1 << 7) + (1 << 11)
		chordtype.AugmentedMajorSeventh:          2321, // (1 << 0) + (1 << 4) + (1 << 8) + (1 << 11)
		chordtype.AugmentedAugmentedSeventh:      4369, // (1 << 0) + (1 << 4) + (1 << 8) + (1 << 12)
		chordtype.DominantSeventhSuspendedSecond: 1157, // (1 << 0) + (1 << 2) + (1 << 7) + (1 << 10)
		chordtype.DominantSeventhSuspendedFourth: 1185, // (1 << 0) + (1 << 5) + (1 << 7) + (1 << 10)
	}

	for givenType, expectedNumber := range chordTypeNumbers {
		require.Equal(t, expectedNumber, givenType.Number())
	}
}
