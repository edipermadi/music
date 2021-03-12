package chordtype

import "fmt"

// Number returns chord type numbering
func (t Type) Number() int {
	if v, found := mapTypeToNumber[t]; found {
		return v
	}

	panic(fmt.Errorf("no such number for chord type %s", t))
}
