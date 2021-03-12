package midi

import (
	"errors"
	"io"
)

// EventRealTimeMetaTempo stores tempo metadata
type EventRealTimeMetaTempo struct {
	DeltaTime                 uint64
	MicroSecondPerQuarterNote uint32 // micro second per quarter note
}

// Type returns event type
func (e EventRealTimeMetaTempo) Type() EventType {
	return EventTypeRealTimeMetaTempo
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaTempo) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	hh := uint8(e.MicroSecondPerQuarterNote >> 16)
	mm := uint8(e.MicroSecondPerQuarterNote >> 8)
	ll := uint8(e.MicroSecondPerQuarterNote)
	_, err := writer.Write([]byte{0xff, 0x51, 0x03, hh, mm, ll})
	return err
}

func (t *Track) parseEventRealTimeMetaTempo(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	if length != 0x03 {
		return errors.New("invalid tempo length")
	}

	payload, err := parseBytes(reader, 3)
	if err != nil {
		return err
	}

	hh, mm, ll := payload[0], payload[1], payload[2]
	value := uint32(hh)<<16 + uint32(mm)<<8 + uint32(ll)
	t.Events = append(t.Events, EventRealTimeMetaTempo{deltaTime, value})
	return nil
}
