package midi

import (
	"io"
)

// EventCommonTune is a tune event type
type EventCommonTune struct {
	DeltaTime uint64
}

// Type returnevent type
func (e EventCommonTune) Type() EventType {
	return EventTypeCommonTune
}

// Marshall marshalls event into writer
func (e EventCommonTune) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xf6})
	return err
}

func (t *Track) parseEventCommonTuneRequest(deltaTime uint64, reader io.Reader) error {
	t.Events = append(t.Events, EventCommonTune{deltaTime})
	t.lastEventType = EventTypeCommonTune
	return nil
}
