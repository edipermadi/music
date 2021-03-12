package alphabet

// Alphabet defines note alhabet
type Alphabet int

// Alphabet enumeration
const (
	Invalid Alphabet = iota
	C       Alphabet = iota
	D       Alphabet = iota
	E       Alphabet = iota
	F       Alphabet = iota
	G       Alphabet = iota
	A       Alphabet = iota
	B       Alphabet = iota
)

// Alphabets retuns slice of all alphabets
func Alphabets() []Alphabet {
	return []Alphabet{C, D, E, F, G, A, B}
}

// String returns stringified version of alphabet
func (a Alphabet) String() string {
	switch a {
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case A:
		return "A"
	case B:
		return "B"
	default:
		return "Invalid"
	}
}

// GoString satisfies go stringer
func (a Alphabet) GoString() string {
	return a.String()
}
