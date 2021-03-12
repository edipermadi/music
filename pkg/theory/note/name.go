package note

import (
	"fmt"
)

// Name returns note name
func (n Note) Name() string {
	if v, found := mapNoteToName[n]; found {
		return v
	}
	panic(fmt.Errorf("no such name for note %s", n))
}
