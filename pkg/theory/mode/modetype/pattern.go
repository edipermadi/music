package modetype

import (
	"github.com/edipermadi/music/pkg/theory/util"
)

// Transposition return transposition in semitones
func (t Type) Transposition() []int {
	computedScale, shiftAmount := t.Scale()
	return util.RotateLeftSliceInt(computedScale.Transposition(), shiftAmount-1)
}
