package note

import "errors"

var (
	// ErrOctaveIsTooLow when octave is lower than -1
	ErrOctaveIsTooLow = errors.New("octave number is too low")

	// ErrOctaveIsTooHigh when octave is greater than 10
	ErrOctaveIsTooHigh = errors.New("octave number is too high")

	// ErrNoSuchMIDINote when a note does not have equivalent midi note number
	ErrNoSuchMIDINote = errors.New("no such midi note")

	// ErrInvalidMIDINoteNumber when midi note number is negative or greater than 127
	ErrInvalidMIDINoteNumber = errors.New("invalid midi note number")
)
