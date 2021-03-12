package note

// Equal returns true when two notes as equal
func (n Note) Equal(v Note) bool {
	return n.AbsoluteOffset() == v.AbsoluteOffset()
}
