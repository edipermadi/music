package modetype

import "sort"

// Types is a slice of types
type Types []Type

// Sort sorts list of mode
func (m Types) Sort() Types {
	sort.SliceStable(m, func(i int, j int) bool {
		return m[i] < m[j]
	})

	return m
}
