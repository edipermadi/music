package midi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Track is midi track
type Track struct {
	Events        []Event
	lastEventType EventType
	lastChannel   uint8
}

// Unmarshall unmarshalls track from reader
func (t *Track) Unmarshall(reader io.Reader) error {
	if err := expectTag(reader, "MTrk"); err != nil {
		return err
	}

	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var deltaTime uint64
		deltaTime, err = parseVarInt(reader)
		if err != nil {
			break
		}

		var event uint8
		event, err = parseUint8(reader)
		switch {
		case err != nil:
			break
		case event >= 0x80 && event <= 0x8f:
			err = t.parseEventCommonNoteOff(deltaTime, event&0x0f, reader)
		case event >= 0x90 && event <= 0x9f:
			err = t.parseEventCommonNoteOn(deltaTime, event&0x0f, reader)
		case event >= 0xa0 && event <= 0xaf:
			err = t.parseEventCommonKeyPressure(deltaTime, event&0x0f, reader)
		case event >= 0xb0 && event <= 0xbf:
			err = t.parseEventCommonControlChange(deltaTime, event&0x0f, reader)
		case event >= 0xc0 && event <= 0xcf:
			err = t.parseEventCommonProgramChange(deltaTime, event&0x0f, reader)
		case event >= 0xd0 && event <= 0xdf:
			err = t.parseEventCommonChannelPressure(deltaTime, event&0x0f, reader)
		case event >= 0xe0 && event <= 0xef:
			err = t.parseEventCommonPitchWheelChange(deltaTime, event&0x0f, reader)
		case event == 0xf2:
			err = t.parseEventCommonSongPositionPointer(deltaTime, reader)
		case event == 0xf3:
			err = t.parseEventCommonSelectSong(deltaTime, reader)
		case event == 0xf6:
			err = t.parseEventCommonTuneRequest(deltaTime, reader)
		case event == 0xf8:
			err = t.parseEventRealTimeTimingClock(deltaTime, reader)
		case event == 0xfa:
			err = t.parseEventRealTimeStart(deltaTime, reader)
		case event == 0xfb:
			err = t.parseEventRealTimeContinue(deltaTime, reader)
		case event == 0xfc:
			err = t.parseEventRealTimeStop(deltaTime, reader)
		case event == 0xfe:
			err = t.parseEventRealTimeActiveSensing(deltaTime, reader)
		case event == 0xff:
			err = t.parseEventRealTimeMeta(deltaTime, reader)
		default:
			err = t.parseRunningStatus(deltaTime, event, reader)
		}
	}

	// EOF signifies end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

// Marshall marshalls track into writer
func (t Track) Marshall(writer io.Writer) error {
	var buff bytes.Buffer
	for _, event := range t.Events {
		if err := event.Marshall(&buff); err != nil {
			return err
		}
	}

	// write header
	body := buff.Bytes()
	if _, err := writer.Write([]byte("MTrk")); err != nil {
		return err
	}

	// write body length
	if err := binary.Write(writer, binary.BigEndian, uint32(len(body))); err != nil {
		return err
	}

	// write body
	_, err := writer.Write(body)
	return err
}

func (t *Track) parseRunningStatus(deltaTime uint64, event uint8, reader io.Reader) error {
	switch t.lastEventType {
	case EventTypeCommonNoteOn:
		velocity, err := parseUint8(reader)
		if err != nil {
			return err
		}
		t.Events = append(t.Events, EventCommonNoteOn{DeltaTime: deltaTime, Channel: t.lastChannel, Key: event, Velocity: velocity})
	case EventTypeCommonNoteOff:
		velocity, err := parseUint8(reader)
		if err != nil {
			return err
		}
		t.Events = append(t.Events, EventCommonNoteOff{DeltaTime: deltaTime, Channel: t.lastChannel, Key: event, Velocity: velocity})
	default:
		return fmt.Errorf("unexpected event 0x%x", event)
	}

	return nil
}
