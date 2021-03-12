package midi

import (
	"io"
)

// EventRealTimeMetaCopyright stores metadata copyright
type EventRealTimeMetaCopyright struct {
	DeltaTime uint64
	Text      string
}

// Type returns event type
func (e EventRealTimeMetaCopyright) Type() EventType {
	return EventTypeRealTimeMetaCopyright
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaCopyright) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	if _, err := writer.Write([]byte{0xff, 0x02, byte(len(e.Text))}); err != nil {
		return err
	}

	_, err := writer.Write([]byte(e.Text))

	return err
}

func (t *Track) parseEventRealTimeMetaCopyright(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	text, err := parseString(reader, int(length))
	if err != nil {
		return err
	}

	t.Events = append(t.Events, EventRealTimeMetaCopyright{deltaTime, text})
	return nil
}
