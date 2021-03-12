package midi

import (
	"errors"
	"io"
)

// EventRealTimeMetaOffset stores offset metadata
type EventRealTimeMetaOffset struct {
	DeltaTime uint64
	Hour      uint8
	Minute    uint8
	Second    uint8
	Frame     uint8
	Fraction  uint8 // hundredth of frame
}

// Type returns event type
func (e EventRealTimeMetaOffset) Type() EventType {
	return EventTypeRealTimeMetaOffset
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaOffset) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xff, 0x54, 0x05, e.Hour, e.Second, e.Second, e.Frame, e.Fraction})
	return err
}

func (t *Track) parseEventRealTimeMetaOffset(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	// offset
	if length != 0x05 {
		return errors.New("invalid offset length")
	}

	payload, err := parseBytes(reader, 5)
	if err != nil {
		return err
	}

	hh, mm, ss, fr, ff := payload[0], payload[1], payload[2], payload[3], payload[4]
	t.Events = append(t.Events, EventRealTimeMetaOffset{deltaTime, hh, mm, ss, fr, ff})
	return nil
}
