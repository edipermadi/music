package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/edipermadi/music/pkg/illustration"
	"github.com/edipermadi/music/pkg/theory/chord"
	"github.com/edipermadi/music/pkg/theory/chord/chordtype"
	"github.com/edipermadi/music/pkg/theory/mode"
	"github.com/edipermadi/music/pkg/theory/mode/modetype"
	"github.com/edipermadi/music/pkg/theory/note"
	"github.com/edipermadi/music/pkg/theory/pitch"
	"github.com/edipermadi/music/pkg/theory/scale"
	"github.com/edipermadi/music/pkg/theory/signature"
	"go.uber.org/zap"
)

var mapModeNumberToModes map[int]mode.Modes
var mapScaleToModes map[scale.Scale]mode.Modes
var mapChordNumberToModes map[int]mode.Modes

func init() {
	mapModeNumberToModes = make(map[int]mode.Modes)
	mapScaleToModes = make(map[scale.Scale]mode.Modes)
	mapChordNumberToModes = make(map[int]mode.Modes)

	for _, givenMode := range mode.AllModes() {
		computedNotes := givenMode.Notes()

		for _, givenRoot := range computedNotes {
			for _, givenChord := range chord.AllChords() {
				if givenChord.Root() == givenRoot &&
					(givenMode.Number()|givenChord.Number() == givenMode.Number()) {
					mapChordNumberToModes[givenChord.Number()] = append(mapChordNumberToModes[givenChord.Number()], givenMode)
				}
			}
		}
		mapScaleToModes[givenMode.Scale()] = append(mapScaleToModes[givenMode.Scale()], givenMode)
	}

	for _, givenMode := range mode.AllModes() {
		src := mapModeNumberToModes[givenMode.Number()]
		dst := append(src, givenMode)
		sort.SliceStable(dst, func(i, j int) bool {
			return dst[i].CanonicalNumber() < dst[j].CanonicalNumber()
		})
		mapModeNumberToModes[givenMode.Number()] = dst
	}

	for k, v := range mapChordNumberToModes {
		mapChordNumberToModes[k] = v.Unique().Sort()
	}

	for k, v := range mapScaleToModes {
		mapScaleToModes[k] = v.Unique().Sort()
	}
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logger.Sync() }()

	wikiDir := os.Getenv("WIKI_DIR")
	if wikiDir == "" {
		logger.Panic("please set WIKI_DIR", zap.Error(err))
	}

	generateMainPage(logger, wikiDir)
	generateScalesTable(logger, wikiDir)
	generateModesTable(logger, wikiDir)
	generateChordsTable(logger, wikiDir)
	generateScalePages(logger, wikiDir)
	generateModePages(logger, wikiDir)
	generatePitchClassPages(logger, wikiDir)
	generateModeImages(logger, wikiDir)
	generateChordImages(logger, wikiDir)
	generateChordPages(logger, wikiDir)
}

func generateMainPage(logger *zap.Logger, wikiDir string) {
	filename := fmt.Sprintf("%s/README.md", wikiDir)

	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# Documentation\n\n")
	_, _ = fmt.Fprintf(&buff, "This documentation is auto generated. Use `$ make docs` to regenerate\n\n")
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n")

	if err := compareAndWrite(filename, buff.Bytes()); err != nil {
		logger.Panic("failed to generate file", zap.String("filename", filename), zap.Error(err))
	}
}

func generateScalesTable(logger *zap.Logger, wikiDir string) {
	filename := fmt.Sprintf("%s/Scales.md", wikiDir)

	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# Scales\n\n")
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")

	_, _ = fmt.Fprintf(&buff, "## Index\n\n")
	_, _ = fmt.Fprintf(&buff, "| Scale | Cardinality | Transposition | Perfection | Imperfection |\n")
	_, _ = fmt.Fprintf(&buff, "|-------|-------------|---------------|------------|--------------|\n")

	for _, givenScale := range scale.AllScales() {
		var computedTransposition []string
		for _, givenInterval := range givenScale.Transposition() {
			computedTransposition = append(computedTransposition, strconv.Itoa(givenInterval))
		}

		perfection, imperfection, _ := givenScale.Perfection()
		_, _ = fmt.Fprintf(&buff, "| [%s](Scale%s.md) | %d | %s | %d | %d | \n", givenScale.String(), givenScale.String(), givenScale.Cardinality(), strings.Join(computedTransposition, ", "), perfection, imperfection)
	}

	if err := compareAndWrite(filename, buff.Bytes()); err != nil {
		logger.Panic("failed to generate file", zap.String("filename", filename), zap.Error(err))
	}
}

