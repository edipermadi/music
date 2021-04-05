package main

import (
	"fmt"
	"github.com/edipermadi/music/pkg/theory/chord"
	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/edipermadi/music/pkg/theory/note"
	"os"
	"strings"
)

func main() {
	_, _ = fmt.Fprint(os.Stdout, `package chord

import (
	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/edipermadi/music/pkg/theory/note"
)

// Chord enumeration
var (
`)

	var chordNames []string
	for _, givenRoot := range note.AllSimpleNotes() {
		for _, givenType := range chordtype.AllTypes() {
			computedChord := chord.New(givenRoot, givenType)
			var notes []string
			for _, givenNote := range computedChord.Notes() {
				notes = append(notes, fmt.Sprintf("note.%s", givenNote))
			}
			_, _ = fmt.Fprintf(os.Stdout, "\t%s = Chord {note.%s, chordtype.%s, 0, %d, %d, \"%s\", note.Notes{%s}}\n", computedChord, givenRoot, givenType, computedChord.Number(), computedChord.CanonicalNumber(), computedChord, strings.Join(notes, ", "))
			chordNames = append(chordNames, computedChord.String())
		}
	}
	_, _ = fmt.Fprint(os.Stdout, ")\n\n")
	_, _ = fmt.Fprintf(os.Stdout, "var allChords = Chords {%s}\n", strings.Join(chordNames, ", "))
}
