package scale

import "sort"

// Scale is musical scale type it corresponds to pitch interval patterns
type Scale int

// Scale enumeration
const (
	Invalid Scale = iota

	// 3 notes
	Minoric Scale = iota

	// 4 notes
	Thaptic  Scale = iota
	Epathic  Scale = iota
	Zyphic   Scale = iota
	Aeoloric Scale = iota
	Daric    Scale = iota
	Saric    Scale = iota
	Aeolic   Scale = iota
	Stathic  Scale = iota
	Phrynic  Scale = iota

	// 5 notes
	Epathitonic   Scale = iota
	Magitonic     Scale = iota
	Gathitonic    Scale = iota
	Zolitonic     Scale = iota
	Phraditonic   Scale = iota
	Aeracritonic  Scale = iota
	Lothitonic    Scale = iota
	Dolitonic     Scale = iota
	Molitonic     Scale = iota
	Ionaditonic   Scale = iota
	Zacritonic    Scale = iota
	Aeolyritonic  Scale = iota
	Zathitonic    Scale = iota
	Aeolanitonic  Scale = iota
	Aeolacritonic Scale = iota
	Lycritonic    Scale = iota
	Epygitonic    Scale = iota
	Zothitonic    Scale = iota
	Ionyptitonic  Scale = iota
	Thoditonic    Scale = iota
	Kataritonic   Scale = iota
	Lyditonic     Scale = iota
	Zylitonic     Scale = iota
	Ranitonic     Scale = iota
	Ryphitonic    Scale = iota
	Phronitonic   Scale = iota
	Aerynitonic   Scale = iota
	Ionoditonic   Scale = iota
	Mixitonic     Scale = iota
	Thyritonic    Scale = iota
	Bolitonic     Scale = iota

	// 6 notes
	Epathimic  Scale = iota
	Dyrimic    Scale = iota
	Daptimic   Scale = iota
	Epygimic   Scale = iota
	Bylimic    Scale = iota
	Garimic    Scale = iota
	Thonimic   Scale = iota
	Mythimic   Scale = iota
	Mixolimic  Scale = iota
	Ionacrimic Scale = iota
	Dalimic    Scale = iota
	Darmic     Scale = iota
	Phracrimic Scale = iota
	Starimic   Scale = iota
	Bygimic    Scale = iota
	Stalimic   Scale = iota
	Rycrimic   Scale = iota
	Katythimic Scale = iota
	Aerycrimic Scale = iota
	Phralimic  Scale = iota
	Solimic    Scale = iota
	Thogimic   Scale = iota
	Laptimic   Scale = iota
	Modimic    Scale = iota
	Syrimic    Scale = iota
	Bolimic    Scale = iota
	Kanimic    Scale = iota
	Pynimic    Scale = iota
	Kytrimic   Scale = iota
	Palimic    Scale = iota
	Ionodimic  Scale = iota
	Mydimic    Scale = iota
	Zagimic    Scale = iota
	Mothimic   Scale = iota
	Bocrimic   Scale = iota
	Zacrimic   Scale = iota
	Ionythimic Scale = iota
	Dynimic    Scale = iota
	Ponimic    Scale = iota
	Lathimic   Scale = iota
	Galimic    Scale = iota
	Katanimic  Scale = iota
	Manimic    Scale = iota
	Kocrimic   Scale = iota
	Mycrimic   Scale = iota
	Aerothimic Scale = iota
	Epynimic   Scale = iota
	Thoptimic  Scale = iota
	Thagimic   Scale = iota
	Dagimic    Scale = iota
	Thalimic   Scale = iota
	Stythimic  Scale = iota
	Aeragimic  Scale = iota
	Aeradimic  Scale = iota
	Lythimic   Scale = iota
	Boptimic   Scale = iota
	Dathimic   Scale = iota
	Gacrimic   Scale = iota
	WholeTone  Scale = iota

	// 7 notes
	Lydian     Scale = iota
	Ionythian  Scale = iota
	Ionarian   Scale = iota
	Aeolacrian Scale = iota
	Zogian     Scale = iota
	Phrolian   Scale = iota
	Soptian    Scale = iota
	Thonian    Scale = iota
	Epathian   Scale = iota
	Stathian   Scale = iota
	Epogian    Scale = iota
	Ionycrian  Scale = iota
	Aerathian  Scale = iota
	Aeopian    Scale = iota
	Epycrian   Scale = iota
	Parian     Scale = iota
	Stygian    Scale = iota
	Stythian   Scale = iota
	Zorian     Scale = iota
	Thacrian   Scale = iota
	Aeolathian Scale = iota
	Saptian    Scale = iota
	Thycrian   Scale = iota
	Stalian    Scale = iota
	Morian     Scale = iota
	Phraptian  Scale = iota
	Katocrian  Scale = iota
	Ionopian   Scale = iota
	Rythian    Scale = iota
	Laptian    Scale = iota
	Katathian  Scale = iota
	Tholian    Scale = iota
	Zarian     Scale = iota
	Ranian     Scale = iota
	Golian     Scale = iota
	Palian     Scale = iota
	Loptian    Scale = iota
	Bycrian    Scale = iota
	Dolian     Scale = iota
	Pagian     Scale = iota
	Larian     Scale = iota
	Thadian    Scale = iota
	Pythian    Scale = iota
	Katogian   Scale = iota
	Katacrian  Scale = iota
	Stanian    Scale = iota
	Katyptian  Scale = iota
	Bonian     Scale = iota
	Epagian    Scale = iota
	Gacrian    Scale = iota
	Katoptian  Scale = iota
	Epacrian   Scale = iota
	Pogian     Scale = iota
	Eporian    Scale = iota
	Korian     Scale = iota
	Mycrian    Scale = iota
	Ionanian   Scale = iota
	Gydian     Scale = iota
	Aeolynian  Scale = iota

	// 8 notes
	Aerycryllic     Scale = iota
	Pynyllic        Scale = iota
	Pothyllic       Scale = iota
	Locryllic       Scale = iota
	Aeracryllic     Scale = iota
	Dydyllic        Scale = iota
	Dagyllic        Scale = iota
	Sydyllic        Scale = iota
	Bacryllic       Scale = iota
	Rocryllic       Scale = iota
	Ionocryllic     Scale = iota
	Ionoptyllic     Scale = iota
	Dolyllic        Scale = iota
	Thagyllic       Scale = iota
	Aeoladyllic     Scale = iota
	Phroryllic      Scale = iota
	Phranyllic      Scale = iota
	Lydyllic        Scale = iota
	Zoryllic        Scale = iota
	Mixolydyllic    Scale = iota
	Ioniptyllic     Scale = iota
	Aeolothyllic    Scale = iota
	Maptyllic       Scale = iota
	Thyptyllic      Scale = iota
	Doptyllic       Scale = iota
	Lonyllic        Scale = iota
	Aeolathyllic    Scale = iota
	Zagyllic        Scale = iota
	Gythyllic       Scale = iota
	Tharyllic       Scale = iota
	Kataryllic      Scale = iota
	Stogyllic       Scale = iota
	Dalyllic        Scale = iota
	Stycryllic      Scale = iota
	Stolyllic       Scale = iota
	Zaryllic        Scale = iota
	Aeronyllic      Scale = iota
	Stacryllic      Scale = iota
	Thyryllic       Scale = iota
	Racryllic       Scale = iota
	Epotyllic       Scale = iota
	MajorDiminished Scale = iota

	// 9 Notes
	Aerycrygic  Scale = iota
	Kalygic     Scale = iota
	Mixolydygic Scale = iota
	Dycrygic    Scale = iota
	Loptygic    Scale = iota
	Bagygic     Scale = iota
	Apinygic    Scale = iota
	Epyrygic    Scale = iota
	Aeolorygic  Scale = iota
	Manygic     Scale = iota
	Sathygic    Scale = iota
	Phrygic     Scale = iota
	Zothygic    Scale = iota
	Aeolacrygic Scale = iota
	Kyrygic     Scale = iota
	Porygic     Scale = iota
	Kaptygic    Scale = iota
	Koptygic    Scale = iota
	Phronygic   Scale = iota

	// 10 Notes
	Aerycryllian Scale = iota
	Katoryllian  Scale = iota
	Sydyllian    Scale = iota
	Gothyllian   Scale = iota
	Staptyllian  Scale = iota
	Thydyllian   Scale = iota

	// 11 Notes
	Aerycratic Scale = iota

	// 12 Notes
	Chromatic Scale = iota
)

