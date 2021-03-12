# music

Music theory exploration with golang

## Features

- General Fetures:
  - 12 semitones per octave system
  - [William Zeitler](https://allthescales.org) based scale and mode enumeration and naming
  - [Ian Ring](https://ianring.com/musictheory/scales/) based scale numbering
- Scale Features:
  - Support 228 scales, from 3 to 12 notes scales
  - Scale cardinality
  - Scale interval patterns
  - Modes in given scale
  - Perfection profile
- Mode Features:
  - Support 1489 modes out of 228 scales
  - Notes derivation
  - Chords derivation
  - Mode numbering (based on Ian Ring numbering system)
  - Relative modes lookup
  - Mode brightness (darken, brighten)
- Chord Features:
  - Supports 1836 chords, from 2 to 7 notes chords
  - Chord numbering (based on Ian Ring numbering system)
  - Modes lookup
  - Chord inversion
- Note Features:
  - Normalization
  - Enharmonic lookup
  - MIDI note values
- Illustration features:
  - Keyboard based visual illustration (37 keys)
  - MIDI and OGG based auditorial illustration
  - Mode audio visual illustration in various tonic notes
  - Chord audio visual illustration in various root notes
- Utilities:
  - Generate music cross reference markdown and images
  - Modulate MIDI from one mode to another
- MIDI features:
  - Marshall
  - Unmarshall

## Documentation

- [Music Cross Reference](docs/index.md)