func generateModesTable(logger *zap.Logger, wikiDir string) {
	filename := fmt.Sprintf("%s/Modes.md", wikiDir)
	logger.Info("processing mode table", zap.String("filename", filename))

	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# Modes\n\n")
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")

	_, _ = fmt.Fprintf(&buff, "## Index\n\n")
	_, _ = fmt.Fprintf(&buff, "| Number | Scale | Mode | Transposition | Notes | Illustration | Audio |\n")
	_, _ = fmt.Fprintf(&buff, "|--------|-------|------|---------------|-------|--------------|-------|\n")

	allModes := mode.AllModes()
	sort.SliceStable(allModes, func(i int, j int) bool {
		return allModes[i].Number() < allModes[j].Number()
	})

	for _, givenMode := range allModes {
		if !givenMode.Tonic().Equal(note.CNatural) {
			continue
		}

		computedNotes := givenMode.Notes()
		_, _, perfectionProfile := givenMode.Perfection()
		var computedNoteNames []string
		for idx, computedNote := range computedNotes {
			if perfectionProfile[idx%len(perfectionProfile)] {
				computedNoteNames = append(computedNoteNames, computedNote.Name())
			} else {
				computedNoteNames = append(computedNoteNames, fmt.Sprintf("**%s**", computedNote.Name()))
			}
		}

		var computedIntervals []string
		for _, interval := range givenMode.Transposition() {
			computedIntervals = append(computedIntervals, strconv.Itoa(interval))
		}

		modeURL := fmt.Sprintf("[%d](https://ianring.com/musictheory/scales/%d)", givenMode.CanonicalNumber(), givenMode.CanonicalNumber())
		_, _ = fmt.Fprintf(&buff, "| %s | [%s](Scale%s.md) | [%s](Mode%s.md) | %s | %s | ![%s](Mode%s.png) | [midi](https://github.com/edipermadi/music/blob/main/docs/Mode%s.mid?raw=true) |\n", modeURL, givenMode.Scale(), givenMode.Scale(), givenMode.Type(), givenMode.Type(), strings.Join(computedIntervals, ", "), strings.Join(computedNoteNames, ", "), givenMode, givenMode, givenMode)
	}

	if err := compareAndWrite(filename, buff.Bytes()); err != nil {
		logger.Panic("failed to generate file", zap.String("filename", filename), zap.Error(err))
	}
}

func generateScalePages(logger *zap.Logger, wikiDir string) {
	for _, givenScale := range scale.AllScales() {
		filename := fmt.Sprintf("%s/Scale%s.md", wikiDir, givenScale.String())
		if err := generateScalePage(logger, filename, givenScale); err != nil {
			logger.Panic("failed to generate mode page", zap.String("scale", givenScale.String()), zap.Error(err))
		}
	}
}

func generateScalePage(logger *zap.Logger, filename string, givenScale scale.Scale) error {
	logger.Info("processing scale page", zap.String("filename", filename))

	var buff bytes.Buffer
	perfection, imperfection, perfectionProfile := givenScale.Perfection()
	_, _ = fmt.Fprintf(&buff, "# Scale %s\n\n", givenScale.String())
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")

	_, _ = fmt.Fprintf(&buff, "## Cardinality\n\n%d Notes\n\n", givenScale.Cardinality())

	_, _ = fmt.Fprintf(&buff, "## Perfection\n\n")
	_, _ = fmt.Fprintf(&buff, "- %d Perfect Pitch\n", perfection)
	_, _ = fmt.Fprintf(&buff, "- %d Imperfect Pitch\n", imperfection)
	_, _ = fmt.Fprintf(&buff, "- %v Perfection Profile\n", perfectionProfile)
	_, _ = fmt.Fprintf(&buff, "\n")

	_, _ = fmt.Fprintf(&buff, "## Modes\n\n")
	_, _ = fmt.Fprintf(&buff, "| Number | Mode | Notes | Illustration | Audio |\n")
	_, _ = fmt.Fprintf(&buff, "|--------|------|-------|--------------|-------|\n")

	allModes := mapScaleToModes[givenScale]
	sort.SliceStable(allModes, func(i, j int) bool {
		return allModes[i].CanonicalNumber() < allModes[j].CanonicalNumber()
	})

	givenTonic := note.CNatural
	for _, givenMode := range allModes {
		if givenMode.Tonic() != givenTonic {
			continue
		}

		computedModeNumber := givenMode.Number()
		_, _, perfectionProfile := givenMode.Perfection()
		computedNoteNames := make([]string, 0)
		for idx, givenNote := range givenMode.Notes() {
			if !perfectionProfile[idx%len(perfectionProfile)] {
				computedNoteNames = append(computedNoteNames, fmt.Sprintf("**%s**", givenNote.Name()))
			} else {
				computedNoteNames = append(computedNoteNames, givenNote.Name())
			}

		}
		computedModeURL := fmt.Sprintf("https://ianring.com/musictheory/scales/%d", computedModeNumber)
		_, _ = fmt.Fprintf(&buff, "| [%d](%s) | [%s](Mode%s.md) | %s | ![%s](Mode%s.png) | [midi](https://github.com/edipermadi/music/blob/main/docs/Mode%s.mid?raw=true) | \n", computedModeNumber, computedModeURL, givenMode.Type(), givenMode.Type(), strings.Join(computedNoteNames, ", "), givenMode, givenMode, givenMode)
	}

	return compareAndWrite(filename, buff.Bytes())
}

