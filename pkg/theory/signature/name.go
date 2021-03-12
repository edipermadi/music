package signature

import "fmt"

var mapSignatureToName = map[Signature]string{
	Invalid:       "Invalid",
	CNaturalMajor: "C",
	GNaturalMajor: "G",
	FNaturalMajor: "F",
	DNaturalMajor: "D",
	BFlatMajor:    "Bb",
	ANaturalMajor: "A",
	EFlatMajor:    "Eb",
	ENaturalMajor: "E",
	AFlatMajor:    "Ab",
	BNaturalMajor: "B",
	DFlatMajor:    "Db",
	FSharpMajor:   "F#",
	GFlatMajor:    "Gb",
	CSharpMajor:   "C#",
}

// Name return signature name
func (s Signature) Name() string {
	if v, found := mapSignatureToName[s]; found {
		return v
	}

	panic(fmt.Errorf("no such name for signature %s", s.String()))
}
