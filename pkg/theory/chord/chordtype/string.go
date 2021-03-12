package chordtype

var mapTypeToString = map[Type]string{
	Invalid: "Invalid",

	// Dyad
	PowerChord: "PowerChord",

	// Triad
	Diminished:                            "Diminished",
	MinorDoubleFlatFifth:                  "MinorDoubleFlatFifth",
	Minor:                                 "Minor",
	MinorAddEleventh:                      "MinorAddEleventh",
	MinorAddNinth:                         "MinorAddNinth",
	MinorAddFourth:                        "MinorAddFourth",
	MinorAddSharpFourth:                   "MinorAddSharpFourth",
	MinorAddSharpNinth:                    "MinorAddSharpNinth",
	MinorSharpFifth:                       "MinorSharpFifth",
	MajorFlatFifth:                        "MajorFlatFifth",
	Major:                                 "Major",
	MajorAddEleventh:                      "MajorAddEleventh",
	MajorAddNinth:                         "MajorAddNinth",
	MajorAddFourth:                        "MajorAddFourth",
	MajorAddSharpFourth:                   "MajorAddSharpFourth",
	MajorAddSharpNinth:                    "MajorAddSharpNinth",
	MajorDoubleSharpFifth:                 "MajorDoubleSharpFifth",
	Augmented:                             "Augmented",
	SuspendedSecondDoubleFlatFifth:        "SuspendedSecondDoubleFlatFifth",
	SuspendedSecondFlatFifth:              "SuspendedSecondFlatFifth",
	SuspendedSecondFlatFifthAddSharpFifth: "SuspendedSecondFlatFifthAddSharpFifth",
	SuspendedSecond:                       "SuspendedSecond",
	SuspendedSecondSharpFifth:             "SuspendedSecondSharpFifth",
	SuspendedFourthFlatFifth:              "SuspendedFourthFlatFifth",
	SuspendedFourth:                       "SuspendedFourth",
	SuspendedFourthSharpFifth:             "SuspendedFourthSharpFifth",
	SuspendedFourthDoubleSharpFifth:       "SuspendedFourthDoubleSharpFifth",
	Phrygian:                              "Phrygian",
	PhrygianAddSeventh:                    "PhrygianAddSeventh",
	Lydian:                                "Lydian",
	LydianMajorSeventh:                    "LydianMajorSeventh",
	Locrian:                               "Locrian",
	Quartal:                               "Quartal",
	QuartalAugmented:                      "QuartalAugmented",

	// Sixth
	MinorSixth:                               "MinorSixth",
	MinorSixthAddNinth:                       "MinorSixthAddNinth",
	MinorSixthAddFlatNinth:                   "MinorSixthAddFlatNinth",
	MajorSixth:                               "MajorSixth",
	MajorSixthAddFlatNinth:                   "MajorSixthAddFlatNinth",
	MajorSixthFlatFifth:                      "MajorSixthFlatFifth",
	MajorSixthAddNinth:                       "MajorSixthAddNinth",
	MajorSixthSuspendedSecond:                "MajorSixthSuspendedSecond",
	MajorSixthSuspendedSecondFlatFifth:       "MajorSixthSuspendedSecondFlatFifth",
	MajorSixthSuspendedSecondDoubleFlatFifth: "MajorSixthSuspendedSecondDoubleFlatFifth",
	MajorSixthSuspendedFourth:                "MajorSixthSuspendedFourth",

	// Seventh
	FullDiminishedSeventh:                       "FullDiminishedSeventh",
	HalfDiminishedSeventh:                       "HalfDiminishedSeventh",
	MinorSeventhDoubleFlatFifth:                 "MinorSeventhDoubleFlatFifth",
	MinorSeventh:                                "MinorSeventh",
	MinorSeventhAddEleventh:                     "MinorSeventhAddEleventh",
	MinorSeventhAddThirteenth:                   "MinorSeventhAddThirteenth",
	MinorSeventhAddSharpEleventh:                "MinorSeventhAddSharpEleventh",
	MinorSeventhSharpFifth:                      "MinorSeventhSharpFifth",
	MinorSeventhFlatNinth:                       "MinorSeventhFlatNinth",
	MinorMajorSeventh:                           "MinorMajorSeventh",
	MinorMajorSeventhAddEleventh:                "MinorMajorSeventhAddEleventh",
	MinorMajorSeventhAddThirteenth:              "MinorMajorSeventhAddThirteenth",
	DominantSeventhSuspendedSecondFlatFifth:     "DominantSeventhSuspendedSecondFlatFifth",
	DominantSeventhFlatFifth:                    "DominantSeventhFlatFifth",
	DominantSeventhSuspendedSecond:              "DominantSeventhSuspendedSecond",
	DominantSeventhSuspendedFourth:              "DominantSeventhSuspendedFourth",
	DominantSeventh:                             "DominantSeventh",
	DominantSeventhAddFourth:                    "DominantSeventhAddFourth",
	DominantSeventhAddEleventh:                  "DominantSeventhAddEleventh",
	DominantSeventhAddThirteenth:                "DominantSeventhAddThirteenth",
	DominantSeventhAddSharpFourth:               "DominantSeventhAddSharpFourth",
	DominantSeventhFlatFifthFlatNinth:           "DominantSeventhFlatFifthFlatNinth",
	DominantSeventhSharpFifthFlatNinth:          "DominantSeventhSharpFifthFlatNinth",
	DominantSeventhFlatNinth:                    "DominantSeventhFlatNinth",
	DominantSeventhFlatNinthFlatThirteenth:      "DominantSeventhFlatNinthFlatThirteenth",
	DominantSeventhSharpNinth:                   "DominantSeventhSharpNinth",
	DominantSeventhSharpNinthSharpEleventh:      "DominantSeventhSharpNinthSharpEleventh",
	DominantSeventhSharpEleventh:                "DominantSeventhSharpEleventh",
	DiminishedMajorSeventh:                      "DiminishedMajorSeventh",
	MajorSeventhFlatFifth:                       "MajorSeventhFlatFifth",
	MajorSeventhSuspendedSecond:                 "MajorSeventhSuspendedSecond",
	MajorSeventhSuspendedFourth:                 "MajorSeventhSuspendedFourth",
	MajorSeventhSuspendedFourthSharpFifth:       "MajorSeventhSuspendedFourthSharpFifth",
	MajorSeventhSuspendedFourthDoubleSharpFifth: "MajorSeventhSuspendedFourthDoubleSharpFifth",
	MajorSeventh:                                "MajorSeventh",
	MajorSeventhAddFourth:                       "MajorSeventhAddFourth",
	MajorSeventhAddEleventh:                     "MajorSeventhAddEleventh",
	MajorSeventhAddThirteenth:                   "MajorSeventhAddThirteenth",
	MajorSeventhAddSharpEleventh:                "MajorSeventhAddSharpEleventh",
	MajorSeventhAddSharpFourth:                  "MajorSeventhAddSharpFourth",
	MajorSeventhDoubleSharpFifth:                "MajorSeventhDoubleSharpFifth",
	AugmentedMajorSeventh:                       "AugmentedMajorSeventh",
	AugmentedAugmentedSeventh:                   "AugmentedAugmentedSeventh",

	// Ninth
	MinorNinth:                   "MinorNinth",
	MinorMajorNinth:              "MinorMajorNinth",
	DominantNinth:                "DominantNinth",
	DominantNinthSuspendedSecond: "DominantNinthSuspendedSecond",
	DominantNinthSuspendedFourth: "DominantNinthSuspendedFourth",
	DominantNinthSharpEleventh:   "DominantNinthSharpEleventh",
	DominantNinthFlatThirteenth:  "DominantNinthFlatThirteenth",
	MajorNinth:                   "MajorNinth",
	MajorNinthSuspendedSecond:    "MajorNinthSuspendedSecond",
	MajorNinthSuspendedFourth:    "MajorNinthSuspendedFourth",

	// Eleventh
	MinorEleventh:      "MinorEleventh",
	MinorMajorEleventh: "MinorMajorEleventh",
	DominantEleventh:   "DominantEleventh",
	MajorEleventh:      "MajorEleventh",

	// Thirteenth
	MinorThirteenth:             "MinorThirteenth",
	MinorMajorThirteenth:        "MinorMajorThirteenth",
	DominantThirteenth:          "DominantThirteenth",
	DominantThirteenthFlatNinth: "DominantThirteenthFlatNinth",
	MajorThirteenth:             "MajorThirteenth",
}

// String returns stringified chord type
func (t Type) String() string {
	return mapTypeToString[t]
}

// GoString satisfies GoStringer
func (t Type) GoString() string {
	return t.String()
}