func generateModePages(logger *zap.Logger, wikiDir string) {
	for _, givenType := range modetype.AllTypes() {
		filename := fmt.Sprintf("%s/Mode%s.md", wikiDir, givenType)
		if err := generateModePage(logger, filename, givenType); err != nil {
			logger.Panic("failed to generate mode page", zap.String("scale", givenType.String()), zap.Error(err))
		}
	}
}

func generateModePage(logger *zap.Logger, filename string, givenType modetype.Type) error {
	logger.Info("processing mode file", zap.String("filename", filename))

	computedScale, _ := givenType.Scale()
	perfection, imperfection, perfectionProfile := givenType.Perfection()

	computedMode := mode.New(note.CNatural, givenType)
	modeNumber := computedMode.Number()
	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# Mode %s\n\n", givenType.String())
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")

	_, _ = fmt.Fprintf(&buff, "## Parent Scale\n\n[%s](Scale%s.md)\n\n", computedScale.String(), computedScale.String())
	_, _ = fmt.Fprintf(&buff, "## Number\n\n[%d](https://ianring.com/musictheory/scales/%d)\n\n", modeNumber, modeNumber)

	var intervals []string
	for _, interval := range givenType.Transposition() {
		intervals = append(intervals, strconv.Itoa(interval))
	}
	_, _ = fmt.Fprintf(&buff, "## Transposition\n\n%s\n\n", strings.Join(intervals, ", "))

	romanNumeralChords := chordRomanNumeralPattern(computedMode)
	_, _ = fmt.Fprintf(&buff, "## Chord Pattern\n\n%s\n\n", strings.Join(romanNumeralChords, ", "))

	_, _ = fmt.Fprintf(&buff, "## Perfection\n\n")
	_, _ = fmt.Fprintf(&buff, "- %d Perfect notes\n", perfection)
	_, _ = fmt.Fprintf(&buff, "- %d Perfect notes\n\n", imperfection)

	_, _ = fmt.Fprintf(&buff, "## Perfection Profile\n\n%v\n\n", perfectionProfile)

	_, _ = fmt.Fprintf(&buff, "## Permutations\n\n")
	_, _ = fmt.Fprintf(&buff, "| Tonic | Notes | Signature | Illustration | Audio |\n")
	_, _ = fmt.Fprintf(&buff, "|-------|-------|-----------|--------------|-------|\n")

	for _, givenTonic := range note.AllSimpleNotes() {
		givenMode := mode.New(givenTonic, givenType)
		givenTonic := givenMode.Tonic()
		computedNotes := givenMode.Notes()
		var noteNames []string
		for idx, computedNote := range computedNotes {
			if perfectionProfile[idx%len(perfectionProfile)] {
				noteNames = append(noteNames, computedNote.Name())
			} else {
				noteNames = append(noteNames, fmt.Sprintf("**%s**", computedNote.Name()))
			}
		}

		computedSignature := signature.FromNotes(computedNotes)
		_, _ = fmt.Fprintf(&buff, "| [%s](Mode%s.md) | %s | %s | ![%s](Mode%s.png) | [midi](https://github.com/edipermadi/music/blob/main/docs/Mode%s.mid?raw=true) |\n", givenTonic.Name(), givenMode, strings.Join(noteNames, ", "), strings.Join(computedSignature.Names(), ", "), givenMode, givenMode, givenMode)

	}

	return compareAndWrite(filename, buff.Bytes())
}

