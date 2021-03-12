package midi

import (
	"fmt"
	"io"
)

// EventRealTimeMetaKeySignature stores key-signature metadata
type EventRealTimeMetaKeySignature struct {
	DeltaTime  uint64
	Accidental uint8
	Minor      uint8
}

// Type returns event type
func (e EventRealTimeMetaKeySignature) Type() EventType {
	return EventTypeRealTimeMetaKeySignature
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaKeySignature) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xff, 0x59, 0x02, e.Accidental, e.Minor & 0x01})
	return err
}

func (t *Track) parseEventRealTimeMetaKeySignature(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	// key signature
	if length != 0x02 {
		return fmt.Errorf("invalid key signature length")
	}

	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}
	accidental, minor := payload[0], payload[1]
	t.Events = append(t.Events, EventRealTimeMetaKeySignature{deltaTime, accidental, minor})
	return nil
}
