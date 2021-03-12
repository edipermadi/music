package note

import "github.com/edipermadi/music/pkg/theory/limit"

// MIDINoteNumber retuns MIDI note number
func (n Note) MIDINoteNumber(octave int) int {
	ret := ((octave + 1) * limit.SemitonesInOneOctave) + noteOffsets[n]
	if ret < limit.MIDINoteNumberMin || ret > limit.MIDINoteNumberMax {
		return -1
	}
	return ret
}

// FromMIDINoteNumber return notes from MIDI note number
func FromMIDINoteNumber(v int) (Notes, error) {
	if v < limit.MIDINoteNumberMin || v > limit.MIDINoteNumberMax {
		return nil, ErrInvalidMIDINoteNumber
	}

	return mapMIDIToNotes[v], nil
}
