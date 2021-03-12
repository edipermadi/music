package midi

import (
	"io"
)

// EventType is event type
type EventType int

// EventType enumeration
const (
	EventTypeCommonNoteOff              EventType = iota
	EventTypeCommonNoteOn               EventType = iota
	EventTypeCommonKeyPressure          EventType = iota
	EventTypeCommonControlChange        EventType = iota
	EventTypeCommonProgramChange        EventType = iota
	EventTypeCommonChannelPressure      EventType = iota
	EventTypeCommonPitchWheelChange     EventType = iota
	EventTypeCommonSongPositionPointer  EventType = iota
	EventTypeCommonSelectSong           EventType = iota
	EventTypeCommonTune                 EventType = iota
	EventTypeRealTimeTimingClock        EventType = iota
	EventTypeRealTimeStart              EventType = iota
	EventTypeRealTimeContinue           EventType = iota
	EventTypeRealTimeStop               EventType = iota
	EventTypeRealTimeActiveSensing      EventType = iota
	EventTypeRealTimeSequenceNumber     EventType = iota
	EventTypeRealTimeMetaText           EventType = iota
	EventTypeRealTimeMetaCopyright      EventType = iota
	EventTypeRealTimeMetaTrackName      EventType = iota
	EventTypeRealTimeMetaInstrumentName EventType = iota
	EventTypeRealTimeMetaLyric          EventType = iota
	EventTypeRealTimeMetaMarker         EventType = iota
	EventTypeRealTimeMetaCuePoint       EventType = iota
	EventTypeRealTimeMetaChannelPrefix  EventType = iota
	EventTypeRealTimeMetaTempo          EventType = iota
	EventTypeRealTimeMetaOffset         EventType = iota
	EventTypeRealTimeMetaTimeSignature  EventType = iota
	EventTypeRealTimeMetaKeySignature   EventType = iota
	EventTypeRealTimeMetaEndOfTrack     EventType = iota
)

// Event is event interface
type Event interface {
	Type() EventType
	Marshall(writer io.Writer) error
}

func (t *Track) parseEventRealTimeMeta(deltaTime uint64, reader io.Reader) error {
	eventType, err := parseUint8(reader)
	switch {
	case err != nil:
		return err
	case eventType == 0x00:
		err = t.parseEventRealTimeSequenceNumber(deltaTime, reader)
	case eventType == 0x01:
		err = t.parseEventRealTimeMetaText(deltaTime, reader)
	case eventType == 0x02:
		err = t.parseEventRealTimeMetaCopyright(deltaTime, reader)
	case eventType == 0x03:
		err = t.parseEventRealTimeMetaTrackName(deltaTime, reader)
	case eventType == 0x04:
		err = t.parseEventRealTimeMetaInstrumentName(deltaTime, reader)
	case eventType == 0x05:
		err = t.parseEventRealTimeMetaLyric(deltaTime, reader)
	case eventType == 0x06:
		err = t.parseEventRealTimeMetaMarker(deltaTime, reader)
	case eventType == 0x07:
		err = t.parseEventRealTimeMetaCuePoint(deltaTime, reader)
	case eventType == 0x20:
		err = t.parseEventRealTimeMetaChannelPrefix(deltaTime, reader)
	case eventType == 0x51:
		err = t.parseEventRealTimeMetaTempo(deltaTime, reader)
	case eventType == 0x54:
		err = t.parseEventRealTimeMetaOffset(deltaTime, reader)
	case eventType == 0x58:
		err = t.parseEventRealTimeMetaTimeSignature(deltaTime, reader)
	case eventType == 0x59:
		err = t.parseEventRealTimeMetaKeySignature(deltaTime, reader)
	case eventType == 0x2f:
		err = t.parseEventRealTimeMetaEndOfTrack(deltaTime, reader)
	}
	return err
}