func generatePitchClassPages(logger *zap.Logger, wikiDir string) {
	for _, givenMode := range mode.AllModes() {
		filename := fmt.Sprintf("%s/Mode%s.md", wikiDir, givenMode)
		if err := generatePitchClassPage(logger, filename, givenMode); err != nil {
			logger.Panic("failed to generate pitch class page", zap.String("scale", givenMode.String()), zap.Error(err))
		}
		if err := generatePitchClassMIDI(logger, wikiDir, givenMode); err != nil {
			logger.Panic("failed to generate pitch class page", zap.String("scale", givenMode.String()), zap.Error(err))
		}
		if err := generatePitchClassCircleDiagram(logger, wikiDir, givenMode); err != nil {
			logger.Panic("failed to generate pitch class diagram", zap.String("scale", givenMode.String()), zap.Error(err))
		}
	}
}

func generatePitchClassMIDI(logger *zap.Logger, wikiDir string, givenMode mode.Mode) error {
	trackName := fmt.Sprintf("Mode%s", givenMode)
	fileName := fmt.Sprintf("%s/%s.mid", wikiDir, trackName)

	logger.Info("generating midi for pitch-class", zap.String("mode", givenMode.String()))

	var buff bytes.Buffer
	midiFile, err := illustration.SaveAsMIDI(trackName, givenMode.Pitches(4)...)
	if err != nil {
		return err
	}

	if err := midiFile.Marshall(&buff); err != nil {
		return err
	}

	return compareAndWrite(fileName, buff.Bytes())
}

func generatePitchClassCircleDiagram(logger *zap.Logger, wikiDir string, givenMode mode.Mode) error {
	diagramName := fmt.Sprintf("CircleMode%s", givenMode)
	fileName := fmt.Sprintf("%s/%s.dot", wikiDir, diagramName)

	logger.Info("generating circle-of-fifth for pitch-class", zap.String("mode", givenMode.String()))

	diagramTemplate := `
graph {

layout = circo;
mindist = .1

node [shape = circle, fontname = Helvetica, margin = 0, style = filled]
edge [style=invis]

subgraph 1 {
	E -- B -- Gb -- Db -- Ab -- Eb -- Bb -- F -- C -- G -- D -- A -- E
}

%s
}
`

	notes := []note.Note{
		note.ENatural, note.BNatural, note.GFlat, note.DFlat, note.AFlat, note.EFlat, note.BFlat,
		note.FNatural, note.CNatural, note.GNatural, note.DNatural, note.ANatural,
	}

	var parts []string
	for _, givenNote := range notes {
		if givenMode.Tonic().Equal(givenNote) {
			parts = append(parts, fmt.Sprintf("%s [fillcolor = cadetblue1];", givenNote.Name()))
		} else if givenMode.Notes().Contains(givenNote) {
			parts = append(parts, fmt.Sprintf("%s [fillcolor = gray];", givenNote.Name()))
		} else {
			parts = append(parts, fmt.Sprintf("%s [fillcolor = white];", givenNote.Name()))
		}
	}

	var buff bytes.Buffer
	if _, err := fmt.Fprintf(&buff, diagramTemplate, strings.Join(parts, "\n")); err != nil {
		return err
	}

	return compareAndWrite(fileName, buff.Bytes())
}

