package chord

import (
	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/edipermadi/music/pkg/theory/note"
)

func init() {
	initMapRootToMapTypeToNotes()
}

var mapRootToMapTypeToNotes map[note.Note]map[chordtype.Type]note.Notes

func initMapRootToMapTypeToNotes() {
	mapRootToMapTypeToNotes = make(map[note.Note]map[chordtype.Type]note.Notes)
	for _, givenRoot := range note.AllSimpleNotes() {
		computedMapTypeToNotes := make(map[chordtype.Type]note.Notes)
		for _, givenType := range chordtype.AllTypes() {
			computedNotes := make(note.Notes, 0)
			for _, interval := range givenType.Intervals() {
				computedNotes = append(computedNotes, givenRoot.Raise(interval))
			}
			computedMapTypeToNotes[givenType] = computedNotes
		}
		mapRootToMapTypeToNotes[givenRoot] = computedMapTypeToNotes
	}
}
