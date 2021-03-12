package midi

import (
	"io"
)

// EventCommonSelectSong is a select song event type
type EventCommonSelectSong struct {
	DeltaTime uint64
	Song      uint8
}

// Type returnevent type
func (e EventCommonSelectSong) Type() EventType {
	return EventTypeCommonSelectSong
}

// Marshall marshalls event into writer
func (e EventCommonSelectSong) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0xf3,
		e.Song & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonSelectSong(deltaTime uint64, reader io.Reader) error {
	value, err := parseUint8(reader)
	if err != nil {
		return err
	}

	t.Events = append(t.Events, EventCommonSelectSong{deltaTime, value & 0x7f})
	t.lastEventType = EventTypeCommonSelectSong
	return nil
}