func generatePitchClassPage(logger *zap.Logger, filename string, givenMode mode.Mode) error {
	logger.Info("processing pitch class", zap.String("filename", filename))

	computedScale := givenMode.Scale()
	computedNotes := givenMode.Notes()
	perfection, imperfection, perfectionProfile := givenMode.Perfection()

	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# Mode %s\n\n", givenMode)
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")

	_, _ = fmt.Fprintf(&buff, "## Scale\n\n[%s](Scale%s.md)\n\n", computedScale.String(), computedScale.String())

	_, _ = fmt.Fprintf(&buff, "## Mode\n\n[%s](Mode%s.md)\n\n", givenMode.String(), givenMode.String())

	_, _ = fmt.Fprintf(&buff, "## Tonic\n\n%s\n\n", givenMode.Tonic().Name())

	_, _ = fmt.Fprintf(&buff, "## Signature\n\n%s\n\n", signature.FromNotes(computedNotes))

	var computedTransposition []string
	for _, interval := range givenMode.Transposition() {
		computedTransposition = append(computedTransposition, strconv.Itoa(interval))
	}
	_, _ = fmt.Fprintf(&buff, "## Transposition\n\n%s\n\n", strings.Join(computedTransposition, ", "))

	numeralNumeralChords := chordRomanNumeralPattern(givenMode)
	_, _ = fmt.Fprintf(&buff, "## Chord Pattern\n\n%s\n\n", strings.Join(numeralNumeralChords, ", "))

	_, _ = fmt.Fprintf(&buff, "## Perfection\n\n")
	_, _ = fmt.Fprintf(&buff, " - %d Perfect Notes\n\n", perfection)
	_, _ = fmt.Fprintf(&buff, " - %d Imperfect Notes\n\n", imperfection)

	_, _ = fmt.Fprintf(&buff, "## Notes\n\n")
	for idx, givenNote := range computedNotes {
		if perfectionProfile[idx%len(perfectionProfile)] {
			_, _ = fmt.Fprintf(&buff, "- %s\n", givenNote.Name())
		} else {
			_, _ = fmt.Fprintf(&buff, "- %s (Imperfect)\n", givenNote.Name())
		}
	}

	_, _ = fmt.Fprintf(&buff, "\n")

	_, _ = fmt.Fprintf(&buff, "## Illustration\n\n![%s](Mode%s.png)\n\n", givenMode, givenMode)

	_, _ = fmt.Fprintf(&buff, "## Diagram\n\n![%s](CircleMode%s.png)\n\n", givenMode, givenMode)

	_, _ = fmt.Fprintf(&buff, "## Relative Modes\n\n")
	_, _ = fmt.Fprintf(&buff, "| Number | Mode | Tonic | Notes | Illustration |\n")
	_, _ = fmt.Fprintf(&buff, "|--------|------|-------|-------|--------------|\n")
	for idx, expectedTonic := range givenMode.Notes() {
		if idx < givenMode.Cardinality() {
			for _, relativeMode := range mapModeNumberToModes[givenMode.Number()] {
				if relativeMode.Tonic() == expectedTonic {
					canonicalNumber := relativeMode.CanonicalNumber()
					computedNoteNames := relativeMode.Notes().Names()
					modeURL := fmt.Sprintf("https://ianring.com/musictheory/scales/%d", canonicalNumber)
					_, _ = fmt.Fprintf(&buff, "| [%d](%s) | [%s](Mode%s.md) | %s | %s | ![%s](Mode%s.png) |\n", canonicalNumber, modeURL, relativeMode.Type(), relativeMode.Type(), relativeMode.Tonic().Name(), strings.Join(computedNoteNames, ", "), relativeMode, relativeMode)
				}
			}
		}
	}

	_, _ = fmt.Fprintf(&buff, "## Relative Brightness\n\n")
	_, _ = fmt.Fprintf(&buff, "| Number | Mode | Tonic | Notes | Illustration |\n")
	_, _ = fmt.Fprintf(&buff, "|--------|------|-------|-------|--------------|\n")
	for idx, expectedTonic := range givenMode.Notes() {
		if idx < givenMode.Cardinality() {
			for _, relativeMode := range mapModeNumberToModes[givenMode.Number()] {
				if relativeMode.Tonic() == expectedTonic {
					canonicalNumber := relativeMode.CanonicalNumber()
					computedNoteNames := relativeMode.Notes().Names()
					modeURL := fmt.Sprintf("https://ianring.com/musictheory/scales/%d", canonicalNumber)
					_, _ = fmt.Fprintf(&buff, "| [%d](%s) | [%s](Mode%s.md) | %s | %s | ![%s](CircleMode%s.png) |\n", canonicalNumber, modeURL, relativeMode.Type(), relativeMode.Type(), relativeMode.Tonic().Name(), strings.Join(computedNoteNames, ", "), relativeMode, relativeMode)
				}
			}
		}
	}

	_, _ = fmt.Fprintf(&buff, "\n")
	_, _ = fmt.Fprintf(&buff, "## Chords\n\n")

	for i := 0; i < givenMode.Cardinality(); i++ {
		givenNote := computedNotes[i]
		_, _ = fmt.Fprintf(&buff, "### %s\n\n", givenNote.Name())
		_, _ = fmt.Fprintf(&buff, "| Number | Root | Name | Notes | Illustration | Audio |\n")
		_, _ = fmt.Fprintf(&buff, "|--------|------|------|-------|--------------|-------|\n")

		// list possible chords
		possibleChords := make(chord.Chords, 0)
		for _, givenChord := range chord.AllChords() {
			if givenChord.Root() == givenNote &&
				(givenMode.Number()|givenChord.Number() == givenMode.Number()) {
				possibleChords = append(possibleChords, givenChord)
			}
		}

		possibleChords = possibleChords.Sort()
		for _, givenChord := range possibleChords {
			chordNoteNames := givenChord.Notes().Names()
			_, _ = fmt.Fprintf(&buff, "| %d | %s | [%s](Chord%s.md) | %s | ![%s](Chord%sRootPosition.png) | [midi](Chord%sRootPosition.mid) |\n", givenChord.Number(), givenNote.Name(), givenChord.Name(), givenChord, strings.Join(chordNoteNames, ", "), givenChord.Name(), givenChord, givenChord)
		}

		_, _ = fmt.Fprintf(&buff, "\n")
	}

	return compareAndWrite(filename, buff.Bytes())
}

