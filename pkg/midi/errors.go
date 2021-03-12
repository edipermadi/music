package midi

import "errors"

var (
	// ErrInsufficientChunkData when SF2 file does not contain enough payload as described in header
	ErrInsufficientChunkData = errors.New("insufficient chunk data")

	// ErrInsufficientTagBytes when available data is less than 4 bytes
	ErrInsufficientTagBytes = errors.New("insufficient data while reading tag")

	// ErrUnexpectedTag when tag is unexpected
	ErrUnexpectedTag = errors.New("unexpected tag")
)
