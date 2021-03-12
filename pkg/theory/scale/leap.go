package scale

// Leap return largest semitone jump
func (s Scale) Leap() int {
	var leap int
	for _, givenInterval := range s.IntervalPattern() {
		if givenInterval > leap {
			leap = givenInterval
		}
	}

	return leap
}
