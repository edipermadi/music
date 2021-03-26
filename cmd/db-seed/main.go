package main

import (
	"fmt"
	"io"
	"os"

	"github.com/edipermadi/music/pkg/theory/pitch"
)

func main() {
	generatePitchSeed(os.Stdout)
}

func generatePitchSeed(writer io.Writer) {
	// put comment
	if _, err := fmt.Fprint(writer, "-- seed for pitches table\n"); err != nil{
		panic(err)
	}

	// generate seed
	for _, p := range pitch.AllPitches() {
		name := fmt.Sprintf("%s%d", p.Note().Alphabet(), p.Octave())
		if _, err := fmt.Fprintf(writer, "INSERT INTO pitches (number, octave, frequency, label, name) VALUES (%d, %d, %f, '%s', '%s');\n", p.MIDINoteNumber(), p.Octave(), p.Frequency(), p.String(), name); err != nil {
			panic(err)
		}
	}

	// trailing new line
	if _, err := fmt.Fprint(writer, "\n"); err != nil{
		panic(err)
	}
}
