package interval

var mapIntervalToSemitone = map[Interval]int{
	PerfectUnison:        0,
	DiminishedSecond:     0,
	MinorSecond:          1,
	AugmentedUnison:      1,
	MajorSecond:          2,
	DiminishedThird:      2,
	MinorThird:           3,
	AugmentedSecond:      3,
	MajorThird:           4,
	DiminishedFourth:     4,
	PerfectFourth:        5,
	AugmentedThird:       5,
	DiminishedFifth:      6,
	AugmentedFourth:      6,
	PerfectFifth:         7,
	DiminishedSixth:      7,
	MinorSixth:           8,
	AugmentedFifth:       8,
	MajorSixth:           9,
	DiminishedSeventh:    9,
	MinorSeventh:         10,
	AugmentedSixth:       10,
	MajorSeventh:         11,
	DiminishedOctave:     11,
	PerfectOctave:        12,
	AugmentedSeventh:     12,
	DiminishedNinth:      12,
	MinorNinth:           13,
	MajorNinth:           14,
	AugmentedNinth:       15,
	DiminishedTenth:      14,
	MinorTenth:           15,
	MajorTenth:           16,
	AugmentedTenth:       17,
	DiminishedEleventh:   16,
	PerfectEleventh:      17,
	AugmentedEleventh:    18,
	DiminishedTwelfth:    18,
	PerfectTwelfth:       19,
	AugmentedTwelfth:     20,
	DiminishedThirteenth: 19,
	MinorThirteenth:      20,
	MajorThirteenth:      21,
	AugmentedThirteenth:  22,
	DiminishedFourteenth: 21,
	MinorFourteenth:      22,
	MajorFourteenth:      23,
	AugmentedFourteenth:  24,
	DiminihsedFifteenth:  23,
	PerfectFifteenth:     24,
	AugmentedFifteenth:   25,
}

// Semitones converts interval into semitones
func (i Interval) Semitones() int {
	return mapIntervalToSemitone[i]
}

// FromSemitones returns possible interval from given semitones count
func FromSemitones(semintones int) ([]Interval, error) {
	if semintones < 0 {
		return nil, ErrInvalidSemitonesCount
	}

	if semintones, found := mapSemitoneToIntervals[semintones]; found {
		return semintones, nil
	}

	return nil, ErrNoSuchInterval
}
