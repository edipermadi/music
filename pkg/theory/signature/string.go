package signature

var mapSignatureToString = map[Signature]string{
	Invalid:       "Invalid",
	CNaturalMajor: "CNaturalMajor",
	GNaturalMajor: "GNaturalMajor",
	FNaturalMajor: "FNaturalMajor",
	DNaturalMajor: "DNaturalMajor",
	BFlatMajor:    "BFlatMajor",
	ANaturalMajor: "ANaturalMajor",
	EFlatMajor:    "EFlatMajor",
	ENaturalMajor: "ENaturalMajor",
	AFlatMajor:    "AFlatMajor",
	BNaturalMajor: "BNaturalMajor",
	DFlatMajor:    "DFlatMajor",
	FSharpMajor:   "FSharpMajor",
	GFlatMajor:    "GFlatMajor",
	CSharpMajor:   "CSharpMajor",
}

// String return stringified name
func (s Signature) String() string {
	return mapSignatureToString[s]
}

// GoString satisfies GoStringer
func (s Signature) GoString() string {
	return s.String()
}
