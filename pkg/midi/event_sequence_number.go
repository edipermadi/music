package midi

import (
	"encoding/binary"
	"errors"
	"io"
)

// EventRealTimeSequenceNumber is a sequence number event type
type EventRealTimeSequenceNumber struct {
	DeltaTime      uint64
	SequenceNumber uint16
}

// Type returnevent type
func (e EventRealTimeSequenceNumber) Type() EventType {
	return EventTypeRealTimeSequenceNumber
}

// Marshall marshalls event into writer
func (e EventRealTimeSequenceNumber) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	if _, err := writer.Write([]byte{0xff, 0x00, 0x02}); err != nil {
		return err
	}

	return binary.Write(writer, binary.BigEndian, e.SequenceNumber)
}

func (t *Track) parseEventRealTimeSequenceNumber(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	if length != 2 {
		return errors.New("invalud sequence number length")
	}

	sequenceNumber, err := parseUint16(reader)
	if err != nil {
		return err
	}

	t.Events = append(t.Events, EventRealTimeSequenceNumber{DeltaTime: deltaTime, SequenceNumber: sequenceNumber})
	return nil
}
