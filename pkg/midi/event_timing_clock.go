package midi

import (
	"io"
)

// EventRealTimeTimingClock is a timing-clock event type
type EventRealTimeTimingClock struct {
	DeltaTime uint64
}

// Type returnevent type
func (e EventRealTimeTimingClock) Type() EventType {
	return EventTypeRealTimeTimingClock
}

// Marshall marshalls event into writer
func (e EventRealTimeTimingClock) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xf8})
	return err
}

func (t *Track) parseEventRealTimeTimingClock(deltaTime uint64, reader io.Reader) error {
	t.Events = append(t.Events, EventRealTimeTimingClock{deltaTime})
	t.lastEventType = EventTypeRealTimeTimingClock
	return nil
}
