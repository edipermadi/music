package midi

import (
	"io"
)

// EventRealTimeStart is a start event type
type EventRealTimeStart struct {
	DeltaTime uint64
}

// Type returnevent type
func (e EventRealTimeStart) Type() EventType {
	return EventTypeRealTimeStart
}

// Marshall marshalls event into writer
func (e EventRealTimeStart) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xfa})
	return err
}

func (t *Track) parseEventRealTimeStart(deltaTime uint64, reader io.Reader) error {
	t.Events = append(t.Events, EventRealTimeStart{deltaTime})
	t.lastEventType = EventTypeRealTimeStart
	return nil
}
