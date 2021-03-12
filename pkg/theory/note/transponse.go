package note

import (
	"fmt"

	"github.com/edipermadi/music/pkg/theory/interval"
	"github.com/edipermadi/music/pkg/theory/limit"
)

// Transpose shifts note by given semitones
func (n Note) Transpose(semitones int) Notes {
	if semitones == 0 {
		return Notes{n}
	}

	offset := (n.AbsoluteOffset() + semitones) % limit.SemitonesInOneOctave
	return mapOffsetToNotes[offset]
}

var raisedNotes = map[Note]map[interval.Interval]Note{
	CNatural: {
		interval.PerfectUnison:   CNatural,
		interval.AugmentedUnison: CSharp,

		interval.DiminishedSecond: DDoubleFlat,
		interval.MinorSecond:      DFlat,
		interval.MajorSecond:      DNatural,
		interval.AugmentedSecond:  DSharp,

		interval.DiminishedThird: EDoubleFlat,
		interval.MinorThird:      EFlat,
		interval.MajorThird:      ENatural,
		interval.AugmentedThird:  ESharp,

		interval.DiminishedFourth: FFlat,
		interval.PerfectFourth:    FNatural,
		interval.AugmentedFourth:  FSharp,

		interval.DiminishedFifth: GFlat,
		interval.PerfectFifth:    GNatural,
		interval.AugmentedFifth:  GSharp,

		interval.DiminishedSixth: ADoubleFlat,
		interval.MinorSixth:      AFlat,
		interval.MajorSixth:      ANatural,
		interval.AugmentedSixth:  ASharp,

		interval.DiminishedSeventh: BDoubleFlat,
		interval.MinorSeventh:      BFlat,
		interval.MajorSeventh:      BNatural,
		interval.AugmentedSeventh:  BSharp,

		interval.DiminishedOctave: CFlat,
		interval.PerfectOctave:    CNatural,
		interval.AugmentedOctave:  CSharp,

		interval.DiminishedNinth: DDoubleFlat,
		interval.MinorNinth:      DFlat,
		interval.MajorNinth:      DNatural,
		interval.AugmentedNinth:  DSharp,

		interval.DiminishedTenth: EDoubleFlat,
		interval.MinorTenth:      EFlat,
		interval.MajorTenth:      ENatural,
		interval.AugmentedTenth:  ESharp,

		interval.DiminishedEleventh: FFlat,
		interval.PerfectEleventh:    FNatural,
		interval.AugmentedEleventh:  FSharp,

		interval.DiminishedTwelfth: GFlat,
		interval.PerfectTwelfth:    GNatural,
		interval.AugmentedTwelfth:  GSharp,

		interval.DiminishedThirteenth: ADoubleFlat,
		interval.MinorThirteenth:      AFlat,
		interval.MajorThirteenth:      ANatural,
		interval.AugmentedThirteenth:  ASharp,

		interval.DiminishedFourteenth: BDoubleFlat,
		interval.MinorFourteenth:      BFlat,
		interval.MajorFourteenth:      BNatural,
		interval.AugmentedFourteenth:  BSharp,

		interval.DiminihsedFifteenth: CFlat,
		interval.PerfectFifteenth:    CNatural,
		interval.AugmentedFifteenth:  CSharp,
	},
	CSharp: {
		interval.PerfectUnison:   CSharp,
		interval.AugmentedUnison: CDoubleSharp,

		interval.DiminishedSecond: DFlat,
		interval.MinorSecond:      DNatural,
		interval.MajorSecond:      DSharp,
		interval.AugmentedSecond:  DDoubleSharp,

		interval.DiminishedThird: EFlat,
		interval.MinorThird:      ENatural,
		interval.MajorThird:      ESharp,
		interval.AugmentedThird:  EDoubleSharp,

		interval.DiminishedFourth: FNatural,
		interval.PerfectFourth:    FSharp,
		interval.AugmentedFourth:  FDoubleSharp,

		interval.DiminishedFifth: GNatural,
		interval.PerfectFifth:    GSharp,
		interval.AugmentedFifth:  GDoubleSharp,

		interval.DiminishedSixth: AFlat,
		interval.MinorSixth:      ANatural,
		interval.MajorSixth:      ASharp,
		interval.AugmentedSixth:  ADoubleSharp,

		interval.DiminishedSeventh: BFlat,
		interval.MinorSeventh:      BNatural,
		interval.MajorSeventh:      BSharp,
		interval.AugmentedSeventh:  BDoubleSharp,

		interval.DiminishedOctave: CNatural,
		interval.PerfectOctave:    CSharp,
		interval.AugmentedOctave:  CDoubleSharp,

		interval.DiminishedNinth: DFlat,
		interval.MinorNinth:      DNatural,
		interval.MajorNinth:      DSharp,
		interval.AugmentedNinth:  DDoubleSharp,

		interval.DiminishedTenth: EFlat,
		interval.MinorTenth:      ENatural,
		interval.MajorTenth:      ESharp,
		interval.AugmentedTenth:  EDoubleSharp,

		interval.DiminishedEleventh: FNatural,
		interval.PerfectEleventh:    FSharp,
		interval.AugmentedEleventh:  FDoubleSharp,

		interval.DiminishedTwelfth: GNatural,
		interval.PerfectTwelfth:    GSharp,
		interval.AugmentedTwelfth:  GDoubleSharp,

		interval.DiminishedThirteenth: AFlat,
		interval.MinorThirteenth:      ANatural,
		interval.MajorThirteenth:      ASharp,
		interval.AugmentedThirteenth:  ADoubleSharp,

		interval.DiminishedFourteenth: BFlat,
		interval.MinorFourteenth:      BNatural,
		interval.MajorFourteenth:      BSharp,
		interval.AugmentedFourteenth:  BDoubleSharp,

		interval.DiminihsedFifteenth: CNatural,
		interval.PerfectFifteenth:    CSharp,
		interval.AugmentedFifteenth:  CDoubleSharp,
	},
	DFlat: {
		interval.PerfectUnison:   DFlat,
		interval.AugmentedUnison: DNatural,

		interval.DiminishedSecond: ETripleFlat,
		interval.MinorSecond:      EDoubleFlat,
		interval.MajorSecond:      EFlat,
		interval.AugmentedSecond:  ENatural,

		interval.DiminishedThird: FDoubleFlat,
		interval.MinorThird:      FFlat,
		interval.MajorThird:      FNatural,
		interval.AugmentedThird:  FSharp,

		interval.DiminishedFourth: GDoubleFlat,
		interval.PerfectFourth:    GFlat,
		interval.AugmentedFourth:  GNatural,

		interval.DiminishedFifth: ADoubleFlat,
		interval.PerfectFifth:    AFlat,
		interval.AugmentedFifth:  ANatural,

		interval.DiminishedSixth: BTripleFlat,
		interval.MinorSixth:      BDoubleFlat,
		interval.MajorSixth:      BFlat,
		interval.AugmentedSixth:  BNatural,

		interval.DiminishedSeventh: CDoubleFlat,
		interval.MinorSeventh:      CFlat,
		interval.MajorSeventh:      CNatural,
		interval.AugmentedSeventh:  CSharp,

		interval.DiminishedOctave: DDoubleFlat,
		interval.PerfectOctave:    DFlat,
		interval.AugmentedOctave:  DNatural,

		interval.DiminishedNinth: ETripleFlat,
		interval.MinorNinth:      EDoubleFlat,
		interval.MajorNinth:      EFlat,
		interval.AugmentedNinth:  ENatural,

		interval.DiminishedTenth: FDoubleFlat,
		interval.MinorTenth:      FFlat,
		interval.MajorTenth:      FNatural,
		interval.AugmentedTenth:  FSharp,

		interval.DiminishedEleventh: GDoubleFlat,
		interval.PerfectEleventh:    GFlat,
		interval.AugmentedEleventh:  GNatural,

		interval.DiminishedTwelfth: ADoubleFlat,
		interval.PerfectTwelfth:    AFlat,
		interval.AugmentedTwelfth:  ANatural,

		interval.DiminishedThirteenth: BTripleFlat,
		interval.MinorThirteenth:      BDoubleFlat,
		interval.MajorThirteenth:      BFlat,
		interval.AugmentedThirteenth:  BNatural,

		interval.DiminishedFourteenth: CDoubleFlat,
		interval.MinorFourteenth:      CFlat,
		interval.MajorFourteenth:      CNatural,
		interval.AugmentedFourteenth:  CSharp,

		interval.DiminihsedFifteenth: DDoubleFlat,
		interval.PerfectFifteenth:    DFlat,
		interval.AugmentedFifteenth:  DNatural,
	},
	DNatural: {
		interval.PerfectUnison:   DNatural,
		interval.AugmentedUnison: DSharp,

		interval.DiminishedSecond: EDoubleFlat,
		interval.MinorSecond:      EFlat,
		interval.MajorSecond:      ENatural,
		interval.AugmentedSecond:  ESharp,

		interval.DiminishedThird: FFlat,
		interval.MinorThird:      FNatural,
		interval.MajorThird:      FSharp,
		interval.AugmentedThird:  FDoubleSharp,

		interval.DiminishedFourth: GFlat,
		interval.PerfectFourth:    GNatural,
		interval.AugmentedFourth:  GSharp,

		interval.DiminishedFifth: AFlat,
		interval.PerfectFifth:    ANatural,
		interval.AugmentedFifth:  ASharp,

		interval.DiminishedSixth: BDoubleFlat,
		interval.MinorSixth:      BFlat,
		interval.MajorSixth:      BNatural,
		interval.AugmentedSixth:  BSharp,

		interval.DiminishedSeventh: CFlat,
		interval.MinorSeventh:      CNatural,
		interval.MajorSeventh:      CSharp,
		interval.AugmentedSeventh:  CDoubleSharp,

		interval.DiminishedOctave: DFlat,
		interval.PerfectOctave:    DNatural,
		interval.AugmentedOctave:  DSharp,

		interval.DiminishedNinth: EDoubleFlat,
		interval.MinorNinth:      EFlat,
		interval.MajorNinth:      ENatural,
		interval.AugmentedNinth:  ESharp,

		interval.DiminishedTenth: FFlat,
		interval.MinorTenth:      FNatural,
		interval.MajorTenth:      FSharp,
		interval.AugmentedTenth:  FDoubleSharp,

		interval.DiminishedEleventh: GFlat,
		interval.PerfectEleventh:    GNatural,
		interval.AugmentedEleventh:  GSharp,

		interval.DiminishedTwelfth: AFlat,
		interval.PerfectTwelfth:    ANatural,
		interval.AugmentedTwelfth:  ASharp,

		interval.DiminishedThirteenth: BDoubleFlat,
		interval.MinorThirteenth:      BFlat,
		interval.MajorThirteenth:      BNatural,
		interval.AugmentedThirteenth:  BSharp,

		interval.DiminishedFourteenth: CFlat,
		interval.MinorFourteenth:      CNatural,
		interval.MajorFourteenth:      CSharp,
		interval.AugmentedFourteenth:  CDoubleSharp,

		interval.DiminihsedFifteenth: DFlat,
		interval.PerfectFifteenth:    DNatural,
		interval.AugmentedFifteenth:  DSharp,
	},
	DSharp: {
		interval.PerfectUnison:   DSharp,
		interval.AugmentedUnison: DDoubleSharp,

		interval.DiminishedSecond: EFlat,
		interval.MinorSecond:      ENatural,
		interval.MajorSecond:      ESharp,
		interval.AugmentedSecond:  EDoubleSharp,

		interval.DiminishedThird: FNatural,
		interval.MinorThird:      FSharp,
		interval.MajorThird:      FDoubleSharp,
		interval.AugmentedThird:  FTripleSharp,

		interval.DiminishedFourth: GNatural,
		interval.PerfectFourth:    GSharp,
		interval.AugmentedFourth:  GDoubleSharp,

		interval.DiminishedFifth: ANatural,
		interval.PerfectFifth:    ASharp,
		interval.AugmentedFifth:  ADoubleSharp,

		interval.DiminishedSixth: BFlat,
		interval.MinorSixth:      BNatural,
		interval.MajorSixth:      BSharp,
		interval.AugmentedSixth:  BDoubleSharp,

		interval.DiminishedSeventh: CNatural,
		interval.MinorSeventh:      CSharp,
		interval.MajorSeventh:      CDoubleSharp,
		interval.AugmentedSeventh:  CTripleSharp,

		interval.DiminishedOctave: DNatural,
		interval.PerfectOctave:    DSharp,
		interval.AugmentedOctave:  DDoubleSharp,

		interval.DiminishedNinth: EFlat,
		interval.MinorNinth:      ENatural,
		interval.MajorNinth:      ESharp,
		interval.AugmentedNinth:  EDoubleSharp,

		interval.DiminishedTenth: FNatural,
		interval.MinorTenth:      FSharp,
		interval.MajorTenth:      FDoubleSharp,
		interval.AugmentedTenth:  FTripleSharp,

		interval.DiminishedEleventh: GNatural,
		interval.PerfectEleventh:    GSharp,
		interval.AugmentedEleventh:  GDoubleSharp,

		interval.DiminishedTwelfth: ANatural,
		interval.PerfectTwelfth:    ASharp,
		interval.AugmentedTwelfth:  ADoubleSharp,

		interval.DiminishedThirteenth: BFlat,
		interval.MinorThirteenth:      BNatural,
		interval.MajorThirteenth:      BSharp,
		interval.AugmentedThirteenth:  BDoubleSharp,

		interval.DiminishedFourteenth: CNatural,
		interval.MinorFourteenth:      CSharp,
		interval.MajorFourteenth:      CDoubleSharp,
		interval.AugmentedFourteenth:  CTripleSharp,

		interval.DiminihsedFifteenth: DNatural,
		interval.PerfectFifteenth:    DSharp,
		interval.AugmentedFifteenth:  DDoubleSharp,
	},
	EFlat: {
		interval.PerfectUnison:   EFlat,
		interval.AugmentedUnison: ENatural,

		interval.DiminishedSecond: FDoubleFlat,
		interval.MinorSecond:      FFlat,
		interval.MajorSecond:      FNatural,
		interval.AugmentedSecond:  FSharp,

		interval.DiminishedThird: GDoubleFlat,
		interval.MinorThird:      GFlat,
		interval.MajorThird:      GNatural,
		interval.AugmentedThird:  GSharp,

		interval.DiminishedFourth: ADoubleFlat,
		interval.PerfectFourth:    AFlat,
		interval.AugmentedFourth:  ANatural,

		interval.DiminishedFifth: BDoubleFlat,
		interval.PerfectFifth:    BFlat,
		interval.AugmentedFifth:  BNatural,

		interval.DiminishedSixth: CDoubleFlat,
		interval.MinorSixth:      CFlat,
		interval.MajorSixth:      CNatural,
		interval.AugmentedSixth:  CSharp,

		interval.DiminishedSeventh: DDoubleFlat,
		interval.MinorSeventh:      DFlat,
		interval.MajorSeventh:      DNatural,
		interval.AugmentedSeventh:  DSharp,

		interval.DiminishedOctave: EDoubleFlat,
		interval.PerfectOctave:    EFlat,
		interval.AugmentedOctave:  ENatural,

		interval.DiminishedNinth: FDoubleFlat,
		interval.MinorNinth:      FFlat,
		interval.MajorNinth:      FNatural,
		interval.AugmentedNinth:  FSharp,

		interval.DiminishedTenth: GDoubleFlat,
		interval.MinorTenth:      GFlat,
		interval.MajorTenth:      GNatural,
		interval.AugmentedTenth:  GSharp,

		interval.DiminishedEleventh: ADoubleFlat,
		interval.PerfectEleventh:    AFlat,
		interval.AugmentedEleventh:  ANatural,

		interval.DiminishedTwelfth: BDoubleFlat,
		interval.PerfectTwelfth:    BFlat,
		interval.AugmentedTwelfth:  BNatural,

		interval.DiminishedThirteenth: CDoubleFlat,
		interval.MinorThirteenth:      CFlat,
		interval.MajorThirteenth:      CNatural,
		interval.AugmentedThirteenth:  CSharp,

		interval.DiminishedFourteenth: DDoubleFlat,
		interval.MinorFourteenth:      DFlat,
		interval.MajorFourteenth:      DNatural,
		interval.AugmentedFourteenth:  DSharp,

		interval.DiminihsedFifteenth: EDoubleFlat,
		interval.PerfectFifteenth:    EFlat,
		interval.AugmentedFifteenth:  ENatural,
	},
	ENatural: {
		interval.PerfectUnison:   ENatural,
		interval.AugmentedUnison: ESharp,

		interval.DiminishedSecond: FFlat,
		interval.MinorSecond:      FNatural,
		interval.MajorSecond:      FSharp,
		interval.AugmentedSecond:  FDoubleSharp,

		interval.DiminishedThird: GFlat,
		interval.MinorThird:      GNatural,
		interval.MajorThird:      GSharp,
		interval.AugmentedThird:  GDoubleSharp,

		interval.DiminishedFourth: AFlat,
		interval.PerfectFourth:    ANatural,
		interval.AugmentedFourth:  ASharp,

		interval.DiminishedFifth: BFlat,
		interval.PerfectFifth:    BNatural,
		interval.AugmentedFifth:  BSharp,

		interval.DiminishedSixth: CFlat,
		interval.MinorSixth:      CNatural,
		interval.MajorSixth:      CSharp,
		interval.AugmentedSixth:  CDoubleSharp,

		interval.DiminishedSeventh: DFlat,
		interval.MinorSeventh:      DNatural,
		interval.MajorSeventh:      DSharp,
		interval.AugmentedSeventh:  DDoubleSharp,

		interval.DiminishedOctave: EFlat,
		interval.PerfectOctave:    ENatural,
		interval.AugmentedOctave:  ESharp,

		interval.DiminishedNinth: FFlat,
		interval.MinorNinth:      FNatural,
		interval.MajorNinth:      FSharp,
		interval.AugmentedNinth:  FDoubleSharp,

		interval.DiminishedTenth: GFlat,
		interval.MinorTenth:      GNatural,
		interval.MajorTenth:      GSharp,
		interval.AugmentedTenth:  GDoubleSharp,

		interval.DiminishedEleventh: AFlat,
		interval.PerfectEleventh:    ANatural,
		interval.AugmentedEleventh:  ASharp,

		interval.DiminishedTwelfth: BFlat,
		interval.PerfectTwelfth:    BNatural,
		interval.AugmentedTwelfth:  BSharp,

		interval.DiminishedThirteenth: CFlat,
		interval.MinorThirteenth:      CNatural,
		interval.MajorThirteenth:      CSharp,
		interval.AugmentedThirteenth:  CDoubleSharp,

		interval.DiminishedFourteenth: DFlat,
		interval.MinorFourteenth:      DNatural,
		interval.MajorFourteenth:      DSharp,
		interval.AugmentedFourteenth:  DDoubleSharp,

		interval.DiminihsedFifteenth: EFlat,
		interval.PerfectFifteenth:    ENatural,
		interval.AugmentedFifteenth:  ESharp,
	},
	FNatural: {
		interval.PerfectUnison:   FNatural,
		interval.AugmentedUnison: FSharp,

		interval.DiminishedSecond: GDoubleFlat,
		interval.MinorSecond:      GFlat,
		interval.MajorSecond:      GNatural,
		interval.AugmentedSecond:  GSharp,

		interval.DiminishedThird: ADoubleFlat,
		interval.MinorThird:      AFlat,
		interval.MajorThird:      ANatural,
		interval.AugmentedThird:  ASharp,

		interval.DiminishedFourth: BDoubleFlat,
		interval.PerfectFourth:    BFlat,
		interval.AugmentedFourth:  BNatural,

		interval.DiminishedFifth: CFlat,
		interval.PerfectFifth:    CNatural,
		interval.AugmentedFifth:  CSharp,

		interval.DiminishedSixth: DDoubleFlat,
		interval.MinorSixth:      DFlat,
		interval.MajorSixth:      DNatural,
		interval.AugmentedSixth:  DSharp,

		interval.DiminishedSeventh: EDoubleFlat,
		interval.MinorSeventh:      EFlat,
		interval.MajorSeventh:      ENatural,
		interval.AugmentedSeventh:  ESharp,

		interval.DiminishedOctave: FFlat,
		interval.PerfectOctave:    FNatural,
		interval.AugmentedOctave:  FSharp,

		interval.DiminishedNinth: GDoubleFlat,
		interval.MinorNinth:      GFlat,
		interval.MajorNinth:      GNatural,
		interval.AugmentedNinth:  GSharp,

		interval.DiminishedTenth: ADoubleFlat,
		interval.MinorTenth:      AFlat,
		interval.MajorTenth:      ANatural,
		interval.AugmentedTenth:  ASharp,

		interval.DiminishedEleventh: BDoubleFlat,
		interval.PerfectEleventh:    BFlat,
		interval.AugmentedEleventh:  BNatural,

		interval.DiminishedTwelfth: CFlat,
		interval.PerfectTwelfth:    CNatural,
		interval.AugmentedTwelfth:  CSharp,

		interval.DiminishedThirteenth: DDoubleFlat,
		interval.MinorThirteenth:      DFlat,
		interval.MajorThirteenth:      DNatural,
		interval.AugmentedThirteenth:  DSharp,

		interval.DiminishedFourteenth: EDoubleFlat,
		interval.MinorFourteenth:      EFlat,
		interval.MajorFourteenth:      ENatural,
		interval.AugmentedFourteenth:  ESharp,

		interval.DiminihsedFifteenth: FFlat,
		interval.PerfectFifteenth:    FNatural,
		interval.AugmentedFifteenth:  FSharp,
	},
	FSharp: {
		interval.PerfectUnison:   FSharp,
		interval.AugmentedUnison: FDoubleSharp,

		interval.DiminishedSecond: GFlat,
		interval.MinorSecond:      GNatural,
		interval.MajorSecond:      GSharp,
		interval.AugmentedSecond:  GDoubleSharp,

		interval.DiminishedThird: AFlat,
		interval.MinorThird:      ANatural,
		interval.MajorThird:      ASharp,
		interval.AugmentedThird:  ADoubleSharp,

		interval.DiminishedFourth: BFlat,
		interval.PerfectFourth:    BNatural,
		interval.AugmentedFourth:  BSharp,

		interval.DiminishedFifth: CNatural,
		interval.PerfectFifth:    CSharp,
		interval.AugmentedFifth:  CDoubleSharp,

		interval.DiminishedSixth: DFlat,
		interval.MinorSixth:      DNatural,
		interval.MajorSixth:      DSharp,
		interval.AugmentedSixth:  DDoubleSharp,

		interval.DiminishedSeventh: EFlat,
		interval.MinorSeventh:      ENatural,
		interval.MajorSeventh:      ESharp,
		interval.AugmentedSeventh:  EDoubleSharp,

		interval.DiminishedOctave: FNatural,
		interval.PerfectOctave:    FSharp,

		interval.DiminishedNinth: GFlat,
		interval.MinorNinth:      GNatural,
		interval.MajorNinth:      GSharp,
		interval.AugmentedNinth:  GDoubleSharp,

		interval.DiminishedTenth: AFlat,
		interval.MinorTenth:      ANatural,
		interval.MajorTenth:      ASharp,
		interval.AugmentedTenth:  ADoubleSharp,

		interval.DiminishedEleventh: BFlat,
		interval.PerfectEleventh:    BNatural,
		interval.AugmentedEleventh:  BSharp,

		interval.DiminishedTwelfth: CNatural,
		interval.PerfectTwelfth:    CSharp,
		interval.AugmentedTwelfth:  CDoubleSharp,

		interval.DiminishedThirteenth: DFlat,
		interval.MinorThirteenth:      DNatural,
		interval.MajorThirteenth:      DSharp,
		interval.AugmentedThirteenth:  DDoubleSharp,

		interval.DiminishedFourteenth: EFlat,
		interval.MinorFourteenth:      ENatural,
		interval.MajorFourteenth:      ESharp,
		interval.AugmentedFourteenth:  EDoubleSharp,

		interval.DiminihsedFifteenth: FNatural,
		interval.PerfectFifteenth:    FSharp,
		interval.AugmentedFifteenth:  FDoubleSharp,
	},
	GFlat: {
		interval.PerfectUnison:   GFlat,
		interval.AugmentedUnison: GNatural,

		interval.DiminishedSecond: ATripleFlat,
		interval.MinorSecond:      ADoubleFlat,
		interval.MajorSecond:      AFlat,
		interval.AugmentedSecond:  ANatural,

		interval.DiminishedThird: BTripleFlat,
		interval.MinorThird:      BDoubleFlat,
		interval.MajorThird:      BFlat,
		interval.AugmentedThird:  BNatural,

		interval.DiminishedFourth: CDoubleFlat,
		interval.PerfectFourth:    CFlat,
		interval.AugmentedFourth:  CNatural,

		interval.DiminishedFifth: DDoubleFlat,
		interval.PerfectFifth:    DFlat,
		interval.AugmentedFifth:  DNatural,

		interval.DiminishedSixth: ETripleFlat,
		interval.MinorSixth:      EDoubleFlat,
		interval.MajorSixth:      EFlat,
		interval.AugmentedSixth:  ENatural,

		interval.DiminishedSeventh: FDoubleFlat,
		interval.MinorSeventh:      FFlat,
		interval.MajorSeventh:      FNatural,
		interval.AugmentedSeventh:  FSharp,

		interval.DiminishedOctave: GDoubleFlat,
		interval.PerfectOctave:    GFlat,
		interval.AugmentedOctave:  GNatural,

		interval.DiminishedNinth: ATripleFlat,
		interval.MinorNinth:      ADoubleFlat,
		interval.MajorNinth:      AFlat,
		interval.AugmentedNinth:  ANatural,

		interval.DiminishedTenth: BTripleFlat,
		interval.MinorTenth:      BDoubleFlat,
		interval.MajorTenth:      BFlat,
		interval.AugmentedTenth:  BNatural,

		interval.DiminishedEleventh: CDoubleFlat,
		interval.PerfectEleventh:    CFlat,
		interval.AugmentedEleventh:  CNatural,

		interval.DiminishedTwelfth: DDoubleFlat,
		interval.PerfectTwelfth:    DFlat,
		interval.AugmentedTwelfth:  DNatural,

		interval.DiminishedThirteenth: ETripleFlat,
		interval.MinorThirteenth:      EDoubleFlat,
		interval.MajorThirteenth:      EFlat,
		interval.AugmentedThirteenth:  ENatural,

		interval.DiminishedFourteenth: EDoubleFlat,
		interval.MinorFourteenth:      EFlat,
		interval.MajorFourteenth:      FNatural,
		interval.AugmentedFourteenth:  FSharp,

		interval.DiminihsedFifteenth: GDoubleFlat,
		interval.PerfectFifteenth:    GFlat,
		interval.AugmentedFifteenth:  GNatural,
	},
	GNatural: {
		interval.PerfectUnison:   GNatural,
		interval.AugmentedUnison: GSharp,

		interval.DiminishedSecond: ADoubleFlat,
		interval.MinorSecond:      AFlat,
		interval.MajorSecond:      ANatural,
		interval.AugmentedSecond:  ASharp,

		interval.DiminishedThird: BDoubleFlat,
		interval.MinorThird:      BFlat,
		interval.MajorThird:      BNatural,
		interval.AugmentedThird:  BSharp,

		interval.DiminishedFourth: CFlat,
		interval.PerfectFourth:    CNatural,
		interval.AugmentedFourth:  CSharp,

		interval.DiminishedFifth: DFlat,
		interval.PerfectFifth:    DNatural,
		interval.AugmentedFifth:  DSharp,

		interval.DiminishedSixth: EDoubleFlat,
		interval.MinorSixth:      EFlat,
		interval.MajorSixth:      ENatural,
		interval.AugmentedSixth:  ESharp,

		interval.DiminishedSeventh: FFlat,
		interval.MinorSeventh:      FNatural,
		interval.MajorSeventh:      FSharp,
		interval.AugmentedSeventh:  FDoubleSharp,

		interval.DiminishedOctave: GFlat,
		interval.PerfectOctave:    GNatural,
		interval.AugmentedOctave:  GSharp,

		interval.DiminishedNinth: ADoubleFlat,
		interval.MinorNinth:      AFlat,
		interval.MajorNinth:      ANatural,
		interval.AugmentedNinth:  ASharp,

		interval.DiminishedTenth: BDoubleFlat,
		interval.MinorTenth:      BFlat,
		interval.MajorTenth:      BNatural,
		interval.AugmentedTenth:  BSharp,

		interval.DiminishedEleventh: CFlat,
		interval.PerfectEleventh:    CNatural,
		interval.AugmentedEleventh:  CSharp,

		interval.DiminishedTwelfth: DFlat,
		interval.PerfectTwelfth:    DNatural,
		interval.AugmentedTwelfth:  DSharp,

		interval.DiminishedThirteenth: EDoubleFlat,
		interval.MinorThirteenth:      EFlat,
		interval.MajorThirteenth:      ENatural,
		interval.AugmentedThirteenth:  ESharp,

		interval.DiminishedFourteenth: EFlat,
		interval.MinorFourteenth:      ENatural,
		interval.MajorFourteenth:      FSharp,
		interval.AugmentedFourteenth:  FDoubleSharp,

		interval.DiminihsedFifteenth: GFlat,
		interval.PerfectFifteenth:    GNatural,
		interval.AugmentedFifteenth:  GSharp,
	},
	GSharp: {
		interval.PerfectUnison:   GSharp,
		interval.AugmentedUnison: GDoubleSharp,

		interval.DiminishedSecond: AFlat,
		interval.MinorSecond:      ANatural,
		interval.MajorSecond:      ASharp,
		interval.AugmentedSecond:  ADoubleSharp,

		interval.DiminishedThird: BFlat,
		interval.MinorThird:      BNatural,
		interval.MajorThird:      BSharp,
		interval.AugmentedThird:  BDoubleSharp,

		interval.DiminishedFourth: CNatural,
		interval.PerfectFourth:    CSharp,
		interval.AugmentedFourth:  CDoubleSharp,

		interval.DiminishedFifth: DNatural,
		interval.PerfectFifth:    DSharp,
		interval.AugmentedFifth:  DDoubleSharp,

		interval.DiminishedSixth: EFlat,
		interval.MinorSixth:      ENatural,
		interval.MajorSixth:      ESharp,
		interval.AugmentedSixth:  EDoubleSharp,

		interval.DiminishedSeventh: FNatural,
		interval.MinorSeventh:      FSharp,
		interval.MajorSeventh:      FDoubleSharp,
		interval.AugmentedSeventh:  FTripleSharp,

		interval.DiminishedOctave: GNatural,
		interval.PerfectOctave:    GSharp,
		interval.AugmentedOctave:  GDoubleSharp,

		interval.DiminishedNinth: AFlat,
		interval.MinorNinth:      ANatural,
		interval.MajorNinth:      ASharp,
		interval.AugmentedNinth:  ADoubleSharp,

		interval.DiminishedTenth: BFlat,
		interval.MinorTenth:      BNatural,
		interval.MajorTenth:      BSharp,
		interval.AugmentedTenth:  BDoubleSharp,

		interval.DiminishedEleventh: CNatural,
		interval.PerfectEleventh:    CSharp,
		interval.AugmentedEleventh:  CDoubleSharp,

		interval.DiminishedTwelfth: DNatural,
		interval.PerfectTwelfth:    DSharp,
		interval.AugmentedTwelfth:  DDoubleSharp,

		interval.DiminishedThirteenth: EFlat,
		interval.MinorThirteenth:      ENatural,
		interval.MajorThirteenth:      ESharp,
		interval.AugmentedThirteenth:  EDoubleSharp,

		interval.DiminishedFourteenth: ENatural,
		interval.MinorFourteenth:      ESharp,
		interval.MajorFourteenth:      FDoubleSharp,
		interval.AugmentedFourteenth:  FTripleSharp,

		interval.DiminihsedFifteenth: GNatural,
		interval.PerfectFifteenth:    GSharp,
		interval.AugmentedFifteenth:  GDoubleSharp,
	},
	AFlat: {
		interval.PerfectUnison:   AFlat,
		interval.AugmentedUnison: ANatural,

		interval.DiminishedSecond: BTripleFlat,
		interval.MinorSecond:      BDoubleFlat,
		interval.MajorSecond:      BFlat,
		interval.AugmentedSecond:  BNatural,

		interval.DiminishedThird: CDoubleFlat,
		interval.MinorThird:      CFlat,
		interval.MajorThird:      CNatural,
		interval.AugmentedThird:  CSharp,

		interval.DiminishedFourth: DDoubleFlat,
		interval.PerfectFourth:    DFlat,
		interval.AugmentedFourth:  DNatural,

		interval.DiminishedFifth: EDoubleFlat,
		interval.PerfectFifth:    EFlat,
		interval.AugmentedFifth:  ENatural,

		interval.DiminishedSixth: FDoubleFlat,
		interval.MinorSixth:      FFlat,
		interval.MajorSixth:      FNatural,
		interval.AugmentedSixth:  FSharp,

		interval.DiminishedSeventh: GDoubleFlat,
		interval.MinorSeventh:      GFlat,
		interval.MajorSeventh:      GNatural,
		interval.AugmentedSeventh:  GSharp,

		interval.DiminishedOctave: ADoubleFlat,
		interval.PerfectOctave:    AFlat,
		interval.AugmentedOctave:  ANatural,

		interval.DiminishedNinth: BTripleFlat,
		interval.MinorNinth:      BDoubleFlat,
		interval.MajorNinth:      BFlat,
		interval.AugmentedNinth:  BNatural,

		interval.DiminishedTenth: CDoubleFlat,
		interval.MinorTenth:      CFlat,
		interval.MajorTenth:      CNatural,
		interval.AugmentedTenth:  CSharp,

		interval.DiminishedEleventh: DDoubleFlat,
		interval.PerfectEleventh:    DFlat,
		interval.AugmentedEleventh:  DNatural,

		interval.DiminishedTwelfth: EDoubleFlat,
		interval.PerfectTwelfth:    EFlat,
		interval.AugmentedTwelfth:  ENatural,

		interval.DiminishedThirteenth: FDoubleFlat,
		interval.MinorThirteenth:      FFlat,
		interval.MajorThirteenth:      FNatural,
		interval.AugmentedThirteenth:  FSharp,

		interval.DiminishedFourteenth: GDoubleFlat,
		interval.MinorFourteenth:      GFlat,
		interval.MajorFourteenth:      GNatural,
		interval.AugmentedFourteenth:  GSharp,

		interval.DiminihsedFifteenth: ADoubleFlat,
		interval.PerfectFifteenth:    AFlat,
		interval.AugmentedFifteenth:  ANatural,
	},
	ANatural: {
		interval.PerfectUnison:   ANatural,
		interval.AugmentedUnison: ASharp,

		interval.DiminishedSecond: BDoubleFlat,
		interval.MinorSecond:      BFlat,
		interval.MajorSecond:      BNatural,
		interval.AugmentedSecond:  BSharp,

		interval.DiminishedThird: CFlat,
		interval.MinorThird:      CNatural,
		interval.MajorThird:      CSharp,
		interval.AugmentedThird:  CDoubleSharp,

		interval.DiminishedFourth: DFlat,
		interval.PerfectFourth:    DNatural,
		interval.AugmentedFourth:  DSharp,

		interval.DiminishedFifth: EFlat,
		interval.PerfectFifth:    ENatural,
		interval.AugmentedFifth:  ESharp,

		interval.DiminishedSixth: FFlat,
		interval.MinorSixth:      FNatural,
		interval.MajorSixth:      FSharp,
		interval.AugmentedSixth:  FDoubleSharp,

		interval.DiminishedSeventh: GFlat,
		interval.MinorSeventh:      GNatural,
		interval.MajorSeventh:      GSharp,
		interval.AugmentedSeventh:  GDoubleSharp,

		interval.DiminishedOctave: AFlat,
		interval.PerfectOctave:    ANatural,
		interval.AugmentedOctave:  ASharp,

		interval.DiminishedNinth: BDoubleFlat,
		interval.MinorNinth:      BFlat,
		interval.MajorNinth:      BNatural,
		interval.AugmentedNinth:  BSharp,

		interval.DiminishedTenth: CFlat,
		interval.MinorTenth:      CNatural,
		interval.MajorTenth:      CSharp,
		interval.AugmentedTenth:  CDoubleSharp,

		interval.DiminishedEleventh: DFlat,
		interval.PerfectEleventh:    DNatural,
		interval.AugmentedEleventh:  DSharp,

		interval.DiminishedTwelfth: EFlat,
		interval.PerfectTwelfth:    ENatural,
		interval.AugmentedTwelfth:  ESharp,

		interval.DiminishedThirteenth: FFlat,
		interval.MinorThirteenth:      FNatural,
		interval.MajorThirteenth:      FSharp,
		interval.AugmentedThirteenth:  FDoubleSharp,

		interval.DiminishedFourteenth: GFlat,
		interval.MinorFourteenth:      GNatural,
		interval.MajorFourteenth:      GSharp,
		interval.AugmentedFourteenth:  GDoubleSharp,

		interval.DiminihsedFifteenth: AFlat,
		interval.PerfectFifteenth:    ANatural,
		interval.AugmentedFifteenth:  ASharp,
	},
	ASharp: {
		interval.PerfectUnison:   ASharp,
		interval.AugmentedUnison: ADoubleSharp,

		interval.DiminishedSecond: BFlat,
		interval.MinorSecond:      BNatural,
		interval.MajorSecond:      BSharp,
		interval.AugmentedSecond:  BDoubleSharp,

		interval.DiminishedThird: CNatural,
		interval.MinorThird:      CSharp,
		interval.MajorThird:      CDoubleSharp,
		interval.AugmentedThird:  CTripleSharp,

		interval.DiminishedFourth: DNatural,
		interval.PerfectFourth:    DSharp,
		interval.AugmentedFourth:  DDoubleSharp,

		interval.DiminishedFifth: ENatural,
		interval.PerfectFifth:    ESharp,
		interval.AugmentedFifth:  EDoubleSharp,

		interval.DiminishedSixth: FNatural,
		interval.MinorSixth:      FSharp,
		interval.MajorSixth:      FDoubleSharp,
		interval.AugmentedSixth:  FTripleSharp,

		interval.DiminishedSeventh: GNatural,
		interval.MinorSeventh:      GSharp,
		interval.MajorSeventh:      GDoubleSharp,
		interval.AugmentedSeventh:  GTripleSharp,

		interval.DiminishedOctave: ANatural,
		interval.PerfectOctave:    ASharp,
		interval.AugmentedOctave:  ADoubleSharp,

		interval.DiminishedNinth: BFlat,
		interval.MinorNinth:      BNatural,
		interval.MajorNinth:      BSharp,
		interval.AugmentedNinth:  BDoubleSharp,

		interval.DiminishedTenth: CNatural,
		interval.MinorTenth:      CSharp,
		interval.MajorTenth:      CDoubleSharp,
		interval.AugmentedTenth:  CTripleSharp,

		interval.DiminishedEleventh: DNatural,
		interval.PerfectEleventh:    DSharp,
		interval.AugmentedEleventh:  DDoubleSharp,

		interval.DiminishedTwelfth: ENatural,
		interval.PerfectTwelfth:    ESharp,
		interval.AugmentedTwelfth:  EDoubleSharp,

		interval.DiminishedThirteenth: FNatural,
		interval.MinorThirteenth:      FSharp,
		interval.MajorThirteenth:      FDoubleSharp,
		interval.AugmentedThirteenth:  FTripleSharp,

		interval.DiminishedFourteenth: GNatural,
		interval.MinorFourteenth:      GSharp,
		interval.MajorFourteenth:      GDoubleSharp,
		interval.AugmentedFourteenth:  GTripleSharp,

		interval.DiminihsedFifteenth: ANatural,
		interval.PerfectFifteenth:    ASharp,
		interval.AugmentedFifteenth:  ADoubleSharp,
	},
	BFlat: {
		interval.PerfectUnison:   BFlat,
		interval.AugmentedUnison: BNatural,

		interval.DiminishedSecond: CDoubleFlat,
		interval.MinorSecond:      CFlat,
		interval.MajorSecond:      CNatural,
		interval.AugmentedSecond:  CSharp,

		interval.DiminishedThird: DDoubleFlat,
		interval.MinorThird:      DFlat,
		interval.MajorThird:      DNatural,
		interval.AugmentedThird:  DSharp,

		interval.DiminishedFourth: EDoubleFlat,
		interval.PerfectFourth:    EFlat,
		interval.AugmentedFourth:  ENatural,

		interval.DiminishedFifth: FFlat,
		interval.PerfectFifth:    FNatural,
		interval.AugmentedFifth:  FSharp,

		interval.DiminishedSixth: GDoubleFlat,
		interval.MinorSixth:      GFlat,
		interval.MajorSixth:      GNatural,
		interval.AugmentedSixth:  GSharp,

		interval.DiminishedSeventh: ADoubleFlat,
		interval.MinorSeventh:      AFlat,
		interval.MajorSeventh:      ANatural,
		interval.AugmentedSeventh:  ASharp,

		interval.DiminishedOctave: BDoubleFlat,
		interval.PerfectOctave:    BFlat,
		interval.AugmentedOctave:  BNatural,

		interval.DiminishedNinth: CDoubleFlat,
		interval.MinorNinth:      CFlat,
		interval.MajorNinth:      CNatural,
		interval.AugmentedNinth:  CSharp,

		interval.DiminishedTenth: DDoubleFlat,
		interval.MinorTenth:      DFlat,
		interval.MajorTenth:      DNatural,
		interval.AugmentedTenth:  DSharp,

		interval.DiminishedEleventh: EDoubleFlat,
		interval.PerfectEleventh:    EFlat,
		interval.AugmentedEleventh:  ENatural,

		interval.DiminishedTwelfth: FFlat,
		interval.PerfectTwelfth:    FNatural,
		interval.AugmentedTwelfth:  FSharp,

		interval.DiminishedThirteenth: GDoubleFlat,
		interval.MinorThirteenth:      GFlat,
		interval.MajorThirteenth:      GNatural,
		interval.AugmentedThirteenth:  GSharp,

		interval.DiminishedFourteenth: ADoubleFlat,
		interval.MinorFourteenth:      AFlat,
		interval.MajorFourteenth:      ANatural,
		interval.AugmentedFourteenth:  ASharp,

		interval.DiminihsedFifteenth: BDoubleFlat,
		interval.PerfectFifteenth:    BFlat,
		interval.AugmentedFifteenth:  BNatural,
	},
	BNatural: {
		interval.PerfectUnison:   BNatural,
		interval.AugmentedUnison: BSharp,

		interval.DiminishedSecond: CFlat,
		interval.MinorSecond:      CNatural,
		interval.MajorSecond:      CSharp,
		interval.AugmentedSecond:  CDoubleSharp,

		interval.DiminishedThird: DFlat,
		interval.MinorThird:      DNatural,
		interval.MajorThird:      DSharp,
		interval.AugmentedThird:  DDoubleSharp,

		interval.DiminishedFourth: EFlat,
		interval.PerfectFourth:    ENatural,
		interval.AugmentedFourth:  ESharp,

		interval.DiminishedFifth: FNatural,
		interval.PerfectFifth:    FSharp,
		interval.AugmentedFifth:  FDoubleSharp,

		interval.DiminishedSixth: GFlat,
		interval.MinorSixth:      GNatural,
		interval.MajorSixth:      GSharp,
		interval.AugmentedSixth:  GDoubleSharp,

		interval.DiminishedSeventh: AFlat,
		interval.MinorSeventh:      ANatural,
		interval.MajorSeventh:      ASharp,
		interval.AugmentedSeventh:  ADoubleSharp,

		interval.DiminishedOctave: BFlat,
		interval.PerfectOctave:    BNatural,
		interval.AugmentedOctave:  BSharp,

		interval.DiminishedNinth: CFlat,
		interval.MinorNinth:      CNatural,
		interval.MajorNinth:      CSharp,
		interval.AugmentedNinth:  CDoubleSharp,

		interval.DiminishedTenth: DFlat,
		interval.MinorTenth:      DNatural,
		interval.MajorTenth:      DSharp,
		interval.AugmentedTenth:  DDoubleSharp,

		interval.DiminishedEleventh: EFlat,
		interval.PerfectEleventh:    ENatural,
		interval.AugmentedEleventh:  ESharp,

		interval.DiminishedTwelfth: FNatural,
		interval.PerfectTwelfth:    FSharp,
		interval.AugmentedTwelfth:  FDoubleSharp,

		interval.DiminishedThirteenth: GFlat,
		interval.MinorThirteenth:      GNatural,
		interval.MajorThirteenth:      GSharp,
		interval.AugmentedThirteenth:  GDoubleSharp,

		interval.DiminishedFourteenth: AFlat,
		interval.MinorFourteenth:      ANatural,
		interval.MajorFourteenth:      ASharp,
		interval.AugmentedFourteenth:  ADoubleSharp,

		interval.DiminihsedFifteenth: BFlat,
		interval.PerfectFifteenth:    BNatural,
		interval.AugmentedFifteenth:  BSharp,
	},
}

// Raise raise note by given interval
func (n Note) Raise(interval interval.Interval) Note {
	if v, found := raisedNotes[n][interval]; found {
		return v
	}

	panic(fmt.Errorf("cannot raise %s by %s", n, interval))
}
