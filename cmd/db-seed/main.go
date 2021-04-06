package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/edipermadi/music/pkg/theory/alphabet"
	"github.com/edipermadi/music/pkg/theory/mode"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/edipermadi/music/pkg/theory/note/accidental"
	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/edipermadi/music/pkg/theory/scale"
)

func main() {
	generatePitchSeed(os.Stdout)
	generateAccidentalSeed(os.Stdout)
	generateAlphabetSeed(os.Stdout)
	generateNoteSeed(os.Stdout)
	generateNotePitchSeed(os.Stdout)
	generateScaleSeed(os.Stdout)
	generateModeSeed(os.Stdout)
}

var mapIDToPitch map[int]pitch.Pitch

func generatePitchSeed(writer io.Writer) {
	mapIDToPitch = make(map[int]pitch.Pitch)

	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for pitches table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for id, p := range pitch.AllPitches() {
		name := fmt.Sprintf("%s%d", p.Note().Alphabet(), p.Octave())
		if _, err := fmt.Fprintf(writer, "INSERT INTO pitches (number, octave, frequency, label, name) VALUES (%d, %d, %f, '%s', '%s');\n", p.MIDINoteNumber(), p.Octave(), p.Frequency(), p.String(), name); err != nil {
			panic(err)
		}
		mapIDToPitch[id+1] = p
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

func generateAccidentalSeed(writer io.Writer) {
	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for accidentals table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for _, a := range accidental.AllAccidentals() {
		if _, err := fmt.Fprintf(writer, "INSERT INTO accidentals (displacement, label, name) VALUES (%d, '%s', '%s');\n", a.Displacement(), a.String(), a.Name()); err != nil {
			panic(err)
		}
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

func generateAlphabetSeed(writer io.Writer) {
	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for alphabets table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for _, a := range alphabet.Alphabets() {
		if _, err := fmt.Fprintf(writer, "INSERT INTO alphabets (label) VALUES ('%s');\n", a.String()); err != nil {
			panic(err)
		}
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

var mapIDToNote map[int]note.Note

func generateNoteSeed(writer io.Writer) {
	mapIDToNote = make(map[int]note.Note)

	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for notes table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for i, n := range note.AllNotes() {
		if _, err := fmt.Fprintf(writer, "INSERT INTO notes (number, alphabet_id, accidental_id, label, name) VALUES (%d, %d, %d, '%s', '%s');\n", n.Number(), n.Alphabet(), n.Accidental(), n.String(), n.Name()); err != nil {
			panic(err)
		}
		mapIDToNote[i+1] = n
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

func generateNotePitchSeed(writer io.Writer) {
	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for note_pitches table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for octave := -1; octave < 11; octave++ {
		for _, n := range note.AllNotes() {
			switch {
			case octave == -1 && n == note.CFlat,
				octave == -1 && n == note.CDoubleFlat,
				octave == -1 && n == note.CTripleFlat,
				octave == -1 && n == note.DTripleFlat:
				continue
			case octave == 10 && n == note.BSharp,
				octave == 10 && n == note.BDoubleSharp,
				octave == 10 && n == note.BTripleSharp:
				continue
			default:
				p := pitch.Make(n, octave)
				noteID := getNoteID(n)
				pitchID := getPitchID(p)
				if _, err := fmt.Fprintf(writer, "INSERT INTO note_pitches (octave, note_id, pitch_id) VALUES (%d, %d, %d); -- %s %s\n", octave, noteID, pitchID, n, p); err != nil {
					panic(err)
				}
			}
		}
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

var mapIDToScale map[int]scale.Scale

func generateScaleSeed(writer io.Writer) {
	mapIDToScale = make(map[int]scale.Scale)

	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for scales table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for i, s := range scale.AllScales() {
		var transpositionBuff bytes.Buffer
		if err := json.NewEncoder(&transpositionBuff).Encode(s.Transposition()); err != nil {
			panic(err)
		}

		_, _, perfectionProfile := s.Perfection()
		var perfectionBuff bytes.Buffer
		if err := json.NewEncoder(&perfectionBuff).Encode(perfectionProfile); err != nil {
			panic(err)
		}

		transposition := strings.TrimSpace(transpositionBuff.String())
		perfection := strings.TrimSpace(perfectionBuff.String())
		if _, err := fmt.Fprintf(writer, "INSERT INTO scales (cardinality, transposition, perfection, label, name) VALUES (%d, '%s', '%s', '%s', '%s');\n", s.Cardinality(), transposition, perfection, s, s); err != nil {
			panic(err)
		}

		mapIDToScale[i+1] = s
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

func generateModeSeed(writer io.Writer) {
	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for modes table\n"); err != nil {
		panic(err)
	}

	// generate seed
	for _, m := range mode.AllModes() {
		var transpositionBuff bytes.Buffer
		if err := json.NewEncoder(&transpositionBuff).Encode(m.Transposition()); err != nil {
			panic(err)
		}

		_, _, perfectionProfile := m.Perfection()
		var perfectionBuff bytes.Buffer
		if err := json.NewEncoder(&perfectionBuff).Encode(perfectionProfile); err != nil {
			panic(err)
		}

		var noteIDs []int
		for _, n := range m.Notes() {
			noteIDs = append(noteIDs, getNoteID(n))
		}

		var notesBuff bytes.Buffer
		if err := json.NewEncoder(&notesBuff).Encode(noteIDs); err != nil {
			panic(err)
		}

		transposition := strings.TrimSpace(transpositionBuff.String())
		perfection := strings.TrimSpace(perfectionBuff.String())
		notes := strings.TrimSpace(notesBuff.String())
		if _, err := fmt.Fprintf(writer, "INSERT INTO modes (number, scale_id, tonic_note_id, transposition, perfection, notes, name, label) VALUES (%d, %d, %d, '%s', '%s', '%s', '%s', '%s');\n", m.Number(), getScaleID(m.Scale()), getNoteID(m.Tonic()), transposition, perfection, notes, m.String(), m.String()); err != nil {
			panic(err)
		}
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil {
		panic(err)
	}
}

func getNoteID(givenNote note.Note) int {
	for id, n := range mapIDToNote {
		if n == givenNote {
			return id
		}
	}
	return 0
}

func getPitchID(givenPitch pitch.Pitch) int {
	for id, n := range mapIDToPitch {
		if n == givenPitch {
			return id
		}
	}
	return 0
}

func getScaleID(givenScale scale.Scale) int {
	for id, n := range mapIDToScale {
		if n == givenScale {
			return id
		}
	}
	return 0
}
