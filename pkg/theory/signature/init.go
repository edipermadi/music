package signature

import "sort"

var mapNumberToSignature map[int]Signatures

func init() {
	initAllSignatures()
	initMapNumberToSignature()
}

var allSignatures Signatures

func initAllSignatures() {
	allSignatures = make([]Signature, 0)
	for k := range mapSignatureToString {
		if k != Invalid {
			allSignatures = append(allSignatures, k)
		}
	}

	sort.SliceStable(allSignatures, func(i, j int) bool {
		return allSignatures[i] < allSignatures[j]
	})
}

func initMapNumberToSignature() {
	mapNumberToSignature = make(map[int]Signatures)
	for _, givenSignature := range AllSignatures() {
		computedNumber := Number(givenSignature.Notes())
		if _, found := mapNumberToSignature[computedNumber]; !found {
			mapNumberToSignature[computedNumber] = Signatures{givenSignature}
		} else {
			mapNumberToSignature[computedNumber] = append(mapNumberToSignature[computedNumber], givenSignature)
		}
	}
}
