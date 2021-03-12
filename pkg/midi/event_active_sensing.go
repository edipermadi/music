package midi

import (
	"io"
)

// EventRealTimeActiveSensing is an active-sensing event type
type EventRealTimeActiveSensing struct {
	DeltaTime uint64
}

// Type returns event type
func (e EventRealTimeActiveSensing) Type() EventType {
	return EventTypeRealTimeActiveSensing
}

// Marshall marshalls event into writer
func (e EventRealTimeActiveSensing) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xfe})
	return err
}

func (t *Track) parseEventRealTimeActiveSensing(deltaTime uint64, reader io.Reader) error {
	t.Events = append(t.Events, EventRealTimeActiveSensing{deltaTime})
	t.lastEventType = EventTypeRealTimeActiveSensing
	return nil
}
