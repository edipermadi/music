package modetype

import (
	"sort"

	"github.com/edipermadi/music/pkg/theory/util"
)

func init() {
	initAllTypes()
	initMapTypeToIntervalPattern()
}

var allTypes Types

func initAllTypes() {
	for givenType := range mapModeToString {
		if givenType != Invalid {
			allTypes = append(allTypes, givenType)
		}
	}

	sort.SliceStable(allTypes, func(i int, j int) bool {
		return allTypes[i] < allTypes[j]
	})
}

var mapTypeToIntervalPattern map[Type][]int

func initMapTypeToIntervalPattern() {
	mapTypeToIntervalPattern = make(map[Type][]int)
	for _, givenType := range AllTypes() {
		computedScale, shiftAmount := givenType.Scale()
		mapTypeToIntervalPattern[givenType] = util.RotateLeftSliceInt(computedScale.IntervalPattern(), shiftAmount-1)
	}
}
