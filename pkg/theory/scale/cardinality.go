package scale

// Cardinality returns scale cardinality
func (s Scale) Cardinality() int {
	return len(s.Transposition())
}