func generateModeImages(logger *zap.Logger, wikiDir string) {
	for _, givenMode := range mode.AllModes() {
		filename := fmt.Sprintf("%s/Mode%s.png", wikiDir, givenMode)
		if err := drawModeImage(logger, filename, givenMode); err != nil {
			logger.Panic("failed to generate pitch class image", zap.String("scale", givenMode.String()), zap.Error(err))
		}
	}
}

func drawModeImage(_ *zap.Logger, filename string, givenMode mode.Mode) error {
	// build pitch class
	var buffer bytes.Buffer
	computedPitches := givenMode.Pitches(4)
	if err := illustration.NewKeyboard(&buffer).Draw(computedPitches...); err != nil {
		return err
	}

	return compareAndWrite(filename, buffer.Bytes())
}

func generateChordsTable(logger *zap.Logger, wikiDir string) {
	filename := fmt.Sprintf("%s/Chords.md", wikiDir)
	logger.Info("processing chords table", zap.String("filename", filename))

	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# Chords\n\n")

	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")

	_, _ = fmt.Fprintf(&buff, "## Chord Index\n\n")

	_, _ = fmt.Fprintf(&buff, "| Number | Name | Notes | Illustration | Audio |\n")
	_, _ = fmt.Fprintf(&buff, "|--------|------|-------|--------------|-------|\n")

	givenRoot := note.CNatural
	allChords := chord.AllChords()
	sort.SliceStable(allChords, func(i int, j int) bool {
		return allChords[i].Number() < allChords[j].Number()
	})

	for _, givenChord := range allChords {
		if givenChord.Root() != givenRoot {
			continue
		}
		computedChordNotes := givenChord.Notes()

		noteNames := strings.Join(computedChordNotes.Names(), ",")
		_, _ = fmt.Fprintf(&buff, "| %d | [%s](ChordCNatural%s.md) | %s | ![%s](Chord%sRootPosition.png) | [midi](Chord%sRootPosition.mid) | \n", givenChord.Number(), givenChord.Type().Name(), givenChord.Type(), noteNames, givenChord, givenChord, givenChord)
	}

	if err := compareAndWrite(filename, buff.Bytes()); err != nil {
		panic(err)
	}
}

func generateChordImages(logger *zap.Logger, wikiDir string) {
	for _, givenChord := range chord.AllChords() {
		computedChordNotes := givenChord.Notes()
		for i := 0; i < len(computedChordNotes); i++ {
			rotated := givenChord.Inversion(i)
			if err := drawChordImage(logger, wikiDir, rotated); err != nil {
				logger.Panic("failed to draw chord", zap.Error(err))
			}
			if err := saveChordMIDI(logger, wikiDir, rotated); err != nil {
				logger.Panic("failed to save chord MIDI", zap.Error(err))
			}
		}
	}
}

func generateChordPages(logger *zap.Logger, wikiDir string) {
	for _, givenChord := range chord.AllChords() {
		chordPageFilename := fmt.Sprintf("%s/Chord%s.md", wikiDir, givenChord.String())
		if err := generateChordPage(logger, chordPageFilename, givenChord); err != nil {
			logger.Panic("failed to create chord documentation", zap.Error(err))
		}
	}
}

