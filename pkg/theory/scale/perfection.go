package scale

// Perfection return perfection profile
func (s Scale) Perfection() (int, int, []bool) {
	activePitches := []bool{true, false, false, false, false, false, false, false, false, false, false, false}

	var currentIndex int
	for _, givenInterval := range s.Transposition() {
		currentIndex += givenInterval
		activePitches[currentIndex%len(activePitches)] = true
	}

	var perfect, imperfect int
	perfectionProfile := make([]bool, 0)
	for currentIndex, active := range activePitches {
		if !active {
			continue
		}

		nextFifth := (currentIndex + 7) % len(activePitches)
		if activePitches[nextFifth] {
			perfect++
			perfectionProfile = append(perfectionProfile, true)
		} else {
			imperfect++
			perfectionProfile = append(perfectionProfile, false)
		}
	}

	return perfect, imperfect, perfectionProfile
}
