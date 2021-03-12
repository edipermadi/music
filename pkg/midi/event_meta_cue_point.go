package midi

import (
	"io"
)

// EventRealTimeMetaCuePoint stores metadata cue-point
type EventRealTimeMetaCuePoint struct {
	DeltaTime uint64
	Text      string
}

// Type returns event type
func (e EventRealTimeMetaCuePoint) Type() EventType {
	return EventTypeRealTimeMetaCuePoint
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaCuePoint) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	if _, err := writer.Write([]byte{0xff, 0x07, byte(len(e.Text))}); err != nil {
		return err
	}

	_, err := writer.Write([]byte(e.Text))

	return err
}

func (t *Track) parseEventRealTimeMetaCuePoint(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	text, err := parseString(reader, int(length))
	if err != nil {
		return err
	}
	t.Events = append(t.Events, EventRealTimeMetaCuePoint{deltaTime, text})
	return nil
}
