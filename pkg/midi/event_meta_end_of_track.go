package midi

import (
	"fmt"
	"io"
)

// EventRealTimeMetaEndOfTrack stores end-of-track metadata
type EventRealTimeMetaEndOfTrack struct {
	DeltaTime uint64
}

// Type returns event type
func (e EventRealTimeMetaEndOfTrack) Type() EventType {
	return EventTypeRealTimeMetaEndOfTrack
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaEndOfTrack) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xff, 0x2f, 0x00})
	return err
}

func (t *Track) parseEventRealTimeMetaEndOfTrack(deltaTime uint64, reader io.Reader) error {
	v, err := parseUint8(reader)
	if err != nil {
		return err
	}

	if v != 0x00 {
		return fmt.Errorf("unexpected end of track length")
	}

	t.Events = append(t.Events, EventRealTimeMetaEndOfTrack{deltaTime})
	return nil
}
