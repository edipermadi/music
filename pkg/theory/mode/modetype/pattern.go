package modetype

import (
	"fmt"
)

// IntervalPattern return interval pattern in semitones
func (t Type) IntervalPattern() []int {
	if v, found := mapTypeToIntervalPattern[t]; found {
		return v
	}

	panic(fmt.Errorf("no such interval pattern for mode type %s", t))
}
