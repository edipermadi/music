package midi

import (
	"io"
)

// EventCommonProgramChange is a program change event type
type EventCommonProgramChange struct {
	DeltaTime uint64
	Channel   uint8
	Program   uint8
}

// Type returnevent type
func (e EventCommonProgramChange) Type() EventType {
	return EventTypeCommonProgramChange
}

// Marshall marshalls event into writer
func (e EventCommonProgramChange) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0xc0 | (e.Channel & 0x0f),
		e.Program & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonProgramChange(deltaTime uint64, channel uint8, reader io.Reader) error {
	program, err := parseUint8(reader)
	if err != nil {
		return err
	}

	t.Events = append(t.Events, EventCommonProgramChange{deltaTime, channel, program & 0x7f})
	t.lastEventType = EventTypeCommonProgramChange
	t.lastChannel = channel
	return nil
}
