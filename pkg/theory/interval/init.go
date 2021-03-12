package interval

import (
	"sort"

	"github.com/edipermadi/music/pkg/theory/limit"
)

func init() {
	initAllIntervals()
	initSemitoneIntervals()
}

var allIntervals []Interval

func initAllIntervals() {
	allIntervals = make([]Interval, 0)
	for v := range intervalStringifieds {
		if v != Invalid {
			allIntervals = append(allIntervals, v)
		}
	}

	// sort slice
	sort.SliceStable(allIntervals, func(i, j int) bool {
		return allIntervals[i] < allIntervals[j]
	})
}

var mapSemitoneToIntervals map[int][]Interval

func initSemitoneIntervals() {
	mapSemitoneToIntervals = make(map[int][]Interval, limit.SemitonesInOneOctave+1)

	for i := 0; i <= limit.SemitonesInOneOctave; i++ {
		intervals := make([]Interval, 0)
		unique := make(map[Interval]struct{}, 0)
		for _, name := range AllIntervals() {
			if _, found := unique[name]; found {
				continue
			}

			if name.Semitones() == i {
				unique[name] = struct{}{}
				intervals = append(intervals, name)
			}
		}

		if i == 0 {
			intervals = append(intervals, PerfectOctave, AugmentedSeventh)
		}

		mapSemitoneToIntervals[i] = intervals
	}
}
