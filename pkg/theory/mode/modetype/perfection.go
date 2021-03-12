package modetype

// Perfection returns perfect and imperferct pitch indexes
func (t Type) Perfection() (int, int, []bool) {
	computedScale, shiftAmount := t.Scale()
	perfection, imperfection, perfectionProfile := computedScale.Perfection()

	r := (shiftAmount - 1) % len(perfectionProfile)
	shiftedProfile := append(perfectionProfile[shiftAmount-1:], perfectionProfile[:r]...)
	return perfection, imperfection, shiftedProfile
}