func saveChordMIDI(logger *zap.Logger, wikiDir string, givenChord chord.Chord) error {
	filename := fmt.Sprintf("%s/Chord%s%s%s.mid", wikiDir, givenChord.Root(), givenChord.Type(), givenChord.InversionName())
	logger.Info("processing chord MIDI", zap.String("filename", filename))

	var buff bytes.Buffer
	trackName := fmt.Sprintf("Chord%s%s%s", givenChord.Root(), givenChord.Type(), givenChord.InversionName())
	midiFile, err := illustration.SaveAsMIDI(trackName, pitch.PitchesFromNotes(4, givenChord.Notes()...)...)
	if err != nil {
		return err
	}

	if err := midiFile.Marshall(&buff); err != nil {
		return err
	}

	return compareAndWrite(filename, buff.Bytes())
}

func drawChordImage(logger *zap.Logger, wikiDir string, givenChord chord.Chord) error {
	filename := fmt.Sprintf("%s/Chord%s%s%s.png", wikiDir, givenChord.Root().String(), givenChord.Type().String(), givenChord.InversionName())
	logger.Info("processing chord image", zap.String("filename", filename))

	var buffer bytes.Buffer
	pitches := pitch.PitchesFromNotes(4, givenChord.Notes()...)
	if err := illustration.NewKeyboard(&buffer).Draw(pitches...); err != nil {
		return err
	}

	return compareAndWrite(filename, buffer.Bytes())
}

func generateChordPage(logger *zap.Logger, filename string, givenChord chord.Chord) error {
	logger.Info("processing chord page", zap.String("filename", filename))

	computedRoot := givenChord.Root()
	computedNotes := givenChord.Notes()

	var buff bytes.Buffer
	_, _ = fmt.Fprintf(&buff, "# %s\n\n", givenChord)
	_, _ = fmt.Fprintf(&buff, "## Links\n\n")
	_, _ = fmt.Fprintf(&buff, "- [Documentation](README.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Scales Index](Scales.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Modes Index](Modes.md)\n")
	_, _ = fmt.Fprintf(&buff, "- [Chords Index](Chords.md)\n\n")
	_, _ = fmt.Fprintf(&buff, "## Root\n\n%s\n\n", computedRoot.Name())
	_, _ = fmt.Fprintf(&buff, "## Notes\n\n")
	_, _ = fmt.Fprintf(&buff, "| Position | Notes | Illustration |\n")
	_, _ = fmt.Fprintf(&buff, "|----------|------|--------------|\n")
	for i := 0; i < len(computedNotes); i++ {
		rotated := givenChord.Inversion(i)
		noteNames := strings.Join(rotated.Notes().Names(), ",")
		_, _ = fmt.Fprintf(&buff, "| %s | %s | ![%s](Chord%s%s.png) |\n", rotated.InversionName(), noteNames, givenChord, givenChord, rotated.InversionName())
	}

	_, _ = fmt.Fprintf(&buff, "## Modes\n\n")

	_, _ = fmt.Fprintf(&buff, "| Number | Mode | Tonic | Notes | Illustration |\n")
	_, _ = fmt.Fprintf(&buff, "|--------|------|-------|-------|--------------|\n")

	allModes := mapChordNumberToModes[givenChord.Number()]
	sort.SliceStable(allModes, func(i, j int) bool {
		return allModes[i].CanonicalNumber() < allModes[j].CanonicalNumber()
	})

	for _, givenMode := range allModes {
		computedModeNames := givenMode.Notes().Names()

		canonicalModeNumber := givenMode.CanonicalNumber()
		modeURL := fmt.Sprintf("https://ianring.com/musictheory/scales/%d", canonicalModeNumber)
		_, _ = fmt.Fprintf(&buff, "| [%d](%s) | [%s](Mode%s.md) | %s | %s | ![Mode%s](Mode%s.png) |\n", canonicalModeNumber, modeURL, givenMode.Type(), givenMode, givenMode.Tonic().Name(), strings.Join(computedModeNames, ", "), givenMode, givenMode)
	}

	return compareAndWrite(filename, buff.Bytes())
}

