package signature

// Signature is key signature
type Signature int

// Signature enumeration
const (
	Invalid Signature = iota

	CNaturalMajor Signature = iota // no accidental

	GNaturalMajor Signature = iota // F#
	FNaturalMajor Signature = iota // Bb

	DNaturalMajor Signature = iota // F# C#
	BFlatMajor    Signature = iota // Bb Eb

	ANaturalMajor Signature = iota // F# C# G#
	EFlatMajor    Signature = iota // Bb Eb Ab

	ENaturalMajor Signature = iota // F# C# G# D#
	AFlatMajor    Signature = iota // Bb Eb Ab Db

	BNaturalMajor Signature = iota // F# C# G# D# A#
	DFlatMajor    Signature = iota // Bb Eb Ab Db Gb

	FSharpMajor Signature = iota // F# C# G# D# A# E#
	GFlatMajor  Signature = iota // Bb Eb Ab Db Gb Cb

	CSharpMajor Signature = iota // F# C# G# D# A# E# B#
)

// AllSignatures returns all signature
func AllSignatures() []Signature {
	return allSignatures
}
