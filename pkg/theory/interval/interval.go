package interval

// Interval represents interval between notes
type Interval int

// Interval enumeration
const (
	Invalid Interval = iota

	// unison
	PerfectUnison   Interval = iota // 0 semitones
	AugmentedUnison Interval = iota // 1 semitones

	// second
	DiminishedSecond Interval = iota // 0 semitones
	MinorSecond      Interval = iota // 1 semitones
	MajorSecond      Interval = iota // 2 semitones
	AugmentedSecond  Interval = iota // 3 semitones

	// third
	DiminishedThird Interval = iota // 2 semitones
	MinorThird      Interval = iota // 3 semitones
	MajorThird      Interval = iota // 4 semitones
	AugmentedThird  Interval = iota // 5 semitones

	// fourth
	DiminishedFourth Interval = iota // 4 semitones
	PerfectFourth    Interval = iota // 5 semitones
	AugmentedFourth  Interval = iota // 6 semitones

	// fifth
	DiminishedFifth Interval = iota // 6 semitones
	PerfectFifth    Interval = iota // 7 semitones
	AugmentedFifth  Interval = iota // 8 semitones

	// sixth
	DiminishedSixth Interval = iota // 7 semitones
	MinorSixth      Interval = iota // 8 semitones
	MajorSixth      Interval = iota // 9 semitones
	AugmentedSixth  Interval = iota // 10 semitones

	// seventh
	DiminishedSeventh Interval = iota // 9 semitones
	MinorSeventh      Interval = iota // 10 semitones
	MajorSeventh      Interval = iota // 11 semitones
	AugmentedSeventh  Interval = iota // 12 semitones

	// octave
	DiminishedOctave Interval = iota // 11 semitones
	PerfectOctave    Interval = iota // 12 semitones
	AugmentedOctave  Interval = iota // 13 semitones

	// ninth
	DiminishedNinth Interval = iota // 12 semitones
	MinorNinth      Interval = iota // 13 semitones
	MajorNinth      Interval = iota // 14 semitones
	AugmentedNinth  Interval = iota // 15 semitones

	// tenth
	DiminishedTenth Interval = iota // 14 semitones
	MinorTenth      Interval = iota // 15 semitones
	MajorTenth      Interval = iota // 16 semitones
	AugmentedTenth  Interval = iota // 17 semitones

	// eleventh
	DiminishedEleventh Interval = iota // 16 semitones
	PerfectEleventh    Interval = iota // 17 semitones
	AugmentedEleventh  Interval = iota // 18 semitones

	// twelfth
	DiminishedTwelfth Interval = iota // 18 semitones
	PerfectTwelfth    Interval = iota // 19 semitones
	AugmentedTwelfth  Interval = iota // 20 semitones

	// thirteenth
	DiminishedThirteenth Interval = iota // 19 semitones
	MinorThirteenth      Interval = iota // 20 semitones
	MajorThirteenth      Interval = iota // 21 semitones
	AugmentedThirteenth  Interval = iota // 22 semitones

	// fourteenth
	DiminishedFourteenth Interval = iota // 21 semitones
	MinorFourteenth      Interval = iota // 22 semitones
	MajorFourteenth      Interval = iota // 23 semitones
	AugmentedFourteenth  Interval = iota // 24 semitones

	// fifteenth
	DiminihsedFifteenth Interval = iota // 23 semitones
	PerfectFifteenth    Interval = iota // 24 semitones
	AugmentedFifteenth  Interval = iota // 25 semitones
)

var intervalStringifieds = map[Interval]string{
	Invalid:              "Invalid",
	PerfectUnison:        "PerfectUnison",
	AugmentedUnison:      "AugmentedUnison",
	DiminishedSecond:     "DiminishedSecond",
	MinorSecond:          "MinorSecond",
	MajorSecond:          "MajorSecond",
	AugmentedSecond:      "AugmentedSecond",
	DiminishedThird:      "DiminishedThird",
	MinorThird:           "MinorThird",
	MajorThird:           "MajorThird",
	AugmentedThird:       "AugmentedThird",
	DiminishedFourth:     "DiminishedFourth",
	PerfectFourth:        "PerfectFourth",
	AugmentedFourth:      "AugmentedFourth",
	DiminishedFifth:      "DiminishedFifth",
	PerfectFifth:         "PerfectFifth",
	AugmentedFifth:       "AugmentedFifth",
	DiminishedSixth:      "DiminishedSixth",
	MinorSixth:           "MinorSixth",
	MajorSixth:           "MajorSixth",
	AugmentedSixth:       "AugmentedSixth",
	DiminishedSeventh:    "DiminishedSeventh",
	MinorSeventh:         "MinorSeventh",
	MajorSeventh:         "MajorSeventh",
	AugmentedSeventh:     "AugmentedSeventh",
	DiminishedOctave:     "DiminishedOctave",
	PerfectOctave:        "PerfectOctave",
	AugmentedOctave:      "AugmentedOctave",
	DiminishedNinth:      "DiminishedNinth",
	MinorNinth:           "MinorNinth",
	MajorNinth:           "MajorNinth",
	AugmentedNinth:       "AugmentedNinth",
	DiminishedTenth:      "DiminishedTenth",
	MinorTenth:           "MinorTenth",
	MajorTenth:           "MajorTenth",
	AugmentedTenth:       "AugmentedTenth",
	DiminishedEleventh:   "DiminishedEleventh",
	PerfectEleventh:      "PerfectEleventh",
	AugmentedEleventh:    "AugmentedEleventh",
	DiminishedTwelfth:    "DiminishedTwelfth",
	PerfectTwelfth:       "PerfectTwelfth",
	AugmentedTwelfth:     "AugmentedTwelfth",
	DiminishedThirteenth: "DiminishedThirteenth",
	MinorThirteenth:      "MinorThirteenth",
	MajorThirteenth:      "MajorThirteenth",
	AugmentedThirteenth:  "AugmentedThirteenth",
	DiminishedFourteenth: "DiminishedFourteenth",
	MinorFourteenth:      "MinorFourteenth",
	MajorFourteenth:      "MajorFourteenth",
	AugmentedFourteenth:  "AugmentedFourteenth",
	DiminihsedFifteenth:  "DiminihsedFifteenth",
	PerfectFifteenth:     "PerfectFifteenth",
	AugmentedFifteenth:   "AugmentedFifteenth",
}

// String returns stringified name
func (i Interval) String() string {
	return intervalStringifieds[i]
}

// GoString satisfies GoStringer
func (i Interval) GoString() string {
	return i.String()
}

// AllIntervals retuns slice of intervals
func AllIntervals() []Interval {
	return allIntervals
}
