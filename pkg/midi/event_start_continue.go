package midi

import (
	"io"
)

// EventRealTimeContinue is a continue event type
type EventRealTimeContinue struct {
	DeltaTime uint64
}

// Type returns event type
func (e EventRealTimeContinue) Type() EventType {
	return EventTypeRealTimeContinue
}

// Marshall marshalls event into writer
func (e EventRealTimeContinue) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xfb})
	return err
}

func (t *Track) parseEventRealTimeContinue(deltaTime uint64, reader io.Reader) error {
	t.Events = append(t.Events, EventRealTimeContinue{deltaTime})
	t.lastEventType = EventTypeRealTimeContinue
	return nil
}
