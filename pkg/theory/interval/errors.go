package interval

import "errors"

var (
	// ErrInvalidSemitonesCount when count of semitones is negative
	ErrInvalidSemitonesCount = errors.New("invalid semitones count")

	// ErrNoSuchInterval when given semitones count does not have corresponding interval
	ErrNoSuchInterval = errors.New("no such interval")
)
