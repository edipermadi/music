package illustration

import (
	"github.com/edipermadi/music/pkg/midi"
	"github.com/edipermadi/music/pkg/theory/pitch"
)

// SaveAsMIDI convert notes to MIDI File
func SaveAsMIDI(trackName string, givenPitches ...pitch.Pitch) (*midi.File, error) {
	events := []midi.Event{
		midi.EventRealTimeSequenceNumber{},
		midi.EventRealTimeMetaTrackName{DeltaTime: 0, Text: trackName},
		midi.EventRealTimeMetaChannelPrefix{},
		midi.EventRealTimeMetaInstrumentName{DeltaTime: 0, Text: "Grand Piano"},
		midi.EventCommonProgramChange{},
		midi.EventRealTimeMetaTempo{DeltaTime: 0, MicroSecondPerQuarterNote: 0xb71b0},
		midi.EventRealTimeMetaTimeSignature{DeltaTime: 0, NN: 4, DD: 2, CC: 0x18, BB: 0x18},
		midi.EventRealTimeMetaKeySignature{},
	}

	if len(givenPitches) > 0 {
		for _, givenPitch := range givenPitches {
			key := uint8(givenPitch.MIDINoteNumber())
			events = append(events,
				midi.EventCommonNoteOn{DeltaTime: 0x00, Channel: 0, Key: key, Velocity: 0x7e},
				midi.EventCommonNoteOff{DeltaTime: 0x1e0, Channel: 0, Key: key, Velocity: 0x00},
			)
		}

		// switch all notes on
		for _, givenPitch := range givenPitches {
			key := uint8(givenPitch.MIDINoteNumber())
			events = append(events, midi.EventCommonNoteOn{DeltaTime: 0x00, Channel: 0, Key: key, Velocity: 0x7e})
		}

		// switch all notes off
		length := len(givenPitches)
		for i := 0; i < length; i++ {
			key := uint8(givenPitches[length-i-1].MIDINoteNumber())
			var deltaTime uint64
			if i == 0 {
				deltaTime = 0x960
			}

			events = append(events, midi.EventCommonNoteOff{DeltaTime: deltaTime, Channel: 0, Key: key, Velocity: 0})
		}
	}

	// add end of track
	events = append(events, midi.EventRealTimeMetaEndOfTrack{DeltaTime: 0x00})

	var midiTrack midi.Track
	midiTrack.Events = append(midiTrack.Events)
	midiFile := midi.File{
		Header: midi.Header{Type: 0x01, NumberOfTracks: 0x01, Division: 0x1e0},
		Tracks: []midi.Track{{Events: events}},
	}
	return &midiFile, nil
}