func compareAndWrite(filename string, newPayload []byte) error {
	var proceed bool

	newChecksum := sha256.Sum256(newPayload)
	oldChecksum, err := computeFileHash(filename)
	switch {
	case errors.Is(err, os.ErrNotExist):
		proceed = true
	case err != nil:
		return err
	case hex.EncodeToString(newChecksum[:]) != oldChecksum:
		proceed = true
	}

	if !proceed {
		return nil
	}

	imageFile, err := os.Create(filename)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	defer func() { _ = imageFile.Close() }()
	_, err = imageFile.Write(newPayload)
	return err
}

func computeFileHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer func() { _ = file.Close() }()

	payload, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	digest := sha256.Sum256(payload)
	return hex.EncodeToString(digest[:]), nil
}

var diminishedRomanNumeralChords = map[int]string{
	0:  "i⁰",
	1:  "ii⁰",
	2:  "iii⁰",
	3:  "iv⁰",
	4:  "v⁰",
	5:  "vi⁰",
	6:  "vii⁰",
	7:  "viii⁰",
	8:  "ix⁰",
	9:  "x⁰",
	10: "xi⁰",
	11: "xii⁰",
}

var diminishedFlatThirdRomanNumeralChords = map[int]string{
	0:  "i⁰b3",
	1:  "ii⁰b3",
	2:  "iii⁰b3",
	3:  "iv⁰b3",
	4:  "v⁰b3",
	5:  "vi⁰b3",
	6:  "vii⁰b3",
	7:  "viii⁰b3",
	8:  "ix⁰b3",
	9:  "x⁰b3",
	10: "xi⁰b3",
	11: "xii⁰b3",
}

var minorRomanNumeralChords = map[int]string{
	0:  "i",
	1:  "ii",
	2:  "iii",
	3:  "iv",
	4:  "v",
	5:  "vi",
	6:  "vii",
	7:  "viii",
	8:  "ix",
	9:  "x",
	10: "xi",
	11: "xii",
}

var majorRomanNumeralChords = map[int]string{
	0:  "I",
	1:  "II",
	2:  "III",
	3:  "IV",
	4:  "V",
	5:  "VI",
	6:  "VII",
	7:  "VIII",
	8:  "IX",
	9:  "X",
	10: "XI",
	11: "XII",
}

var majorFlatFifthRomanNumeralChords = map[int]string{
	0:  "Ib5",
	1:  "IIb5",
	2:  "IIIb5",
	3:  "IVb5",
	4:  "Vb5",
	5:  "VIb5",
	6:  "VIIb5",
	7:  "VIIIb5",
	8:  "IXb5",
	9:  "Xb5",
	10: "XIb5",
	11: "XIIb5",
}

var augmentedRomanNumeralChords = map[int]string{
	0:  "I⁺",
	1:  "II⁺",
	2:  "III⁺",
	3:  "IV⁺",
	4:  "V⁺",
	5:  "VI⁺",
	6:  "VII⁺",
	7:  "VIII⁺",
	8:  "IX⁺",
	9:  "X⁺",
	10: "XI⁺",
	11: "XII⁺",
}

func chordRomanNumeralPattern(givenMode mode.Mode) []string {
	var numeralChords []string
	max := givenMode.Cardinality()
	notes := givenMode.Notes().Normalized()
	for i := 0; i < max; i++ {
		givenNote := notes[i]
		allowedChordNotes := note.Notes{notes[i], notes[(i+2)%max], notes[(i+4)%max]}
		allowedChordNotesChecksum := allowedChordNotes.Checksum()
		for _, givenChord := range chord.AllChords() {
			if (givenChord.Root() == givenNote) &&
				((givenMode.Number() | givenChord.Number()) == givenMode.Number()) &&
				givenChord.Number() == allowedChordNotesChecksum {
				switch givenChord.Type() {
				case chordtype.Diminished:
					numeralChords = append(numeralChords, diminishedRomanNumeralChords[i])
				case chordtype.DiminishedFlatThird:
					numeralChords = append(numeralChords, diminishedFlatThirdRomanNumeralChords[i])
				case chordtype.Minor:
					numeralChords = append(numeralChords, minorRomanNumeralChords[i])
				case chordtype.Major:
					numeralChords = append(numeralChords, majorRomanNumeralChords[i])
				case chordtype.MajorFlatFifth:
					numeralChords = append(numeralChords, majorFlatFifthRomanNumeralChords[i])
				case chordtype.Augmented:
					numeralChords = append(numeralChords, augmentedRomanNumeralChords[i])
				}
			}

		}
	}

	return numeralChords

}
