package midi

import (
	"io"
)

// EventRealTimeStop is a stop event type
type EventRealTimeStop struct {
	DeltaTime uint64
}

// Type returns event type
func (e EventRealTimeStop) Type() EventType {
	return EventTypeRealTimeStop
}

// Marshall marshalls event into writer
func (e EventRealTimeStop) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xfc})
	return err
}

func (t *Track) parseEventRealTimeStop(deltaTime uint64, reader io.Reader) error {
	t.Events = append(t.Events, EventRealTimeStop{deltaTime})
	t.lastEventType = EventTypeRealTimeStop
	return nil
}
