package note

import (
	"github.com/edipermadi/music/pkg/theory/interval"
)

// Interval returns interval between two notes
func (n Note) Interval(v Note) ([]interval.Interval, error) {
	v1 := n.MIDINoteNumber(4)

	var v2 int
	if n.AbsoluteOffset() <= v.AbsoluteOffset() {
		v2 = v.MIDINoteNumber(4)
	} else {
		v2 = v.MIDINoteNumber(5)
	}

	return interval.FromSemitones(v2 - v1)
}