var scaleStringifieds = map[Scale]string{
	Invalid: "Invalid",

	// 3 Notes
	Minoric: "Minoric",

	// 4 Notes
	Thaptic:  "Thaptic",
	Epathic:  "Epathic",
	Zyphic:   "Zyphic",
	Aeoloric: "Aeoloric",
	Daric:    "Daric",
	Saric:    "Saric",
	Aeolic:   "Aeolic",
	Stathic:  "Stathic",
	Phrynic:  "Phrynic",

	// 5 Notes
	Epathitonic:   "Epathitonic",
	Magitonic:     "Magitonic",
	Gathitonic:    "Gathitonic",
	Zolitonic:     "Zolitonic",
	Phraditonic:   "Phraditonic",
	Aeracritonic:  "Aeracritonic",
	Lothitonic:    "Lothitonic",
	Dolitonic:     "Dolitonic",
	Molitonic:     "Molitonic",
	Ionaditonic:   "Ionaditonic",
	Zacritonic:    "Zacritonic",
	Aeolyritonic:  "Aeolyritonic",
	Zathitonic:    "Zathitonic",
	Aeolanitonic:  "Aeolanitonic",
	Aeolacritonic: "Aeolacritonic",
	Lycritonic:    "Lycritonic",
	Epygitonic:    "Epygitonic",
	Zothitonic:    "Zothitonic",
	Ionyptitonic:  "Ionyptitonic",
	Thoditonic:    "Thoditonic",
	Kataritonic:   "Kataritonic",
	Lyditonic:     "Lyditonic",
	Zylitonic:     "Zylitonic",
	Ranitonic:     "Ranitonic",
	Ryphitonic:    "Ryphitonic",
	Phronitonic:   "Phronitonic",
	Aerynitonic:   "Aerynitonic",
	Ionoditonic:   "Ionoditonic",
	Mixitonic:     "Mixitonic",
	Thyritonic:    "Thyritonic",
	Bolitonic:     "Bolitonic",

	// 6 Notes
	Epathimic:  "Epathimic",
	Dyrimic:    "Dyrimic",
	Daptimic:   "Daptimic",
	Epygimic:   "Epygimic",
	Bylimic:    "Bylimic",
	Garimic:    "Garimic",
	Thonimic:   "Thonimic",
	Mythimic:   "Mythimic",
	Mixolimic:  "Mixolimic",
	Ionacrimic: "Ionacrimic",
	Dalimic:    "Dalimic",
	Darmic:     "Darmic",
	Phracrimic: "Phracrimic",
	Starimic:   "Starimic",
	Bygimic:    "Bygimic",
	Stalimic:   "Stalimic",
	Rycrimic:   "Rycrimic",
	Katythimic: "Katythimic",
	Aerycrimic: "Aerycrimic",
	Phralimic:  "Phralimic",
	Solimic:    "Solimic",
	Thogimic:   "Thogimic",
	Laptimic:   "Laptimic",
	Modimic:    "Modimic",
	Syrimic:    "Syrimic",
	Bolimic:    "Bolimic",
	Kanimic:    "Kanimic",
	Pynimic:    "Pynimic",
	Kytrimic:   "Kytrimic",
	Palimic:    "Palimic",
	Ionodimic:  "Ionodimic",
	Mydimic:    "Mydimic",
	Zagimic:    "Zagimic",
	Mothimic:   "Mothimic",
	Bocrimic:   "Bocrimic",
	Zacrimic:   "Zacrimic",
	Ionythimic: "Ionythimic",
	Dynimic:    "Dynimic",
	Ponimic:    "Ponimic",
	Lathimic:   "Lathimic",
	Galimic:    "Galimic",
	Katanimic:  "Katanimic",
	Manimic:    "Manimic",
	Kocrimic:   "Kocrimic",
	Mycrimic:   "Mycrimic",
	Aerothimic: "Aerothimic",
	Epynimic:   "Epynimic",
	Thoptimic:  "Thoptimic",
	Thagimic:   "Thagimic",
	Dagimic:    "Dagimic",
	Thalimic:   "Thalimic",
	Stythimic:  "Stythimic",
	Aeragimic:  "Aeragimic",
	Aeradimic:  "Aeradimic",
	Lythimic:   "Lythimic",
	Boptimic:   "Boptimic",
	Dathimic:   "Dathimic",
	Gacrimic:   "Gacrimic",
	WholeTone:  "WholeTone",

	// 7 Notes
	Lydian:     "Lydian",
	Ionythian:  "Ionythian",
	Ionarian:   "Ionarian",
	Aeolacrian: "Aeolacrian",
	Zogian:     "Zogian",
	Phrolian:   "Phrolian",
	Soptian:    "Soptian",
	Thonian:    "Thonian",
	Epathian:   "Epathian",
	Stathian:   "Stathian",
	Epogian:    "Epogian",
	Ionycrian:  "Ionycrian",
	Aerathian:  "Aerathian",
	Aeopian:    "Aeopian",
	Epycrian:   "Epycrian",
	Parian:     "Parian",
	Stygian:    "Stygian",
	Stythian:   "Stythian",
	Zorian:     "Zorian",
	Thacrian:   "Thacrian",
	Aeolathian: "Aeolathian",
	Saptian:    "Saptian",
	Thycrian:   "Thycrian",
	Stalian:    "Stalian",
	Morian:     "Morian",
	Phraptian:  "Phraptian",
	Katocrian:  "Katocrian",
	Ionopian:   "Ionopian",
	Rythian:    "Rythian",
	Laptian:    "Laptian",
	Katathian:  "Katathian",
	Tholian:    "Tholian",
	Zarian:     "Zarian",
	Ranian:     "Ranian",
	Golian:     "Golian",
	Palian:     "Palian",
	Loptian:    "Loptian",
	Bycrian:    "Bycrian",
	Dolian:     "Dolian",
	Pagian:     "Pagian",
	Larian:     "Larian",
	Thadian:    "Thadian",
	Pythian:    "Pythian",
	Katogian:   "Katogian",
	Katacrian:  "Katacrian",
	Stanian:    "Stanian",
	Katyptian:  "Katyptian",
	Bonian:     "Bonian",
	Epagian:    "Epagian",
	Gacrian:    "Gacrian",
	Katoptian:  "Katoptian",
	Epacrian:   "Epacrian",
	Pogian:     "Pogian",
	Eporian:    "Eporian",
	Korian:     "Korian",
	Mycrian:    "Mycrian",
	Ionanian:   "Ionanian",
	Gydian:     "Gydian",
	Aeolynian:  "Aeolynian",

	// 8 notes
	Aerycryllic:     "Aerycryllic",
	Pynyllic:        "Pynyllic",
	Pothyllic:       "Pothyllic",
	Locryllic:       "Locryllic",
	Aeracryllic:     "Aeracryllic",
	Dydyllic:        "Dydyllic",
	Dagyllic:        "Dagyllic",
	Sydyllic:        "Sydyllic",
	Bacryllic:       "Bacryllic",
	Rocryllic:       "Rocryllic",
	Ionocryllic:     "Ionocryllic",
	Ionoptyllic:     "Ionoptyllic",
	Dolyllic:        "Dolyllic",
	Thagyllic:       "Thagyllic",
	Aeoladyllic:     "Aeoladyllic",
	Phroryllic:      "Phroryllic",
	Phranyllic:      "Phranyllic",
	Lydyllic:        "Lydyllic",
	Zoryllic:        "Zoryllic",
	Mixolydyllic:    "Mixolydyllic",
	Ioniptyllic:     "Ioniptyllic",
	Aeolothyllic:    "Aeolothyllic",
	Maptyllic:       "Maptyllic",
	Thyptyllic:      "Thyptyllic",
	Doptyllic:       "Doptyllic",
	Lonyllic:        "Lonyllic",
	Aeolathyllic:    "Aeolathyllic",
	Zagyllic:        "Zagyllic",
	Gythyllic:       "Gythyllic",
	Tharyllic:       "Tharyllic",
	Kataryllic:      "Kataryllic",
	Stogyllic:       "Stogyllic",
	Dalyllic:        "Dalyllic",
	Stycryllic:      "Stycryllic",
	Stolyllic:       "Stolyllic",
	Zaryllic:        "Zaryllic",
	Aeronyllic:      "Aeronyllic",
	Stacryllic:      "Stacryllic",
	Thyryllic:       "Thyryllic",
	Racryllic:       "Racryllic",
	Epotyllic:       "Epotyllic",
	MajorDiminished: "MajorDiminished",

	// 9 Notes
	Aerycrygic:  "Aerycrygic",
	Kalygic:     "Kalygic",
	Mixolydygic: "Mixolydygic",
	Dycrygic:    "Dycrygic",
	Loptygic:    "Loptygic",
	Bagygic:     "Bagygic",
	Apinygic:    "Apinygic",
	Epyrygic:    "Epyrygic",
	Aeolorygic:  "Aeolorygic",
	Manygic:     "Manygic",
	Sathygic:    "Sathygic",
	Phrygic:     "Phrygic",
	Zothygic:    "Zothygic",
	Aeolacrygic: "Aeolacrygic",
	Kyrygic:     "Kyrygic",
	Porygic:     "Porygic",
	Kaptygic:    "Kaptygic",
	Koptygic:    "Koptygic",
	Phronygic:   "Phronygic",

	// 10 Notes
	Aerycryllian: "Aerycryllian",
	Katoryllian:  "Katoryllian",
	Sydyllian:    "Sydyllian",
	Gothyllian:   "Gothyllian",
	Staptyllian:  "Staptyllian",
	Thydyllian:   "Thydyllian",

	// 11 Notes
	Aerycratic: "Aerycratic",

	// 12 Notes
	Chromatic: "Chromatic",
}

// String returns stringified version of a scale
func (s Scale) String() string {
	return scaleStringifieds[s]
}

// AllScales returns valid scale
func AllScales() []Scale {
	scales := make([]Scale, 0)
	for k := range scaleStringifieds {
		if k != Invalid {
			scales = append(scales, k)
		}
	}

	// sort scales
	sort.SliceStable(scales, func(i, j int) bool {
		return scales[i] < scales[j]
	})
	return scales
}
