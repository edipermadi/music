package illustration

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/edipermadi/music/pkg/theory/pitch"
)

const pianoKeys = 37

// NewKeyboard returns an instance of keyboard scale illustrator
func NewKeyboard(writer io.Writer) Keyboard {
	numWhiteKeys := 22
	return Keyboard{
		margin:       2,
		width:        numWhiteKeys * 20,
		height:       72,
		numWhiteKeys: numWhiteKeys,
		writer:       writer,
	}
}

// Keyboard represents keybord scale illustrator
type Keyboard struct {
	margin       int
	width        int
	height       int
	numWhiteKeys int
	writer       io.Writer
}

type keyInfo struct {
	WhiteKeyIndex int
	IsBlackKey    bool
}

var pitchKeyDetails = map[pitch.Pitch]keyInfo{
	pitch.C4Natural: {0, false}, pitch.C4Sharp: {0, true}, pitch.D4Flat: {0, true},
	pitch.D4Natural: {1, false}, pitch.D4Sharp: {1, true}, pitch.E4Flat: {1, true},
	pitch.E4Natural: {2, false},
	pitch.F4Natural: {3, false}, pitch.F4Sharp: {3, true}, pitch.G4Flat: {3, true},
	pitch.G4Natural: {4, false}, pitch.G4Sharp: {4, true}, pitch.A4Flat: {4, true},
	pitch.A4Natural: {5, false}, pitch.A4Sharp: {5, true}, pitch.B4Flat: {5, true},
	pitch.B4Natural: {6, false},

	pitch.C5Natural: {7, false}, pitch.C5Sharp: {7, true}, pitch.D5Flat: {7, true},
	pitch.D5Natural: {8, false}, pitch.D5Sharp: {8, true}, pitch.E5Flat: {8, true},
	pitch.E5Natural: {9, false},
	pitch.F5Natural: {10, false}, pitch.F5Sharp: {10, true}, pitch.G5Flat: {10, true},
	pitch.G5Natural: {11, false}, pitch.G5Sharp: {11, true}, pitch.A5Flat: {11, true},
	pitch.A5Natural: {12, false}, pitch.A5Sharp: {12, true}, pitch.B5Flat: {12, true},
	pitch.B5Natural: {13, false},

	pitch.C6Natural: {14, false}, pitch.C6Sharp: {14, true}, pitch.D6Flat: {14, true},
	pitch.D6Natural: {15, false}, pitch.D6Sharp: {15, true}, pitch.E6Flat: {15, true},
	pitch.E6Natural: {16, false},
	pitch.F6Natural: {17, false}, pitch.F6Sharp: {17, true}, pitch.G6Flat: {17, true},
	pitch.G6Natural: {18, false}, pitch.G6Sharp: {18, true}, pitch.A6Flat: {18, true},
	pitch.A6Natural: {19, false}, pitch.A6Sharp: {19, true}, pitch.B6Flat: {19, true},
	pitch.B6Natural: {20, false},

	pitch.C7Natural: {21, false},
}

// Draw renders a keyboard with pressed key identified by red color
func (k Keyboard) Draw(pressedPitches ...pitch.Pitch) error {
	uniquePressedPitches := pitch.Pitches(pressedPitches).Unique()

	upLeft := image.Point{0, 0}
	lowRight := image.Point{
		X: k.width + k.margin*2,
		Y: k.height + k.margin*2,
	}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	allPitches := pitch.SimplePianoPitches(pianoKeys)

	// draw white keys
	for _, givenPitch := range allPitches {
		if givenPitch.Note().Accidental().Natural() {
			pressed := uniquePressedPitches.Contains(givenPitch)
			if err := k.drawKey(img, givenPitch, pressed); err != nil {
				return err
			}
		}
	}

	// draw black keys
	for _, givenPitch := range allPitches {
		if !givenPitch.Note().Accidental().Natural() {
			pressed := uniquePressedPitches.Contains(givenPitch)
			if err := k.drawKey(img, givenPitch, pressed); err != nil {
				return err
			}
		}
	}

	return png.Encode(k.writer, img)
}

func (k Keyboard) drawKey(img *image.RGBA, givenPitch pitch.Pitch, pressed bool) error {
	detail, found := pitchKeyDetails[givenPitch]
	if !found {
		return fmt.Errorf("no such key for pitch %s", givenPitch)
	}

	if detail.IsBlackKey {
		k.drawBlackKey(img, detail.WhiteKeyIndex, pressed)
	} else {
		k.drawWhiteKey(img, detail.WhiteKeyIndex, pressed)
	}

	return nil
}

func (k Keyboard) drawWhiteKey(img *image.RGBA, idx int, pressed bool) {
	border := color.RGBA{0, 0, 0, 255}
	foreground := color.RGBA{255, 255, 255, 255}

	pos := (idx * k.whiteKeyWidth())
	topLeft := image.Point{pos + k.margin, k.margin}
	bottomRight := image.Point{pos + k.whiteKeyWidth() + k.margin, k.whiteKeyHeight() + k.margin}

	if pressed {
		foreground = color.RGBA{255, 0, 0, 255}
	}

	k.drawBox(img, topLeft, bottomRight, border, foreground)
}

func (k Keyboard) drawBlackKey(img *image.RGBA, idx int, pressed bool) {
	border := color.RGBA{0, 0, 0, 255}
	foreground := color.RGBA{50, 50, 50, 255}

	offset := ((k.whiteKeyWidth() * 2) - k.blackKeyWidth()) / 2
	pos := (idx * k.whiteKeyWidth()) + offset
	topLeft := image.Point{pos + k.margin, k.margin}
	bottomRight := image.Point{pos + k.blackKeyWidth() + k.margin, k.blackKeyHeight() + k.margin}

	if pressed {
		foreground = color.RGBA{255, 0, 0, 255}
	}

	k.drawBox(img, topLeft, bottomRight, border, foreground)
}

func (k Keyboard) drawBox(img *image.RGBA, topLeft image.Point, bottomRight image.Point, border color.Color, foreground color.Color) {
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			switch {
			case y == topLeft.Y:
				img.Set(x, y, border)
			case y == bottomRight.Y:
				img.Set(x, y, border)
			case x == topLeft.X:
				img.Set(x, y, border)
			case x == bottomRight.X:
				img.Set(x, y, border)
			default:
				img.Set(x, y, foreground)
			}
		}
	}
}
func (k Keyboard) whiteKeyWidth() int {
	return k.width / k.numWhiteKeys
}

func (k Keyboard) whiteKeyHeight() int {
	return k.height
}

func (k Keyboard) blackKeyWidth() int {
	return k.whiteKeyWidth() / 2
}

func (k Keyboard) blackKeyHeight() int {
	return (k.whiteKeyHeight() * 2) / 3
}
