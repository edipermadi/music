package midi

import (
	"io"
)

// EventRealTimeMetaLyric stores metadata lyric
type EventRealTimeMetaLyric struct {
	DeltaTime uint64
	Text      string
}

// Type returns event type
func (e EventRealTimeMetaLyric) Type() EventType {
	return EventTypeRealTimeMetaLyric
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaLyric) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	if _, err := writer.Write([]byte{0xff, 0x05, byte(len(e.Text))}); err != nil {
		return err
	}

	_, err := writer.Write([]byte(e.Text))

	return err
}

func (t *Track) parseEventRealTimeMetaLyric(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	text, err := parseString(reader, int(length))
	if err != nil {
		return err
	}
	t.Events = append(t.Events, EventRealTimeMetaLyric{deltaTime, text})
	return nil
}
