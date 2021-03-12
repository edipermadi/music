package chordtype

import (
	"fmt"

	"github.com/edipermadi/music/pkg/theory/interval"
)

var typeIntervals = map[Type][]interval.Interval{
	// Dyad
	PowerChord: {interval.PerfectUnison, interval.PerfectFifth},

	// Triad
	Diminished:                            {interval.PerfectUnison, interval.MinorThird, interval.DiminishedFifth},
	MinorDoubleFlatFifth:                  {interval.PerfectUnison, interval.MinorThird, interval.PerfectFourth},
	Minor:                                 {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth},
	MinorAddNinth:                         {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorNinth},
	MinorAddEleventh:                      {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.PerfectEleventh},
	MinorAddSharpNinth:                    {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.AugmentedNinth},
	MinorAddFourth:                        {interval.PerfectUnison, interval.MinorThird, interval.PerfectFourth, interval.PerfectFifth},
	MinorAddSharpFourth:                   {interval.PerfectUnison, interval.MinorThird, interval.AugmentedFourth, interval.PerfectFifth},
	MinorSharpFifth:                       {interval.PerfectUnison, interval.MinorThird, interval.MinorSixth},
	MajorFlatFifth:                        {interval.PerfectUnison, interval.MajorThird, interval.DiminishedFifth},
	Major:                                 {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth},
	MajorAddNinth:                         {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorNinth},
	MajorAddEleventh:                      {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.PerfectEleventh},
	MajorAddSharpNinth:                    {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.AugmentedNinth},
	MajorDoubleSharpFifth:                 {interval.PerfectUnison, interval.MajorThird, interval.MajorSixth},
	MajorAddFourth:                        {interval.PerfectUnison, interval.MajorThird, interval.PerfectFourth, interval.PerfectFifth},
	MajorAddSharpFourth:                   {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFourth, interval.PerfectFifth},
	Augmented:                             {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFifth},
	SuspendedSecondDoubleFlatFifth:        {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFourth},
	SuspendedSecondFlatFifth:              {interval.PerfectUnison, interval.MajorSecond, interval.DiminishedFifth},
	SuspendedSecondFlatFifthAddSharpFifth: {interval.PerfectUnison, interval.MajorSecond, interval.DiminishedFifth, interval.AugmentedFifth},
	SuspendedSecond:                       {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth},
	SuspendedSecondSharpFifth:             {interval.PerfectUnison, interval.MajorSecond, interval.AugmentedFifth},
	SuspendedFourthFlatFifth:              {interval.PerfectUnison, interval.PerfectFourth, interval.DiminishedFifth},
	SuspendedFourth:                       {interval.PerfectUnison, interval.PerfectFourth, interval.PerfectFifth},
	SuspendedFourthSharpFifth:             {interval.PerfectUnison, interval.PerfectFourth, interval.AugmentedFifth},
	SuspendedFourthDoubleSharpFifth:       {interval.PerfectUnison, interval.PerfectFourth, interval.MajorSixth},
	Phrygian:                              {interval.PerfectUnison, interval.MinorSecond, interval.PerfectFifth},
	PhrygianAddSeventh:                    {interval.PerfectUnison, interval.MinorSecond, interval.PerfectFifth, interval.MajorSeventh},
	Lydian:                                {interval.PerfectUnison, interval.AugmentedFourth, interval.PerfectFifth},
	LydianMajorSeventh:                    {interval.PerfectUnison, interval.AugmentedFourth, interval.PerfectFifth, interval.MajorSeventh},
	Locrian:                               {interval.PerfectUnison, interval.MinorSecond, interval.DiminishedFifth},
	Quartal:                               {interval.PerfectUnison, interval.PerfectFourth, interval.MinorSeventh},
	QuartalAugmented:                      {interval.PerfectUnison, interval.PerfectFourth, interval.MajorSeventh},

	// Sixth
	MinorSixth:                               {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSixth},
	MinorSixthAddNinth:                       {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSixth, interval.MajorNinth},
	MinorSixthAddFlatNinth:                   {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSixth, interval.MinorNinth},
	MajorSixth:                               {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSixth},
	MajorSixthAddFlatNinth:                   {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSixth, interval.MinorNinth},
	MajorSixthAddNinth:                       {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSixth, interval.MajorNinth},
	MajorSixthFlatFifth:                      {interval.PerfectUnison, interval.MajorThird, interval.DiminishedFifth, interval.MajorSixth},
	MajorSixthSuspendedSecond:                {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth, interval.MajorSixth},
	MajorSixthSuspendedSecondFlatFifth:       {interval.PerfectUnison, interval.MajorSecond, interval.DiminishedFifth, interval.MajorSixth},
	MajorSixthSuspendedSecondDoubleFlatFifth: {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFourth, interval.MajorSixth},
	MajorSixthSuspendedFourth:                {interval.PerfectUnison, interval.PerfectFourth, interval.PerfectFifth, interval.MajorSixth},

	// Seventh
	FullDiminishedSeventh:                       {interval.PerfectUnison, interval.MinorThird, interval.DiminishedFifth, interval.DiminishedSeventh},
	HalfDiminishedSeventh:                       {interval.PerfectUnison, interval.MinorThird, interval.DiminishedFifth, interval.MinorSeventh},
	MinorSeventhDoubleFlatFifth:                 {interval.PerfectUnison, interval.MinorThird, interval.PerfectFourth, interval.MinorSeventh},
	MinorSeventh:                                {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh},
	MinorSeventhAddEleventh:                     {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.PerfectEleventh},
	MinorSeventhAddThirteenth:                   {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorThirteenth},
	MinorSeventhAddSharpEleventh:                {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.AugmentedEleventh},
	MinorSeventhSharpFifth:                      {interval.PerfectUnison, interval.MinorThird, interval.AugmentedFifth, interval.MinorSeventh},
	MinorSeventhFlatNinth:                       {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MinorNinth},
	MinorMajorSeventh:                           {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSeventh},
	MinorMajorSeventhAddEleventh:                {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSeventh, interval.PerfectEleventh},
	MinorMajorSeventhAddThirteenth:              {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorThirteenth},
	DominantSeventhFlatFifth:                    {interval.PerfectUnison, interval.MajorThird, interval.DiminishedFifth, interval.MinorSeventh},
	DominantSeventhSuspendedSecondFlatFifth:     {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth, interval.DiminishedSeventh},
	DominantSeventhSuspendedSecond:              {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth, interval.MinorSeventh},
	DominantSeventhSuspendedFourth:              {interval.PerfectUnison, interval.PerfectFourth, interval.PerfectFifth, interval.MinorSeventh},
	DominantSeventh:                             {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh},
	DominantSeventhAddEleventh:                  {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.PerfectEleventh},
	DominantSeventhAddThirteenth:                {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorThirteenth},
	DominantSeventhAddFourth:                    {interval.PerfectUnison, interval.MajorThird, interval.PerfectFourth, interval.PerfectFifth, interval.MinorSeventh},
	DominantSeventhAddSharpFourth:               {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFourth, interval.PerfectFifth, interval.MinorSeventh},
	DominantSeventhFlatFifthFlatNinth:           {interval.PerfectUnison, interval.MajorThird, interval.DiminishedFifth, interval.MinorSeventh, interval.MinorNinth},
	DominantSeventhSharpFifthFlatNinth:          {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFifth, interval.MinorSeventh, interval.MinorNinth},
	DominantSeventhFlatNinth:                    {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MinorNinth},
	DominantSeventhFlatNinthFlatThirteenth:      {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MinorNinth, interval.MinorThirteenth},
	DominantSeventhSharpNinth:                   {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.AugmentedNinth},
	DominantSeventhSharpNinthSharpEleventh:      {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.AugmentedNinth, interval.AugmentedEleventh},
	DominantSeventhSharpEleventh:                {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.AugmentedEleventh},
	DiminishedMajorSeventh:                      {interval.PerfectUnison, interval.MinorThird, interval.DiminishedFifth, interval.MajorSeventh},
	MajorSeventhFlatFifth:                       {interval.PerfectUnison, interval.MajorThird, interval.DiminishedFifth, interval.MajorSeventh},
	MajorSeventhSuspendedSecond:                 {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth, interval.MajorSeventh},
	MajorSeventhSuspendedFourth:                 {interval.PerfectUnison, interval.PerfectFourth, interval.PerfectFifth, interval.MajorSeventh},
	MajorSeventhSuspendedFourthSharpFifth:       {interval.PerfectUnison, interval.PerfectFourth, interval.AugmentedFifth, interval.MajorSeventh},
	MajorSeventhSuspendedFourthDoubleSharpFifth: {interval.PerfectUnison, interval.PerfectFourth, interval.MajorSixth, interval.MajorSeventh},
	MajorSeventh:                                {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh},
	MajorSeventhAddEleventh:                     {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh, interval.PerfectEleventh},
	MajorSeventhAddThirteenth:                   {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorThirteenth},
	MajorSeventhAddSharpEleventh:                {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh, interval.AugmentedEleventh},
	MajorSeventhAddFourth:                       {interval.PerfectUnison, interval.MajorThird, interval.PerfectFourth, interval.PerfectFifth, interval.MajorSeventh},
	MajorSeventhAddSharpFourth:                  {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFourth, interval.PerfectFifth, interval.MajorSeventh},
	MajorSeventhDoubleSharpFifth:                {interval.PerfectUnison, interval.MajorThird, interval.MajorSixth, interval.MajorSeventh},
	AugmentedMajorSeventh:                       {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFifth, interval.MajorSeventh},
	AugmentedAugmentedSeventh:                   {interval.PerfectUnison, interval.MajorThird, interval.AugmentedFifth, interval.AugmentedSeventh},

	// Ninth
	MinorNinth:                   {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth},
	MinorMajorNinth:              {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth},
	DominantNinth:                {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth},
	DominantNinthSuspendedSecond: {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth},
	DominantNinthSuspendedFourth: {interval.PerfectUnison, interval.PerfectFourth, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth},
	DominantNinthSharpEleventh:   {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth, interval.AugmentedEleventh},
	DominantNinthFlatThirteenth:  {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth, interval.MinorThirteenth},
	MajorNinth:                   {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth},
	MajorNinthSuspendedSecond:    {interval.PerfectUnison, interval.MajorSecond, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth},
	MajorNinthSuspendedFourth:    {interval.PerfectUnison, interval.PerfectFourth, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth},

	// Eleventh
	MinorEleventh:      {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth, interval.PerfectEleventh},
	MinorMajorEleventh: {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth, interval.PerfectEleventh},
	DominantEleventh:   {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth, interval.PerfectEleventh},
	MajorEleventh:      {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth, interval.PerfectEleventh},

	// Thirteenth
	MinorThirteenth:             {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth, interval.PerfectEleventh, interval.MajorThirteenth},
	MinorMajorThirteenth:        {interval.PerfectUnison, interval.MinorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth, interval.PerfectEleventh, interval.MajorThirteenth},
	DominantThirteenth:          {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MajorNinth, interval.PerfectEleventh, interval.MajorThirteenth},
	DominantThirteenthFlatNinth: {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MinorSeventh, interval.MinorNinth, interval.PerfectEleventh, interval.MajorThirteenth},
	MajorThirteenth:             {interval.PerfectUnison, interval.MajorThird, interval.PerfectFifth, interval.MajorSeventh, interval.MajorNinth, interval.PerfectEleventh, interval.MajorThirteenth},
}

// Intervals return interval pattern for chord
func (t Type) Intervals() []interval.Interval {
	if intervals, found := typeIntervals[t]; found {
		return intervals
	}

	panic(fmt.Errorf("no such interval pattern for chord %s", t))
}
