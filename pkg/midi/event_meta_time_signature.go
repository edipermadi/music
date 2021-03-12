package midi

import (
	"fmt"
	"io"
)

// EventRealTimeMetaTimeSignature stores time-signature metadata
type EventRealTimeMetaTimeSignature struct {
	DeltaTime uint64
	NN        uint8
	DD        uint8
	CC        uint8
	BB        uint8
}

// Type returns event type
func (e EventRealTimeMetaTimeSignature) Type() EventType {
	return EventTypeRealTimeMetaTimeSignature
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaTimeSignature) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xff, 0x58, 0x04, e.NN, e.DD, e.CC, e.BB})
	return err
}

func (t *Track) parseEventRealTimeMetaTimeSignature(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	// time signature
	if length != 0x04 {
		return fmt.Errorf("invalid time signature length")
	}

	payload, err := parseBytes(reader, 4)
	if err != nil {
		return err
	}

	nn, dd, cc, bb := payload[0], payload[1], payload[2], payload[3]
	t.Events = append(t.Events, EventRealTimeMetaTimeSignature{deltaTime, nn, dd, cc, bb})
	return nil
}
