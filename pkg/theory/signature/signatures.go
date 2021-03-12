package signature

// Signatures wraps a slice of signature
type Signatures []Signature

// Names return signature names
func (s Signatures) Names() []string {
	computedNames := make([]string, 0)
	for _, givenSignature := range s {
		computedNames = append(computedNames, givenSignature.Name())
	}
	return computedNames
}

// Unique filters out duplicated signature
func (s Signatures) Unique() Signatures {
	uniqueSignatureMap := make(map[Signature]struct{})
	uniqueSignatures := make(Signatures, 0)
	for _, givenSignature := range s {
		if _, found := uniqueSignatureMap[givenSignature]; !found {
			uniqueSignatureMap[givenSignature] = struct{}{}
			uniqueSignatures = append(uniqueSignatures, givenSignature)
		}
	}
	return uniqueSignatures
}
