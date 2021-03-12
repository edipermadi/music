package midi

import (
	"io"
)

// EventCommonSongPositionPointer is a song pointer ( 1 beat = 6 MIDI clock)
type EventCommonSongPositionPointer struct {
	DeltaTime uint64
	Pointer   uint16
}

// Type returns event type
func (e EventCommonSongPositionPointer) Type() EventType {
	return EventTypeCommonSongPositionPointer
}

// Marshall marshalls event into writer
func (e EventCommonSongPositionPointer) Marshall(writer io.Writer) error {
	if err := marshallVarInt(writer, e.DeltaTime); err != nil {
		return err
	}

	hh := uint8(e.Pointer >> 7)
	ll := uint8(e.Pointer)
	_, err := writer.Write([]byte{
		0xf2,
		ll & 0x7f,
		hh & 0x7f,
	})

	return err
}

func (t *Track) parseEventCommonSongPositionPointer(deltaTime uint64, reader io.Reader) error {
	payload, err := parseBytes(reader, 2)
	if err != nil {
		return err
	}

	pointer := (uint16(payload[1]&0x7f) << 7) | uint16(payload[0]&0x7f)
	t.Events = append(t.Events, EventCommonSongPositionPointer{deltaTime, pointer & 0x7fff})
	t.lastEventType = EventTypeCommonSongPositionPointer
	return nil
}
