package chordtype

import "sort"

func init() {
	initAllTypes()
	initMapTypeToNumber()
}

var allTypes Types

func initAllTypes() {
	allTypes = make([]Type, 0)
	for v := range mapTypeToString {
		if v != Invalid {
			allTypes = append(allTypes, v)
		}
	}

	// sort chord types
	sort.SliceStable(allTypes, func(i, j int) bool {
		return allTypes[i] < allTypes[j]
	})
}

var mapTypeToNumber map[Type]int

func initMapTypeToNumber() {
	mapTypeToNumber = make(map[Type]int)
	for _, givenType := range AllTypes() {
		intervals := givenType.Intervals()

		var number int
		for _, v := range intervals {
			number += (1 << v.Semitones())
		}

		mapTypeToNumber[givenType] = number
	}
}
