package chordtype

// Type represents chord type
type Type int

// Types is a slice of chord type
type Types []Type

// Type enumeration
const (
	Invalid Type = iota

	// Dyad
	PowerChord Type = iota

	// Triad
	Diminished                            Type = iota
	DiminishedFlatThird                   Type = iota
	MinorDoubleFlatFifth                  Type = iota
	Minor                                 Type = iota
	MinorAddEleventh                      Type = iota
	MinorAddNinth                         Type = iota
	MinorAddFourth                        Type = iota
	MinorAddSharpFourth                   Type = iota
	MinorAddSharpNinth                    Type = iota
	MinorSharpFifth                       Type = iota
	MajorFlatFifth                        Type = iota
	Major                                 Type = iota
	MajorAddNinth                         Type = iota
	MajorAddEleventh                      Type = iota
	MajorDoubleSharpFifth                 Type = iota
	MajorAddFourth                        Type = iota
	MajorAddSharpFourth                   Type = iota
	MajorAddSharpNinth                    Type = iota
	Augmented                             Type = iota
	SuspendedSecondDoubleFlatFifth        Type = iota
	SuspendedSecondFlatFifth              Type = iota
	SuspendedSecondFlatFifthAddSharpFifth Type = iota
	SuspendedSecond                       Type = iota
	SuspendedSecondSharpFifth             Type = iota
	SuspendedFourthFlatFifth              Type = iota
	SuspendedFourth                       Type = iota
	SuspendedFourthSharpFifth             Type = iota
	SuspendedFourthDoubleSharpFifth       Type = iota
	Phrygian                              Type = iota
	PhrygianAddSeventh                    Type = iota
	Lydian                                Type = iota
	LydianMajorSeventh                    Type = iota
	Locrian                               Type = iota
	Quartal                               Type = iota
	QuartalAugmented                      Type = iota

	// Sixth
	MinorSixth                               Type = iota
	MinorSixthAddNinth                       Type = iota
	MinorSixthAddFlatNinth                   Type = iota
	MajorSixth                               Type = iota
	MajorSixthFlatFifth                      Type = iota
	MajorSixthAddFlatNinth                   Type = iota
	MajorSixthAddNinth                       Type = iota
	MajorSixthSuspendedSecond                Type = iota
	MajorSixthSuspendedSecondFlatFifth       Type = iota
	MajorSixthSuspendedSecondDoubleFlatFifth Type = iota
	MajorSixthSuspendedFourth                Type = iota

	// Seventh
	FullDiminishedSeventh                       Type = iota
	HalfDiminishedSeventh                       Type = iota
	MinorSeventhDoubleFlatFifth                 Type = iota
	MinorSeventh                                Type = iota
	MinorSeventhAddEleventh                     Type = iota
	MinorSeventhAddThirteenth                   Type = iota
	MinorSeventhAddSharpEleventh                Type = iota
	MinorSeventhSharpFifth                      Type = iota
	MinorSeventhFlatNinth                       Type = iota
	MinorMajorSeventh                           Type = iota
	MinorMajorSeventhAddEleventh                Type = iota
	MinorMajorSeventhAddThirteenth              Type = iota
	DominantSeventhFlatFifth                    Type = iota
	DominantSeventhSuspendedSecondFlatFifth     Type = iota
	DominantSeventhSuspendedSecond              Type = iota
	DominantSeventhSuspendedFourth              Type = iota
	DominantSeventh                             Type = iota
	DominantSeventhAddFourth                    Type = iota
	DominantSeventhAddSharpFourth               Type = iota
	DominantSeventhAddEleventh                  Type = iota
	DominantSeventhFlatFifthFlatNinth           Type = iota
	DominantSeventhSharpFifthFlatNinth          Type = iota
	DominantSeventhFlatNinth                    Type = iota
	DominantSeventhFlatNinthFlatThirteenth      Type = iota
	DominantSeventhSharpNinth                   Type = iota
	DominantSeventhSharpNinthSharpEleventh      Type = iota
	DominantSeventhSharpEleventh                Type = iota
	DominantSeventhEleventh                     Type = iota
	DominantSeventhAddThirteenth                Type = iota
	DiminishedMajorSeventh                      Type = iota
	MajorSeventhFlatFifth                       Type = iota
	MajorSeventhSuspendedSecond                 Type = iota
	MajorSeventhSuspendedFourth                 Type = iota
	MajorSeventhSuspendedFourthSharpFifth       Type = iota
	MajorSeventhSuspendedFourthDoubleSharpFifth Type = iota
	MajorSeventh                                Type = iota
	MajorSeventhAddFourth                       Type = iota
	MajorSeventhAddEleventh                     Type = iota
	MajorSeventhAddThirteenth                   Type = iota
	MajorSeventhAddSharpEleventh                Type = iota
	MajorSeventhAddSharpFourth                  Type = iota
	MajorSeventhDoubleSharpFifth                Type = iota
	AugmentedMajorSeventh                       Type = iota
	AugmentedAugmentedSeventh                   Type = iota

	// Ninth
	MinorNinth                   Type = iota
	MinorMajorNinth              Type = iota
	DominantNinth                Type = iota
	DominantNinthSuspendedSecond Type = iota
	DominantNinthSuspendedFourth Type = iota
	DominantNinthSharpEleventh   Type = iota
	DominantNinthFlatThirteenth  Type = iota
	MajorNinth                   Type = iota
	MajorNinthSuspendedSecond    Type = iota
	MajorNinthSuspendedFourth    Type = iota

	// Eleventh
	MinorEleventh      Type = iota
	MinorMajorEleventh Type = iota
	DominantEleventh   Type = iota
	MajorEleventh      Type = iota

	// Thirteenth
	MinorThirteenth             Type = iota
	MinorMajorThirteenth        Type = iota
	DominantThirteenth          Type = iota
	DominantThirteenthFlatNinth Type = iota
	MajorThirteenth             Type = iota
)

// AllTypes return all chord type as slice
func AllTypes() Types {
	return allTypes
}
