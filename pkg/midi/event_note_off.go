package midi

import (
	"io"

	"github.com/edipermadi/music/pkg/theory/pitch"
)

// EventCommonNoteOff is a note off common event type
type EventCommonNoteOff struct {
	DeltaTime uint64
	Channel   uint8
	Key       uint8
	Velocity  uint8
}

// Pitch return pitch
func (e EventCommonNoteOff) Pitch() pitch.Pitch {
	return mapKeyToPitch[e.Key&0x7f]
}

// Type returns event type
func (e EventCommonNoteOff) Type() EventType {
	return EventTypeCommonNoteOff
}

// Marshall marshalls event into writer
func (e EventCommonNoteOff) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0x80 | (e.Channel & 0x0f),
		e.Key & 0x7f,
		e.Velocity & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonNoteOff(deltaTime uint64, channel uint8, reader io.Reader) error {
	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}

	key, velocity := payload[0]&0x7f, payload[1]&0x7f
	t.Events = append(t.Events, EventCommonNoteOff{deltaTime, channel, key, velocity})
	t.lastEventType = EventTypeCommonNoteOff
	t.lastChannel = channel
	return nil
}
