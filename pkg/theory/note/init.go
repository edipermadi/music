package note

import (
	"sort"
	"strings"

	"github.com/edipermadi/music/pkg/theory/limit"
)

func init() {
	initAllNotes()
	initMapNoteToName()
	initEnharmonic()
	initMapMidiToNotes()
	initMapOffsetToNotes()
}

var allNotes Notes

func initAllNotes() {
	allNotes = make(Notes, 0)
	for v := range mapNoteToString {
		if v != Invalid {
			allNotes = append(allNotes, v)
		}
	}

	// sort notes
	sort.SliceStable(allNotes, func(i, j int) bool {
		return allNotes[i] < allNotes[j]
	})
}

var mapNoteToEnharmonics map[Note][]Enharmonic

func initEnharmonic() {
	mapNoteToEnharmonics = make(map[Note][]Enharmonic, 0)
	for _, note1 := range AllNotes() {
		enharmonics := make([]Enharmonic, 0)
		for _, note2 := range AllNotes() {
			if note1 != note2 && note1.AbsoluteOffset() == note2.AbsoluteOffset() {
				enharmonic := Enharmonic{Note: note2, OctaveShift: note1.octaveShift() - note2.octaveShift()}
				enharmonics = append(enharmonics, enharmonic)
			}
		}

		mapNoteToEnharmonics[note1] = enharmonics
	}
}

var mapMIDIToNotes map[int]Notes

func initMapMidiToNotes() {
	mapMIDIToNotes = make(map[int]Notes, limit.MIDINoteNumberMax+1)
	for num1 := limit.MIDINoteNumberMin; num1 <= limit.MIDINoteNumberMax; num1++ {
		notes := make(Notes, 0)
		unique := make(map[Note]struct{})

		for octave := -1; octave < 11; octave++ {
			for _, n := range AllNotes() {
				if _, found := unique[n]; found {
					continue
				}

				num2 := n.MIDINoteNumber(octave)
				if num2 < 0 {
					continue
				}

				if num1 == num2 {
					notes = append(notes, n)
					unique[n] = struct{}{}
				}
			}
		}

		mapMIDIToNotes[num1] = notes
	}
}

var mapOffsetToNotes map[int]Notes

func initMapOffsetToNotes() {
	mapOffsetToNotes = make(map[int]Notes, 16)
	for i := 0; i < limit.SemitonesInOneOctave; i++ {
		uniqueNote := make(map[Note]struct{})
		notes := make(Notes, 0)

		for _, n := range AllNotes() {
			if _, found := uniqueNote[n]; found {
				continue
			}

			if i == n.AbsoluteOffset() {
				uniqueNote[n] = struct{}{}
				notes = append(notes, n)
			}
		}

		mapOffsetToNotes[i] = notes
	}
}

var mapNoteToName map[Note]string

func initMapNoteToName() {
	mapNoteToName = make(map[Note]string)
	for _, givenNote := range AllNotes() {
		var sb strings.Builder
		sb.WriteString(givenNote.Alphabet().String())
		sb.WriteString(givenNote.Accidental().Name())
		mapNoteToName[givenNote] = sb.String()
	}
}
