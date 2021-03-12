package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/edipermadi/music/pkg/midi"
	"github.com/edipermadi/music/pkg/theory/mode"
	"github.com/edipermadi/music/pkg/theory/note"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logger.Sync() }()

	var fileIn string
	var modeOut string

	flag.StringVar(&fileIn, "in", "", "input MIDI file for processing")
	flag.StringVar(&modeOut, "mode", "", "desired output musical mode")
	flag.Parse()

	if fileIn == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(fileIn)
	if err != nil {
		logger.Fatal("failed to open MIDI input file", zap.Error(err))
	}

	defer func() { _ = file.Close() }()

	var midiFile midi.File
	if err := midiFile.Unmarshall(file); err != nil {
		logger.Fatal("failed to parse MIDI input file", zap.Error(err))
	}

	uniqueNoteMap := make(map[int]struct{})
	mapNoteCount := make(map[string]int)
	var notes note.Notes

	// analyze tracks
	for _, track := range midiFile.Tracks {
		for _, event := range track.Events {
			if event.Type() == midi.EventTypeCommonNoteOn {
				e := event.(midi.EventCommonNoteOn)
				pitch := e.Pitch()
				givenNote := pitch.Note()
				absoluteOffset := givenNote.AbsoluteOffset()
				mapNoteCount[givenNote.Name()] = mapNoteCount[givenNote.Name()] + 1
				if _, found := uniqueNoteMap[absoluteOffset]; !found {
					uniqueNoteMap[absoluteOffset] = struct{}{}
					notes = append(notes, pitch.Note())
				}
			}
		}
	}

	logger.Info(fmt.Sprintf("involved notes: %s", strings.Join(notes.Names(), ", ")))
	logger.Info("note statistic")
	for note, count := range mapNoteCount {
		logger.Info(fmt.Sprintf("note: %s (%d occurrence)", note, count))
	}
	notesChecksum := notes.Checksum()
	var computedMode mode.Mode
	for _, givenMode := range mode.AllModes() {
		givenModeNumber := givenMode.Number()
		if givenModeNumber|notesChecksum == givenModeNumber {
			logger.Info(fmt.Sprintf("source MIDI is playable in %s (%d) mode", strings.ToLower(givenMode.Type().String()), givenMode.CanonicalNumber()))
		}
		if givenMode.Number() == notesChecksum {
			computedMode = givenMode
			break
		}
	}

	// play in other mode
	for _, givenMode := range mode.AllModes() {
		if givenMode.Tonic() != computedMode.Tonic() {
			continue
		}

		if givenMode.Cardinality() >= computedMode.Cardinality() {
			if modeOut == "" || strings.EqualFold(modeOut, givenMode.Type().String()) {
				if err := modulateMIDI(fileIn, computedMode, givenMode); err != nil {
					logger.Fatal("failed to modulate MIDI", zap.Error(err))
				}
			}
		}
	}
}

func modulateMIDI(fileIn string, sourceMode mode.Mode, targetMode mode.Mode) error {
	file, err := os.Open(fileIn)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	var midiFile midi.File
	if err := midiFile.Unmarshall(file); err != nil {
		return err
	}

	sourceNotes := sourceMode.Notes()
	targetNotes := targetMode.Notes()
	lenTracks := len(midiFile.Tracks)
	for i := 0; i < lenTracks; i++ {
		lenEvents := midiFile.Tracks[i].Events
		for j := 0; j < len(lenEvents); j++ {
			event := midiFile.Tracks[i].Events[j]
			if event.Type() == midi.EventTypeCommonNoteOn {
				e := event.(midi.EventCommonNoteOn)
				pitchIn := e.Pitch()
				noteIn := pitchIn.Note()
				octaveIn := pitchIn.Octave()
				noteIndex := sourceNotes.IndexOf(noteIn)
				if noteIndex != -1 {
					noteOut := targetNotes[noteIndex]
					targetKey := ((octaveIn + 1) * 12) + noteOut.AbsoluteOffset()
					midiFile.Tracks[i].Events[j] = midi.EventCommonNoteOn{
						DeltaTime: e.DeltaTime,
						Channel:   e.Channel,
						Key:       uint8(targetKey) & 0x7f,
						Velocity:  e.Velocity,
					}
				} else {
					return fmt.Errorf("no such target note for (key = %d, note = %s, pitch = %s, mode = %s)", e.Key, noteIn, pitchIn, targetMode)
				}
			} else if event.Type() == midi.EventTypeCommonNoteOff {
				e := event.(midi.EventCommonNoteOff)
				pitchIn := e.Pitch()
				noteIn := pitchIn.Note()
				octaveIn := pitchIn.Octave()
				noteIndex := sourceNotes.IndexOf(noteIn)
				if noteIndex != -1 {
					noteOut := targetNotes[noteIndex]
					if noteOut.Alphabet() < noteIn.Alphabet() {
						octaveIn++
					}
					targetKey := (octaveIn * 12) + noteOut.AbsoluteOffset()
					midiFile.Tracks[i].Events[j] = midi.EventCommonNoteOff{
						DeltaTime: e.DeltaTime,
						Channel:   e.Channel,
						Key:       uint8(targetKey) & 0x7f,
						Velocity:  e.Velocity,
					}
				} else {
					return fmt.Errorf("no such target note for (key = %d, note = %s, pitch = %s, mode = %s)", e.Key, noteIn, pitchIn, targetMode)
				}
			}
		}
	}

	filename := filepath.Base(fileIn)
	extension := filepath.Ext(filename)
	template := filename[0 : len(filename)-len(extension)]

	fileOut, err := os.Create(fmt.Sprintf("%s.%d.%d.%s.midi", template, targetMode.Cardinality(), targetMode.CanonicalNumber(), strings.ToLower(targetMode.Type().String())))
	if err != nil {
		return err
	}

	defer func() { _ = fileOut.Close() }()

	if err := midiFile.Marshall(fileOut); err != nil {
		return err
	}

	return nil
}
