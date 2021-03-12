package midi

import (
	"io"
)

// EventCommonControlChange is a control change event type
type EventCommonControlChange struct {
	DeltaTime uint64
	Channel   uint8
	Control   uint8
	Value     uint8
}

// Type returns event type
func (e EventCommonControlChange) Type() EventType {
	return EventTypeCommonControlChange
}

// Marshall marshalls event into writer
func (e EventCommonControlChange) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	_, err := writer.Write([]byte{
		0xb0 | (e.Channel & 0x0f),
		e.Control & 0x7f,
		e.Value & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonControlChange(deltaTime uint64, channel uint8, reader io.Reader) error {
	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}

	cc, vv := payload[0]&0x7f, payload[1]&0x7f
	t.Events = append(t.Events, EventCommonControlChange{deltaTime, channel, cc, vv})
	t.lastEventType = EventTypeCommonControlChange
	t.lastChannel = channel
	return nil
}
