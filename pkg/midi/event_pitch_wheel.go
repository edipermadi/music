package midi

import (
	"io"
)

// EventCommonPitchWheelChange is a pitch wheel change event type
type EventCommonPitchWheelChange struct {
	DeltaTime uint64
	Channel   uint8
	Change    uint16
}

// Type returns event type
func (e EventCommonPitchWheelChange) Type() EventType {
	return EventTypeCommonPitchWheelChange
}

// Marshall marshalls event into writer
func (e EventCommonPitchWheelChange) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	hh := uint8(e.Change >> 7)
	ll := uint8(e.Change)
	_, err := writer.Write([]byte{
		0xe0 | (e.Channel & 0x0f),
		ll & 0x7f,
		hh & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonPitchWheelChange(deltaTime uint64, channel uint8, reader io.Reader) error {
	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}

	change := (uint16(payload[1]&0x7f) << 7) | uint16(payload[0]&0x7f)
	t.Events = append(t.Events, EventCommonPitchWheelChange{deltaTime, channel, change & 0x7fff})
	t.lastEventType = EventTypeCommonPitchWheelChange
	t.lastChannel = channel
	return nil
}
