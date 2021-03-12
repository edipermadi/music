package midi

import (
	"io"
)

// EventCommonChannelPressure is a channel-pressure event type
type EventCommonChannelPressure struct {
	DeltaTime uint64
	Channel   uint8
	Pressure  uint8
}

// Type returns event type
func (e EventCommonChannelPressure) Type() EventType {
	return EventTypeCommonChannelPressure
}

// Marshall marshalls event into writer
func (e EventCommonChannelPressure) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0xd0 | (e.Channel & 0x0f),
		e.Pressure & 0x7f,
	})
	return err
}

func (t *Track) parseEventCommonChannelPressure(deltaTime uint64, channel uint8, reader io.Reader) error {
	pressure, err := parseUint8(reader)
	if err != nil {
		return err
	}

	t.Events = append(t.Events, EventCommonChannelPressure{deltaTime, channel, pressure & 0x7f})
	t.lastEventType = EventTypeCommonChannelPressure
	t.lastChannel = channel
	return nil
}
