package soundfont

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Version is soundfont version
type Version struct {
	Major uint16
	Minor uint16
}

// Unmarshall unmarshal version from from reader
func (v *Version) Unmarshall(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	v.Major, err = parseUint16(reader)
	if err != nil {
		return err
	}

	v.Minor, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// SoundFont is a soundfond file
type SoundFont struct {
	Version              Version
	Engine               string
	Name                 string
	CreationDate         string
	Engineer             string
	Product              string
	Copyright            string
	Comments             string
	Tools                string
	Samples              []byte
	PresetHeaders        []PresetHeader
	PresetBags           []PresetBag
	PresetModulators     []PresetModulator
	PresetGenerators     []PresetGenerator
	Instruments          []Instrument
	InstrumentBags       []InstrumentBag
	InstrumentModulators []InstrumentModulator
	InstrumentGenerators []InstrumentGenerator
	SampleHeaders        []SampleHeader
}

// PresetHeader is present header
type PresetHeader struct {
	Name       string
	Preset     uint16
	Bank       uint16
	BagIndex   uint16
	Library    uint32
	Genre      uint32
	Morphology uint32
}

// Unmarshall unmarshall preset header from reader
func (p *PresetHeader) Unmarshall(reader io.Reader) error {
	var err error
	p.Name, err = parseString(reader, 20)
	if err != nil {
		return err
	}

	p.Preset, err = parseUint16(reader)
	if err != nil {
		return err
	}

	p.Bank, err = parseUint16(reader)
	if err != nil {
		return err
	}

	p.BagIndex, err = parseUint16(reader)
	if err != nil {
		return err
	}

	p.Library, err = parseUint32(reader)
	if err != nil {
		return err
	}

	p.Genre, err = parseUint32(reader)
	if err != nil {
		return err
	}

	p.Morphology, err = parseUint32(reader)
	if err != nil {
		return err
	}

	return nil
}

// PresetBag is preset bag
type PresetBag struct {
	GenIndex uint32
	ModIndex uint32
}

// Unmarshall unmarshall bag from reader
func (b *PresetBag) Unmarshall(reader io.Reader) error {
	var err error
	b.GenIndex, err = parseUint32(reader)
	if err != nil {
		return err
	}

	b.ModIndex, err = parseUint32(reader)
	if err != nil {
		return err
	}

	return nil
}

// PresetModulator is preset modulator
type PresetModulator struct {
	ModSrcOper    uint16
	ModDestOper   uint16
	ModAmount     uint16
	ModAmtSrcOper uint16
	ModTransOper  uint16
}

// Unmarshall unmarshall preset modulator from reader
func (m *PresetModulator) Unmarshall(reader io.Reader) error {
	var err error
	m.ModSrcOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	m.ModDestOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	m.ModAmount, err = parseUint16(reader)
	if err != nil {
		return err
	}

	m.ModAmtSrcOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	m.ModTransOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// PresetGenerator is preset generator
type PresetGenerator struct {
	GenOper       uint16
	GenAmountType uint16
}

// Unmarshall unmarshall generator from reader
func (g *PresetGenerator) Unmarshall(reader io.Reader) error {
	var err error
	g.GenOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	g.GenAmountType, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// Instrument is preset instrument
type Instrument struct {
	Name     string
	BagIndex uint16
}

// Unmarshall unmarshall instrument from reader
func (i *Instrument) Unmarshall(reader io.Reader) error {
	var err error
	i.Name, err = parseString(reader, 20)
	if err != nil {
		return err
	}

	i.BagIndex, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// InstrumentBag stores instrument bag detail
type InstrumentBag struct {
	InstrumentGenIndex uint16
	InstrumentModIndex uint16
}

// Unmarshall unmarshall instrument bag from reader
func (i *InstrumentBag) Unmarshall(reader io.Reader) error {
	var err error

	i.InstrumentGenIndex, err = parseUint16(reader)
	if err != nil {
		return err
	}

	i.InstrumentModIndex, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// InstrumentModulator stores instrument modulator detail
type InstrumentModulator struct {
	ModSrcOper    uint16
	ModDestOper   uint16
	ModAmount     uint16
	ModAmtSrcOper uint16
	ModTransOper  uint16
}

// Unmarshall unmarshall instrument modulator from reader
func (i *InstrumentModulator) Unmarshall(reader io.Reader) error {
	var err error

	i.ModSrcOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	i.ModDestOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	i.ModAmount, err = parseUint16(reader)
	if err != nil {
		return err
	}

	i.ModAmtSrcOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	i.ModTransOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// InstrumentGenerator stores instrument generator detail
type InstrumentGenerator struct {
	GenOper   uint16
	GenAmount uint16
}

// Unmarshall decodes instrument generator from reader
func (i *InstrumentGenerator) Unmarshall(reader io.Reader) error {
	var err error

	i.GenOper, err = parseUint16(reader)
	if err != nil {
		return err
	}

	i.GenAmount, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// SampleHeader stores sample header
type SampleHeader struct {
	Name            string
	Start           uint32
	End             uint32
	StartLoop       uint32
	EndLoop         uint32
	SampleRate      uint32
	OriginalPitch   uint8
	PitchCorrection uint8
	SampleLink      uint16
	SampleType      uint16
}

// Unmarshall decodes sample reader from reader
func (s *SampleHeader) Unmarshall(reader io.Reader) error {
	var err error
	s.Name, err = parseString(reader, 20)
	if err != nil {
		return err
	}

	s.Start, err = parseUint32(reader)
	if err != nil {
		return err
	}

	s.End, err = parseUint32(reader)
	if err != nil {
		return err
	}

	s.StartLoop, err = parseUint32(reader)
	if err != nil {
		return err
	}

	s.EndLoop, err = parseUint32(reader)
	if err != nil {
		return err
	}

	s.SampleRate, err = parseUint32(reader)
	if err != nil {
		return err
	}

	s.OriginalPitch, err = parseUint8(reader)
	if err != nil {
		return err
	}

	s.PitchCorrection, err = parseUint8(reader)
	if err != nil {
		return err
	}

	s.SampleLink, err = parseUint16(reader)
	if err != nil {
		return err
	}

	s.SampleType, err = parseUint16(reader)
	if err != nil {
		return err
	}

	return nil
}

// Unmarshall unmarshalls soundfont from reader
func (s *SoundFont) Unmarshall(reader io.Reader) error {
	err := s.parseRIFF(reader)
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseRIFF(reader io.Reader) error {
	if err := expectTag(reader, "RIFF"); err != nil {
		return err
	}

	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	err = s.parseSFBK(reader)
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseSFBK(reader io.Reader) error {
	if err := expectTag(reader, "sfbk"); err != nil {
		return err
	}

	var err error
	for err == nil {
		childReader, err := parseLIST(reader)
		if err != nil {
			break
		}

		tag, err := parseTag(childReader)
		switch {
		case err != nil:
			break
		case tag == "INFO":
			err = s.parseINFO(childReader)
		case tag == "sdta":
			err = s.parseSDTA(childReader)
		case tag == "pdta":
			err = s.parsePDTA(childReader)
		default:
			err = ErrUnexpectedTag
		}
	}

	// EOF marks ends of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseINFO(reader io.Reader) error {
	var err error
	for err == nil {
		var tag string
		tag, err = parseTag(reader)
		switch {
		case err != nil:
			break
		case tag == "ifil":
			err = s.Version.Unmarshall(reader)
		case tag == "isng":
			s.Engine, err = parseZString(reader)
		case tag == "INAM":
			s.Name, err = parseZString(reader)
		case tag == "ICRD":
			s.CreationDate, err = parseZString(reader)
		case tag == "IENG":
			s.Engineer, err = parseZString(reader)
		case tag == "IPRD":
			s.Product, err = parseZString(reader)
		case tag == "ICOP":
			s.Copyright, err = parseZString(reader)
		case tag == "ICMT":
			s.Comments, err = parseZString(reader)
		case tag == "ISFT":
			s.Tools, err = parseZString(reader)
		default:
			err = ErrUnexpectedTag
		}
	}

	// EOF marks ends of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseSDTA(reader io.Reader) error {
	if err := s.parseSMPL(reader); err != nil {
		return err
	}

	return nil
}

func (s *SoundFont) parsePDTA(reader io.Reader) error {
	var err error
	for err == nil {
		var tag string
		tag, err = parseTag(reader)
		switch {
		case err != nil:
			break
		case tag == "phdr":
			err = s.parsePHDR(reader)
		case tag == "pbag":
			err = s.parsePBAG(reader)
		case tag == "pmod":
			err = s.parsePMOD(reader)
		case tag == "pgen":
			err = s.parsePGEN(reader)
		case tag == "inst":
			err = s.parseINST(reader)
		case tag == "ibag":
			err = s.parseIBAG(reader)
		case tag == "imod":
			err = s.parseIMOD(reader)
		case tag == "igen":
			err = s.parseIGEN(reader)
		case tag == "shdr":
			err = s.parseSHDR(reader)
		default:
			err = ErrUnexpectedTag
		}
	}

	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parsePHDR(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var presetHeader PresetHeader
		err = presetHeader.Unmarshall(reader)
		if err == nil {
			s.PresetHeaders = append(s.PresetHeaders, presetHeader)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parsePBAG(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var presetBag PresetBag
		err = presetBag.Unmarshall(reader)
		if err == nil {
			s.PresetBags = append(s.PresetBags, presetBag)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parsePMOD(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var presetModulator PresetModulator

		err = presetModulator.Unmarshall(reader)
		if err == nil {
			s.PresetModulators = append(s.PresetModulators, presetModulator)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parsePGEN(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var presetGenerator PresetGenerator
		err = presetGenerator.Unmarshall(reader)
		if err == nil {
			s.PresetGenerators = append(s.PresetGenerators, presetGenerator)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseINST(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var instrument Instrument

		err = instrument.Unmarshall(reader)
		if err == nil {
			s.Instruments = append(s.Instruments, instrument)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseIBAG(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var instrumentBag InstrumentBag

		err = instrumentBag.Unmarshall(reader)
		if err == nil {
			s.InstrumentBags = append(s.InstrumentBags, instrumentBag)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseIMOD(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var instrumentModulator InstrumentModulator

		err = instrumentModulator.Unmarshall(reader)
		if err == nil {
			s.InstrumentModulators = append(s.InstrumentModulators, instrumentModulator)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseIGEN(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var instrumentGenerator InstrumentGenerator

		err = instrumentGenerator.Unmarshall(reader)
		if err == nil {
			s.InstrumentGenerators = append(s.InstrumentGenerators, instrumentGenerator)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseSHDR(reader io.Reader) error {
	var err error
	reader, err = parseBody(reader)
	if err != nil {
		return err
	}

	for err == nil {
		var sampleHeader SampleHeader
		err = sampleHeader.Unmarshall(reader)
		if err == nil {
			s.SampleHeaders = append(s.SampleHeaders, sampleHeader)
		}
	}

	// EOF marks end of stream
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (s *SoundFont) parseSMPL(reader io.Reader) error {
	if err := expectTag(reader, "smpl"); err != nil {
		return err
	}

	length, err := parseUint32(reader)
	if err != nil {
		return err
	}

	written, err := io.Copy(ioutil.Discard, io.LimitReader(reader, int64(length)))
	if err != nil {
		return err
	}

	if written != int64(length) {
		return fmt.Errorf("failed to discard %d bytes", length)
	}

	return nil
}

func parseLIST(reader io.Reader) (io.Reader, error) {
	if err := expectTag(reader, "LIST"); err != nil {
		return nil, err
	}

	return parseBody(reader)
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
	if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
		return 0, err
	}

	return val, nil
}

func parseUint16(reader io.Reader) (uint16, error) {
	var val uint16
	if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
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
