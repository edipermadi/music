package pitch

import "math"

// The frequency derivation is based on https://pages.mtu.edu/~suits/NoteFreqCalcs.html

// Frequency returns pitch frequency
func (p Pitch) Frequency() float64 {
	if p == A4Natural {
		return 440
	}

	mask := float64(1000000)
	a := math.Pow(2, 1.0/12.0)
	steps := p.MIDINoteNumber() - A4Natural.MIDINoteNumber()
	frequency := 440 * math.Pow(a, float64(steps))
	return math.Round(frequency*mask) / mask
}
