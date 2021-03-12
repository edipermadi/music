package midi

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

// Header is MIDI decoder
type Header struct {
	Type           uint16
	NumberOfTracks uint16
	Division       uint16
}

// Unmarshall unmarshalls MIDI header from reader
func (h *Header) Unmarshall(reader io.Reader) error {
	if err := expectTag(reader, "MThd"); err != nil {
		return err
	}

	length, err := parseUint32(reader)
	if err != nil {
		return err
	}

	childReader := io.LimitReader(reader, int64(length))
	h.Type, err = parseUint16(childReader)
	if err != nil {
		return err
	}

	h.NumberOfTracks, err = parseUint16(childReader)
	if err != nil {
		return err
	}

	h.Division, err = parseUint16(childReader)
	if err != nil {
		return err
	}

	return nil
}

// Marshall marshalls header to writer
func (h Header) Marshall(writer io.Writer) error {
	var buff bytes.Buffer

	if err := binary.Write(&buff, binary.BigEndian, h.Type); err != nil {
		return err
	}

	if err := binary.Write(&buff, binary.BigEndian, h.NumberOfTracks); err != nil {
		return err
	}

	if err := binary.Write(&buff, binary.BigEndian, h.Division); err != nil {
		return err
	}

	body := buff.Bytes()

	// write header
	if _, err := writer.Write([]byte("MThd")); err != nil {
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

// File is midi file
type File struct {
	Header Header
	Tracks []Track
}

// Unmarshall unmarshall MIDI file from reader
func (f *File) Unmarshall(reader io.Reader) error {
	if err := f.Header.Unmarshall(reader); err != nil {
		return err
	}

	for i := 0; i < int(f.Header.NumberOfTracks); i++ {
		var track Track
		if err := track.Unmarshall(reader); err != nil {
			return err
		}

		f.Tracks = append(f.Tracks, track)
	}

	return nil
}

// Marshall marshalls file to writer
func (f File) Marshall(writer io.Writer) error {
	if err := f.Header.Marshall(writer); err != nil {
		return err
	}

	for _, track := range f.Tracks {
		if err := track.Marshall(writer); err != nil {
			return err
		}
	}
	return nil
}
func expectTag(reader io.Reader, expectedTag string) error {
	tag, err := parseTag(reader)
	if err != nil {
		return err
	}

	if tag != expectedTag {
		return ErrUnexpectedTag
	}

	return nil
}

func parseTag(reader io.Reader) (string, error) {
	tag := make([]byte, 4)
	length, err := io.ReadFull(reader, tag)
	if err != nil {
		return "", err
	}
	if length != 4 {
		return "", ErrInsufficientTagBytes
	}
	return string(tag), nil
}

func parseUint32(reader io.Reader) (uint32, error) {
	var val uint32
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return 0, err
	}

	return val, nil
}

func parseUint16(reader io.Reader) (uint16, error) {
	var val uint16
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return 0, err
	}

	return val, nil
}

func parseUint8(reader io.Reader) (uint8, error) {
	var val uint8
	if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
		return 0, err
	}

	return val, nil
}

func parseBytes(reader io.Reader, expectedLength int) ([]byte, error) {
	buffer := make([]byte, expectedLength)
	actualLength, err := io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}

	if actualLength != int(expectedLength) {
		return nil, ErrInsufficientChunkData
	}

	return buffer, nil
}

func parseZString(reader io.Reader) (string, error) {
	expectedLength, err := parseUint32(reader)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, expectedLength)
	actualLength, err := io.ReadFull(reader, buffer)
	if err != nil {
		return "", err
	}

	if actualLength != int(expectedLength) {
		return "", ErrInsufficientChunkData
	}

	return strings.Trim(string(buffer), string(rune(0))), nil
}

func parseString(reader io.Reader, expectedLength int) (string, error) {
	buffer := make([]byte, expectedLength)
	actualLength, err := io.ReadFull(reader, buffer)
	if err != nil {
		return "", err
	}

	if actualLength != expectedLength {
		return "", ErrInsufficientChunkData
	}

	return strings.Trim(string(buffer), string(rune(0))), nil
}

func parseBody(reader io.Reader) (io.Reader, error) {
	length, err := parseUint32(reader)
	if err != nil {
		return nil, err
	}

	return io.LimitReader(reader, int64(length)), nil
}

func parseVarInt(reader io.Reader) (uint64, error) {
	var value uint64
	var err error
	for {
		var v uint8
		v, err = parseUint8(reader)
		switch {
		case err != nil:
			return 0, err
		case v < 0x80:
			value = (value << 7) + uint64(v&0x7f)
			return value, nil
		default:
			value = (value << 7) + uint64(v&0x7f)
		}
	}
}
