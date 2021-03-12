package midi

import (
	"errors"
	"io"
)

// EventRealTimeMetaChannelPrefix stores channel-prefix metadata
type EventRealTimeMetaChannelPrefix struct {
	DeltaTime uint64
	Prefix    uint8
}

// Type returns event type
func (e EventRealTimeMetaChannelPrefix) Type() EventType {
	return EventTypeRealTimeMetaChannelPrefix
}

// Marshall marshalls event into writer
func (e EventRealTimeMetaChannelPrefix) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{0xff, 0x20, 0x01, e.Prefix})
	return err
}

func (t *Track) parseEventRealTimeMetaChannelPrefix(deltaTime uint64, reader io.Reader) error {
	length, err := parseUint8(reader)
	if err != nil {
		return err
	}

	if length != 1 {
		return errors.New("invalud channel-prefix length")
	}

	prefix, err := parseUint8(reader)
	if err != nil {
		return err
	}

	t.Events = append(t.Events, EventRealTimeMetaChannelPrefix{deltaTime, prefix})
	return nil
}
