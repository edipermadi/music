package midi

import (
	"io"
)

// EventCommonKeyPressure is a key-pressure event type
type EventCommonKeyPressure struct {
	DeltaTime uint64
	Channel   uint8
	Key       uint8
	Velocity  uint8
}

// Type returns event type
func (e EventCommonKeyPressure) Type() EventType {
	return EventTypeCommonKeyPressure
}

// Marshall marshalls event into writer
func (e EventCommonKeyPressure) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0xa0 | (e.Channel & 0x0f),
		e.Key & 0x7f,
		e.Velocity & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonKeyPressure(deltaTime uint64, channel uint8, reader io.Reader) error {
	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}

	key, velocity := payload[0]&0x7f, payload[1]&0x7f
	t.Events = append(t.Events, EventCommonKeyPressure{deltaTime, channel, key, velocity})
	t.lastEventType = EventTypeCommonKeyPressure
	t.lastChannel = channel
	return nil
}
