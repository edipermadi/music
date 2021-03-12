package midi

import (
	"io"

	"github.com/edipermadi/music/pkg/theory/pitch"
)

// EventCommonNoteOn is a note on common event type
type EventCommonNoteOn struct {
	DeltaTime uint64
	Channel   uint8
	Key       uint8
	Velocity  uint8
}

// Pitch return pitch
func (e EventCommonNoteOn) Pitch() pitch.Pitch {
	return mapKeyToPitch[e.Key&0x7f]
}

// Type returns event type
func (e EventCommonNoteOn) Type() EventType {
	return EventTypeCommonNoteOn
}

// Marshall marshalls event into writer
func (e EventCommonNoteOn) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0x90 | (e.Channel & 0x0f),
		e.Key & 0x7f,
		e.Velocity & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonNoteOn(deltaTime uint64, channel uint8, reader io.Reader) error {
	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}

	key, velocity := payload[0]&0x7f, payload[1]&0x7f
	t.Events = append(t.Events, EventCommonNoteOn{deltaTime, channel, key, velocity})
	t.lastEventType = EventTypeCommonNoteOn
	t.lastChannel = channel
	return nil
}
