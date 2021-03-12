package modetype

// References : https://allthescales.org/scales.php

// Type represents musical mode type
type Type int

// Musical modes enumeration
const (
	Invalid Type = iota

	// 3.1 Notes
	Minoric Type = iota

	// 4.1 Thaptic scale
	Thaptic  Type = iota
	Lothic   Type = iota
	Phratic  Type = iota
	Aerathic Type = iota

	// 4.2 Epathic scale
	Epathic Type = iota
	Mynic   Type = iota
	Rothic  Type = iota
	Eporic  Type = iota

	// 4.3 Zyphic scale
	Zyphic Type = iota
	Epogic Type = iota
	Lanic  Type = iota
	Pyrric Type = iota

	// 4.4 Aeoloric scale
	Aeoloric Type = iota
	Gonic    Type = iota
	Dalic    Type = iota
	Dygic    Type = iota

	// 4.5 Daric scale
	Daric   Type = iota
	Lonic   Type = iota
	Phradic Type = iota
	Bolic   Type = iota

	// 4.6 Saric scale
	Saric    Type = iota
	Zoptic   Type = iota
	Aeraphic Type = iota
	Byptic   Type = iota

	// 4.7 Aeolic scale
	Aeolic    Type = iota
	Koptic    Type = iota
	Mixolyric Type = iota
	Lydic     Type = iota

	// 4.8 Saric scale
	Stathic Type = iota
	Dadic   Type = iota

	// 4.9 Phrynic scale
	Phrynic Type = iota

	// 5.1 Epathitonic scale
	Epathitonic Type = iota
	Mynitonic   Type = iota
	Rocritonic  Type = iota
	Pentatonic  Type = iota
	Thaptitonic Type = iota

	// 5.2 Magitonic scale
	Magitonic      Type = iota
	Daditonic      Type = iota
	Aeolyphritonic Type = iota
	Gycritonic     Type = iota
	Pyritonic      Type = iota

	// 5.3 Gathitonic scale
	Gathitonic  Type = iota
	Ionitonic   Type = iota
	Phrynitonic Type = iota
	Stathitonic Type = iota
	Thalitonic  Type = iota

	// 5.4 Zolitonic scale
	Zolitonic    Type = iota
	Epogitonic   Type = iota
	Lanitonic    Type = iota
	Paptitonic   Type = iota
	Ionacritonic Type = iota

	// 5.5 Phraditonic scale
	Phraditonic  Type = iota
	Aeoloritonic Type = iota
	Gonitonic    Type = iota
	Dalitonic    Type = iota
	Dygitonic    Type = iota

	// 5.6 Aeracritonic scale
	Aeracritonic Type = iota
	Byptitonic   Type = iota
	Daritonic    Type = iota
	Lonitonic    Type = iota
	Ionycritonic Type = iota

	// 5.7 Lothitonic scale
	Lothitonic   Type = iota
	Phratonic    Type = iota
	Aerathitonic Type = iota
	Saritonic    Type = iota
	Zoptitonic   Type = iota

	// 5.8 Dolitonic scale
	Dolitonic   Type = iota
	Poritonic   Type = iota
	Aerylitonic Type = iota
	Zagitonic   Type = iota
	Lagitonic   Type = iota

	// 5.9 Molitonic scale
	Molitonic   Type = iota
	Staptitonic Type = iota
	Mothitonic  Type = iota
	Aeritonic   Type = iota
	Ragitonic   Type = iota

	// 5.10 Ionaditonic scale
	Ionaditonic   Type = iota
	Bocritonic    Type = iota
	Gythitonic    Type = iota
	Pagitonic     Type = iota
	Aeolythitonic Type = iota

	// 5.11 Zacritonic scale
	Zacritonic  Type = iota
	Laritonic   Type = iota
	Thacritonic Type = iota
	Styditonic  Type = iota
	Loritonic   Type = iota

	// 5.12 Aeolyritonic scale
	Aeolyritonic Type = iota
	Goritonic    Type = iota
	Aeoloditonic Type = iota
	Doptitonic   Type = iota
	Aeraphitonic Type = iota

	// 5.13 Zathitonic scale
	Zathitonic   Type = iota
	Raditonic    Type = iota
	Stonitonic   Type = iota
	Syptitonic   Type = iota
	Ionythitonic Type = iota

	// 5.14 Aeolanitonic scale
	Aeolanitonic Type = iota
	Danitonic    Type = iota
	Ionaritonic  Type = iota
	Dynitonic    Type = iota
	Zyditonic    Type = iota

	// 5.15 Aeolacritonic scale
	Aeolacritonic Type = iota
	Zythitonic    Type = iota
	Dyritonic     Type = iota
	Koptitonic    Type = iota
	Thocritonic   Type = iota

	// 5.16 Lycritonic scale
	Lycritonic Type = iota
	Daptitonic Type = iota
	Kygitonic  Type = iota
	Mocritonic Type = iota
	Zynitonic  Type = iota

	// 5.17 Epygitonic scale
	Epygitonic Type = iota
	Zaptitonic Type = iota
	Kagitonic  Type = iota
	Zogitonic  Type = iota
	Epyritonic Type = iota

	// 5.18 Zothitonic scale
	Zothitonic    Type = iota
	Phrolitonic   Type = iota
	Ionagitonic   Type = iota
	Aeolapritonic Type = iota
	Kyritonic     Type = iota

	// 5.19 Ionyptitonic scale
	Ionyptitonic Type = iota
	Gyritonic    Type = iota
	Zalitonic    Type = iota
	Stolitonic   Type = iota
	Bylitonic    Type = iota

	// 5.20 Thoditonic scale
	Thoditonic  Type = iota
	Dogitonic   Type = iota
	Phralitonic Type = iota
	Garitonic   Type = iota
	Soptitonic  Type = iota

	// 5.21 Kataritonic scale
	Kataritonic Type = iota
	Sylitonic   Type = iota
	Thonitonic  Type = iota
	Phropitonic Type = iota
	Staditonic  Type = iota

	// 5.22 Lyditonic scale
	Lyditonic  Type = iota
	Mythitonic Type = iota
	Sogitonic  Type = iota
	Gothitonic Type = iota
	Rothitonic Type = iota

	// 5.23 Zylitonic scale
	Zylitonic    Type = iota
	Zoditonic    Type = iota
	Zaritonic    Type = iota
	Phrythitonic Type = iota
	Rolitonic    Type = iota

	// 5.24 Ranitonic scale
	Ranitonic    Type = iota
	Laditonic    Type = iota
	Poditonic    Type = iota
	Ionothitonic Type = iota
	Kanitonic    Type = iota

	// 5.25 Ryphitonic scale
	Ryphitonic    Type = iota
	Gylitonic     Type = iota
	Aeolycritonic Type = iota
	Pynitonic     Type = iota
	Zanitonic     Type = iota

	// 5.26 Phronitonic scale
	Phronitonic Type = iota
	Banitonic   Type = iota
	Aeronitonic Type = iota
	Golitonic   Type = iota
	Dyptitonic  Type = iota

	// 5.27 Aerynitonic scale
	Aerynitonic  Type = iota
	Palitonic    Type = iota
	Stothitonic  Type = iota
	Aerophitonic Type = iota
	Katagitonic  Type = iota

	// 5.28 Ionoditonic scale
	Ionoditonic Type = iota
	Bogitonic   Type = iota
	Mogitonic   Type = iota
	Docritonic  Type = iota
	Epaditonic  Type = iota

	// 5.29 Mixitonic scale
	Mixitonic    Type = iota
	Phrothitonic Type = iota
	Katycritonic Type = iota
	Ionalitonic  Type = iota
	Loptitonic   Type = iota

	// 5.30 Thyritonic scale
	Thyritonic  Type = iota
	Thoptitonic Type = iota
	Bycritonic  Type = iota
	Pathitonic  Type = iota
	Myditonic   Type = iota

	// 5.31 Bolitonic scale
	Bolitonic   Type = iota
	Bothitonic  Type = iota
	Kataditonic Type = iota
	Koditonic   Type = iota
	Tholitonic  Type = iota

	// 6.1 Epathimic scale
	Epathimic Type = iota
	Mynimic   Type = iota
	Rocrimic  Type = iota
	Eporimic  Type = iota
	Thaptimic Type = iota
	Lothimic  Type = iota

	// 6.2 Dyrimic scale
	Dyrimic    Type = iota
	Koptimic   Type = iota
	Thocrimic  Type = iota
	Aeolanimic Type = iota
	Danimic    Type = iota
	Ionarimic  Type = iota

	// 6.3 Daptimic scale
	Daptimic Type = iota
	Kygimic  Type = iota
	Mocrimic Type = iota
	Zynimic  Type = iota
	Aeolimic Type = iota
	Zythimic Type = iota

	// 6.4 Epygimic scale
	Epygimic Type = iota
	Zaptimic Type = iota
	Kagimic  Type = iota
	Zogimic  Type = iota
	Epyrimic Type = iota
	Lycrimic Type = iota

	// 6.5 Bylimic scale
	Bylimic     Type = iota
	Zothimic    Type = iota
	Phrolimic   Type = iota
	Ionagimic   Type = iota
	Aeolaphimic Type = iota
	Kycrimic    Type = iota

	// 6.6 Garimic scale
	Garimic    Type = iota
	Soptimic   Type = iota
	Ionyptimic Type = iota
	Gyrimic    Type = iota
	Zalimic    Type = iota
	Stolimic   Type = iota

	// 6.7 Thonimic scale
	Thonimic Type = iota
	Stadimic Type = iota
	Thodimic Type = iota

	// 6.8 Mythimic scale
	Mythimic  Type = iota
	Sogimic   Type = iota
	Gogimic   Type = iota
	Rothimic  Type = iota
	Katarimic Type = iota
	Sylimic   Type = iota

	// 6.9 Mixolimic scale
	Mixolimic   Type = iota
	Dadimic     Type = iota
	Aeolyphimic Type = iota
	Gycrimic    Type = iota
	Pyrimic     Type = iota
	Lydimic     Type = iota

	// 6.10 Ionacrimic scale
	Ionacrimic Type = iota
	Gathimic   Type = iota
	Ionynimic  Type = iota
	Phrynimic  Type = iota
	Stathimic  Type = iota
	Thatimic   Type = iota

	// 6.11 Dalimic scale
	Dalimic  Type = iota
	Dygimic  Type = iota
	Zolimic  Type = iota
	Epogimic Type = iota
	Lanimic  Type = iota
	Paptimic Type = iota

	// 6.12 Darmic scale
	Darmic     Type = iota
	Lonimic    Type = iota
	Ionycrimic Type = iota
	Phradimic  Type = iota
	Aeolorimic Type = iota
	Gonimic    Type = iota

	// 6.13 Phracrimic scale
	Phracrimic Type = iota
	Aerathimic Type = iota
	Sarimic    Type = iota
	Zoptimic   Type = iota
	Zeracrimic Type = iota
	Byptimic   Type = iota

	// 6.14 Starimic scale
	Starimic   Type = iota
	Phrathimic Type = iota
	Saptimic   Type = iota
	Aerodimic  Type = iota
	Macrimic   Type = iota
	Rogimic    Type = iota

	// 6.15 Bygimic scale
	Bygimic    Type = iota
	Thycrimic  Type = iota
	Aeoladimic Type = iota
	Dylimic    Type = iota
	Eponimic   Type = iota
	Katygimic  Type = iota

	// 6.16 Stalimic scale
	Stalimic    Type = iota
	Stoptimic   Type = iota
	Zygimic     Type = iota
	Kataptimic  Type = iota
	Aeolaptimic Type = iota
	Pothimic    Type = iota

	// 6.17 Rycrimic scale
	Rycrimic  Type = iota
	Ronimic   Type = iota
	Stycrimic Type = iota
	Katorimic Type = iota
	Epythimic Type = iota
	Kaptimic  Type = iota

	// 6.18 Katythimic scale
	Katythimic Type = iota
	Madimic    Type = iota
	Aerygimic  Type = iota
	Pylimic    Type = iota
	Ionathimic Type = iota
	Morimic    Type = iota

	// 6.19 Aerycrimic scale
	Aerycrimic Type = iota
	Ganimic    Type = iota
	Eparimic   Type = iota
	Lyrimic    Type = iota
	Phraptimic Type = iota
	Bacrimic   Type = iota

	// 6.20 Phralimic scale
	Phralimic  Type = iota
	Phrogimic  Type = iota
	Rathimic   Type = iota
	Katocrimic Type = iota
	Phryptimic Type = iota
	Katynimic  Type = iota

	// 6.21 Solimic scale
	Solimic    Type = iota
	Ionolimic  Type = iota
	Ionophimic Type = iota
	Aeologimic Type = iota
	Zadimic    Type = iota
	Sygimic    Type = iota

	// 6.22 Thogimic scale
	Thogimic    Type = iota
	Rythimic    Type = iota
	Donimic     Type = iota
	Aeoloptimic Type = iota
	Panimic     Type = iota
	Lodimic     Type = iota

	// 6.23 Laptimic scale
	Laptimic   Type = iota
	Lygimic    Type = iota
	Logimic    Type = iota
	Lalimic    Type = iota
	Sothimic   Type = iota
	Phrocrimic Type = iota

	// 6.24 Modimic scale
	Modimic    Type = iota
	Barimic    Type = iota
	Poptimic   Type = iota
	Sagimic    Type = iota
	Aelothimic Type = iota
	Socrimic   Type = iota

	// 6.25 Syrimic scale
	Syrimic    Type = iota
	Stodimic   Type = iota
	Ionocrimic Type = iota
	Zycrimic   Type = iota
	Ionygimic  Type = iota
	Katathimic Type = iota

	// 6.26 Bolimic scale
	Bolimic   Type = iota
	Bothimic  Type = iota
	Katadimic Type = iota
	Kodimic   Type = iota
	Tholimic  Type = iota
	Ralimic   Type = iota

	// 6.27 Kanimic scale
	Kanimic    Type = iota
	Zylimic    Type = iota
	Zodimic    Type = iota
	Zarimic    Type = iota
	Phrythimic Type = iota
	Rorimic    Type = iota

	// 6.28 Pynimic scale
	Pynimic    Type = iota
	Zanimic    Type = iota
	Ranimic    Type = iota
	Ladimic    Type = iota
	Podimic    Type = iota
	Ionothimic Type = iota

	// 6.29 Kytrimic scale
	Kytrimic    Type = iota
	Golimic     Type = iota
	Dyptimic    Type = iota
	Ryrimic     Type = iota
	Gylimic     Type = iota
	Aeolycrimic Type = iota

	// 6.30 Palimic scale
	Palimic   Type = iota
	Stothimic Type = iota
	Aeronimic Type = iota
	Katagimic Type = iota
	Phronimic Type = iota
	Banimic   Type = iota

	// 6.31 Ionodimic scale
	Ionodimic Type = iota
	Bogimic   Type = iota
	Mogimic   Type = iota
	Docrimic  Type = iota
	Epadimic  Type = iota
	Aerynimic Type = iota

	// 6.32 Mydimic scale
	Mydimic    Type = iota
	Thyptimic  Type = iota
	Phrothimic Type = iota
	Katycrimic Type = iota
	Ionalimic  Type = iota
	Loptimic   Type = iota

	// 6.33 Zagimic scale
	Zagimic   Type = iota
	Lagimic   Type = iota
	Thyrimic  Type = iota
	Thothimic Type = iota
	Bycrimic  Type = iota
	Pathimic  Type = iota

	// 6.34 Mothimic scale
	Mothimic  Type = iota
	Aeranimic Type = iota
	Ragimic   Type = iota
	Dolimic   Type = iota
	Porimic   Type = iota
	Aerylimic Type = iota

	// 6.35 Bocrimic scale
	Bocrimic    Type = iota
	Gythimic    Type = iota
	Pagimic     Type = iota
	Aeolythimic Type = iota
	Molimic     Type = iota
	Staptimic   Type = iota

	// 6.36 Zacrimic scale
	Zacrimic  Type = iota
	Larimic   Type = iota
	Thacrimic Type = iota
	Stydimic  Type = iota
	Lorimic   Type = iota
	Ionadimic Type = iota

	// 6.37 Ionythimic scale
	Ionythimic Type = iota
	Aerythimic Type = iota

	// 6.38 Dynimic scale
	Dynimic  Type = iota
	Zydimic  Type = iota
	Zathimic Type = iota
	Radimic  Type = iota
	Stonimic Type = iota
	Syptimic Type = iota

	// 6.39 Ponimic scale
	Ponimic  Type = iota
	Kadimic  Type = iota
	Gynimic  Type = iota
	Thydimic Type = iota
	Polimic  Type = iota
	Thanimic Type = iota

	// 6.40 Lathimic scale
	Lathimic   Type = iota
	Aeralimic  Type = iota
	Kynimic    Type = iota
	Stynimic   Type = iota
	Epytimic   Type = iota
	Katoptimic Type = iota

	// 6.41 Galimic scale
	Galimic   Type = iota
	Kathimic  Type = iota
	Lylimic   Type = iota
	Epalimic  Type = iota
	Epacrimic Type = iota
	Sathimic  Type = iota

	// 6.42 Katanimic scale
	Katanimic  Type = iota
	Katyrimic  Type = iota
	Rynimic    Type = iota
	Pogimic    Type = iota
	Aeraptimic Type = iota
	Epylimic   Type = iota

	// 6.43 Manimic scale
	Manimic  Type = iota
	Marimic  Type = iota
	Locrimic Type = iota
	Rylimic  Type = iota
	Epatimic Type = iota
	Byrimic  Type = iota

	// 6.44 Kocrimic scale
	Kocrimic  Type = iota
	Korimic   Type = iota
	Lynimic   Type = iota
	Malimic   Type = iota
	Synimic   Type = iota
	Phragimic Type = iota

	// 6.45 Mycrimic scale
	Mycrimic   Type = iota
	Ionorimic  Type = iota
	Phrydimic  Type = iota
	Zyptimic   Type = iota
	Katothimic Type = iota
	Phrylimic  Type = iota

	// 6.46 Aerothimic scale
	Aerothimic Type = iota
	Stagimic   Type = iota
	Dorimic    Type = iota
	Phrycrimic Type = iota
	Kyptimic   Type = iota
	Ionylimic  Type = iota

	// 6.47 Epynimic scale
	Epynimic   Type = iota
	Ionogimic  Type = iota
	Kydimic    Type = iota
	Gaptimic   Type = iota
	Tharimic   Type = iota
	Ionaphimic Type = iota

	// 6.48 Thoptimic scale
	Thoptimic  Type = iota
	Bagimic    Type = iota
	Kyrimic    Type = iota
	Sonimic    Type = iota
	Aeolonimic Type = iota
	Rygimic    Type = iota

	// 6.49 Thagimic scale
	Thagimic  Type = iota
	Kolimic   Type = iota
	Dycrimic  Type = iota
	Epycrimic Type = iota
	Gocrimic  Type = iota
	Katolimic Type = iota

	// 6.50 Dagimic scale
	Dagimic    Type = iota
	Aeolydimic Type = iota
	Parimic    Type = iota
	Ionaptimic Type = iota
	Thylimic   Type = iota
	Lolimic    Type = iota

	// 6.51 Thalimic scale
	Thalimic   Type = iota
	Stygimic   Type = iota
	Aeolygimic Type = iota
	Aerogimic  Type = iota
	Dacrimic   Type = iota
	Baptimic   Type = iota

	// 6.52 Stythimic scale
	Stythimic Type = iota
	Kothimic  Type = iota
	Pygimic   Type = iota
	Rodimic   Type = iota
	Sorimic   Type = iota
	Monimic   Type = iota

	// 6.53 Aeragimic scale
	Aeragimic Type = iota
	Epothimic Type = iota
	Salimic   Type = iota
	Lyptimic  Type = iota
	Katonimic Type = iota
	Gygimic   Type = iota

	// 6.54 Aeradimic scale
	Aeradimic Type = iota
	Zyrimic   Type = iota
	Stylimic  Type = iota

	// 6.55 Lythimic scale
	Lythimic  Type = iota
	Dodimic   Type = iota
	Katalimic Type = iota

	// 6.56 Boptimic scale
	Boptimic    Type = iota
	Stogimic    Type = iota
	Thynimic    Type = iota
	Aeolathimic Type = iota
	Bythimic    Type = iota
	Padimic     Type = iota

	// 6.57 Dathimic scale
	Dathimic Type = iota
	Epagimic Type = iota
	Raptimic Type = iota
	Epolimic Type = iota
	Sythimic Type = iota
	Sydimic  Type = iota

	// 6.58 Gacrimic
	Gacrimic    Type = iota
	Borimic     Type = iota
	Sycrimic    Type = iota
	Gadimic     Type = iota
	Aeolocrimic Type = iota
	Phrygimic   Type = iota

	// 6.59 WholeTone
	WholeTone Type = iota

	// 7.1 Lydian scale
	Lydian     Type = iota
	Mixolydian Type = iota
	Aeolian    Type = iota
	Locrian    Type = iota
	Ionian     Type = iota
	Dorian     Type = iota
	Phrygian   Type = iota

	// 7.2 Ionythian scale
	Ionythian Type = iota
	Aeolyrian Type = iota
	Gorian    Type = iota
	Aeolodian Type = iota
	Doptian   Type = iota
	Aeraphian Type = iota
	Zacrian   Type = iota

	// 7.3 Ionarian scale
	Ionarian Type = iota
	Dynian   Type = iota
	Zydian   Type = iota
	Zathian  Type = iota
	Radian   Type = iota
	Stonian  Type = iota
	Syptian  Type = iota

	// 7.4 Aeolacrian scale
	Aeolacrian Type = iota
	Zythian    Type = iota
	Dyrian     Type = iota
	Koptian    Type = iota
	Thocrian   Type = iota
	Aeolanian  Type = iota
	Danian     Type = iota

	// 7.5 Zogian scale
	Zogian  Type = iota
	Epyrian Type = iota
	Lycrian Type = iota
	Daptian Type = iota
	Kygian  Type = iota
	Mocrian Type = iota
	Zynian  Type = iota

	// 7.6 Phrolian scale
	Phrolian Type = iota
	Ionagian Type = iota
	Aeodian  Type = iota
	Kycrian  Type = iota
	Epygian  Type = iota
	Zaptian  Type = iota
	Kagian   Type = iota

	// 7.7 Soptian scale
	Soptian   Type = iota
	Ionyptian Type = iota
	Gyrian    Type = iota
	Zalian    Type = iota
	Stolian   Type = iota
	Bylian    Type = iota
	Zothian   Type = iota

	// 7.8 Thonian scale
	Thonian    Type = iota
	Phrorian   Type = iota
	Stadian    Type = iota
	Thodian    Type = iota
	Dogian     Type = iota
	Mixopyrian Type = iota
	Garian     Type = iota

	// 7.9 Epathian scale
	Epathian Type = iota
	Mythian  Type = iota
	Sogian   Type = iota
	Gogian   Type = iota
	Rothian  Type = iota
	Katarian Type = iota
	Stylian  Type = iota

	// 7.10 Stathian scale
	Stathian    Type = iota
	Mixonyphian Type = iota
	Magian      Type = iota
	Dadian      Type = iota
	Aeolylian   Type = iota
	Gycrian     Type = iota
	Pyrian      Type = iota

	// 7.11 Epogian scale
	Epogian   Type = iota
	Lanian    Type = iota
	Paptian   Type = iota
	Ionacrian Type = iota
	Gathian   Type = iota
	Ionyphian Type = iota
	Phrynian  Type = iota

	// 7.12 Ionycrian scale
	Ionycrian Type = iota
	Phradian  Type = iota
	Aeolorian Type = iota
	Gonian    Type = iota
	Dalian    Type = iota
	Dygian    Type = iota
	Zolian    Type = iota

	// 7.13 Aerathian scale
	Aerathian Type = iota
	Sarian    Type = iota
	Zoptian   Type = iota
	Aeracrian Type = iota
	Byptian   Type = iota
	Darian    Type = iota
	Lonian    Type = iota

	// 7.14 Aeopian scale
	Aeopian  Type = iota
	Rygian   Type = iota
	Epynian  Type = iota
	Ionogian Type = iota
	Kydian   Type = iota
	Gaptian  Type = iota
	Tharian  Type = iota

	// 7.15 Epycrian scale
	Epycrian Type = iota
	Gocrian  Type = iota
	Katolian Type = iota
	Thoptian Type = iota
	Bagian   Type = iota
	Kyrian   Type = iota
	Sonian   Type = iota

	// 7.16 Parian scale
	Parian    Type = iota
	Ionaptian Type = iota
	Thylian   Type = iota
	Lolian    Type = iota
	Thagian   Type = iota
	Kolian    Type = iota
	Dycrian   Type = iota

	// 7.17 Stygian scale
	Stygian   Type = iota
	Aeolygian Type = iota
	Aerogian  Type = iota
	Dacrian   Type = iota
	Baptian   Type = iota
	Dagian    Type = iota
	Aeolydian Type = iota

	// 7.18 Stythian scale
	Stythian Type = iota
	Kothian  Type = iota
	Pygian   Type = iota
	Rodian   Type = iota
	Sorian   Type = iota
	Monian   Type = iota
	Thalian  Type = iota

	// 7.19 Zorian scale
	Zorian   Type = iota
	Aeragian Type = iota
	Epothian Type = iota
	Salian   Type = iota
	Lyptian  Type = iota
	Katonian Type = iota
	Gyphian  Type = iota

	// 7.20 Thacrian scale
	Thacrian   Type = iota
	Dodian     Type = iota
	Aeolyptian Type = iota
	Aeolonian  Type = iota
	Aeradian   Type = iota
	Aeolagian  Type = iota
	Zyrian     Type = iota

	// 7.21 Aeolathian scale
	Aeolathian Type = iota
	Bythian    Type = iota
	Padian     Type = iota
	Rolian     Type = iota
	Pydian     Type = iota
	Thygian    Type = iota
	Katalian   Type = iota

	// 7.22 Saptian scale
	Saptian  Type = iota
	Aerodian Type = iota
	Macrian  Type = iota
	Rogian   Type = iota
	Boptian  Type = iota
	Stogian  Type = iota
	Thynian  Type = iota

	// 7.23 Thycrian scale
	Thycrian  Type = iota
	Aeoladian Type = iota
	Dylian    Type = iota
	Eponian   Type = iota
	Katygian  Type = iota
	Starian   Type = iota
	Phrathian Type = iota

	// 7.24 Stalian scale
	Stalian    Type = iota
	Stoptian   Type = iota
	Zygian     Type = iota
	Kataptian  Type = iota
	Aeolaptian Type = iota
	Pothian    Type = iota
	Bygian     Type = iota

	// 7.25 Morian scale
	Morian   Type = iota
	Rycrian  Type = iota
	Ronian   Type = iota
	Stycrian Type = iota
	Katorian Type = iota
	Epythian Type = iota
	Kaptian  Type = iota

	// 7.26 Phraptian scale
	Phraptian Type = iota
	Bacrian   Type = iota
	Katythian Type = iota
	Madian    Type = iota
	Aerygian  Type = iota
	Pylian    Type = iota
	Ionathian Type = iota

	// 7.27 Katocrian scale
	Katocrian Type = iota
	Phryptian Type = iota
	Katynian  Type = iota
	Aerycrian Type = iota
	Ganian    Type = iota
	Eparian   Type = iota
	Lyrian    Type = iota

	// 7.28 Ionopian scale
	Ionopian  Type = iota
	Aeologian Type = iota
	Zadian    Type = iota
	Sygian    Type = iota
	Phralian  Type = iota
	Phrogian  Type = iota
	Rathian   Type = iota

	// 7.29 Rythian scale
	Rythian    Type = iota
	Donian     Type = iota
	Aeoloptian Type = iota
	Panian     Type = iota
	Lodian     Type = iota
	Solian     Type = iota
	Ionolian   Type = iota

	// 7.30 Laptian scale
	Laptian   Type = iota
	Lygian    Type = iota
	Logian    Type = iota
	Lalian    Type = iota
	Sothian   Type = iota
	Phrocrian Type = iota
	Thogian   Type = iota

	// 7.31 Katathian scale
	Katathian   Type = iota
	Modian      Type = iota
	Barian      Type = iota
	Mixolocrian Type = iota
	Sagian      Type = iota
	Aeolothian  Type = iota
	Socrian     Type = iota

	// 7.32 Tholian scale
	Tholian   Type = iota
	Ralian    Type = iota
	Syrian    Type = iota
	Stodian   Type = iota
	Ionocrian Type = iota
	Zycrian   Type = iota
	Ionygian  Type = iota

	// 7.33 Zarian scale
	Zarian    Type = iota
	Phrythian Type = iota
	Rorian    Type = iota
	Bolian    Type = iota
	Bothian   Type = iota
	Katadian  Type = iota
	Kodian    Type = iota

	// 7.34 Ranian scale
	Ranian    Type = iota
	Ladian    Type = iota
	Podian    Type = iota
	Ionothian Type = iota
	Kanian    Type = iota
	Zylian    Type = iota
	Zodian    Type = iota

	// 7.35 Golian scale
	Golian     Type = iota
	Dyptian    Type = iota
	Ryphian    Type = iota
	Gylian     Type = iota
	Aeolycrian Type = iota
	Pynian     Type = iota
	Zanian     Type = iota

	// 7.36 Palian scale
	Palian   Type = iota
	Stothian Type = iota
	Aerorian Type = iota
	Katagian Type = iota
	Phronian Type = iota
	Banian   Type = iota
	Aeronian Type = iota

	// 7.37 Loptian scale
	Loptian  Type = iota
	Ionodian Type = iota
	Bogian   Type = iota
	Mogian   Type = iota
	Docrian  Type = iota
	Epadian  Type = iota
	Aerynian Type = iota

	// 7.38 Bycrian scale
	Bycrian   Type = iota
	Pathian   Type = iota
	Mydian    Type = iota
	Thyptian  Type = iota
	Phrothian Type = iota
	Katycrian Type = iota
	Ionalian  Type = iota

	// 7.39 Dolian scale
	Dolian     Type = iota
	Porian     Type = iota
	Aerylian   Type = iota
	Zagian     Type = iota
	Lagian     Type = iota
	Tyrian     Type = iota
	Mixonorian Type = iota

	// 7.40 Pagian scale
	Pagian     Type = iota
	Aeolythian Type = iota
	Molian     Type = iota
	Staptian   Type = iota
	Mothian    Type = iota
	Aeranian   Type = iota
	Ragian     Type = iota

	// 7.41 Larian scale
	Larian      Type = iota
	Lythian     Type = iota
	Stydian     Type = iota
	Lorian      Type = iota
	Ionadian    Type = iota
	Bocrian     Type = iota
	Mixolythian Type = iota

	// 7.42 Thadian scale
	Thadian   Type = iota
	Sanian    Type = iota
	Ionydian  Type = iota
	Epydian   Type = iota
	Katydian  Type = iota
	Mathian   Type = iota
	Aeryptian Type = iota

	// 7.43 Pythian scale
	Pythian  Type = iota
	Katylian Type = iota
	Bydian   Type = iota
	Bynian   Type = iota
	Galian   Type = iota
	Zonian   Type = iota
	Myrian   Type = iota

	// 7.44 Katogian scale
	Katogian Type = iota
	Stacrian Type = iota
	Styrian  Type = iota
	Ionyrian Type = iota
	Phrodian Type = iota
	Pycrian  Type = iota
	Gyptian  Type = iota

	// 7.45 Katacrian scale
	Katacrian Type = iota
	Sodian    Type = iota
	Bathian   Type = iota
	Mylian    Type = iota
	Godian    Type = iota
	Thorian   Type = iota
	Zocrian   Type = iota

	// 7.46 Stanian scale
	Stanian   Type = iota
	Epanian   Type = iota
	Konian    Type = iota
	Stocrian  Type = iota
	Kalian    Type = iota
	Phroptian Type = iota
	Dydian    Type = iota

	// 7.47 Katyptian scale
	Katyptian Type = iota
	Epodian   Type = iota
	Mygian    Type = iota
	Pacrian   Type = iota
	Aerocrian Type = iota
	Aeolarian Type = iota
	Kythian   Type = iota

	// 7.48 Bonian scale
	Bonian   Type = iota
	Badian   Type = iota
	Katodian Type = iota
	Sadian   Type = iota
	Dothian  Type = iota
	Moptian  Type = iota
	Aeryrian Type = iota

	// 7.49 Epagian scale
	Epagian  Type = iota
	Raptian  Type = iota
	Epolian  Type = iota
	Sythian  Type = iota
	Sydian   Type = iota
	Epocrian Type = iota
	Kylian   Type = iota

	// 7.50 Gacrian scale
	Gacrian    Type = iota
	Borian     Type = iota
	Sycrian    Type = iota
	Gadian     Type = iota
	Aeolocrian Type = iota
	Mixodorian Type = iota
	Dathian    Type = iota

	// 7.51 Katoptian scale
	Katoptian Type = iota
	Ponian    Type = iota
	Kadian    Type = iota
	Gynian    Type = iota
	Thyphian  Type = iota
	Polian    Type = iota
	Thanian   Type = iota

	// 7.52 Epacrian scale
	Epacrian Type = iota
	Sathian  Type = iota
	Lathian  Type = iota
	Aeralian Type = iota
	Kynian   Type = iota
	Stynian  Type = iota
	Epyphian Type = iota

	// 7.53 Pogian scale
	Pogian    Type = iota
	Aeraptian Type = iota
	Epylian   Type = iota
	Gamian    Type = iota
	Kathian   Type = iota
	Lylian    Type = iota
	Epalian   Type = iota

	// 7.54 Eporian scale
	Eporian  Type = iota
	Rylian   Type = iota
	Epaptian Type = iota
	Byrian   Type = iota
	Katanian Type = iota
	Katyrian Type = iota
	Rynian   Type = iota

	// 7.55 Korian scale
	Korian   Type = iota
	Lynian   Type = iota
	Malian   Type = iota
	Synian   Type = iota
	Phragian Type = iota
	Manian   Type = iota
	Marian   Type = iota

	// 7.56 Mycrian scale
	Mycrian   Type = iota
	Ionorian  Type = iota
	Phrydian  Type = iota
	Zyptian   Type = iota
	Katothian Type = iota
	Phrylian  Type = iota
	Kocrian   Type = iota

	// 7.57 Ionanian scale
	Ionanian  Type = iota
	Aerothian Type = iota
	Stagian   Type = iota
	Lothian   Type = iota
	Phrycrian Type = iota
	Kyptian   Type = iota
	Ionylian  Type = iota

	// 7.58 Gydian scale
	Gydian   Type = iota
	Kogian   Type = iota
	Rarian   Type = iota
	Aerolian Type = iota
	Karian   Type = iota
	Myptian  Type = iota
	Rydian   Type = iota

	// 7.59 Aeolynian scale
	Aeolynian Type = iota
	Aeroptian Type = iota
	Phryrian  Type = iota
	Gothian   Type = iota
	Storian   Type = iota
	Pyptian   Type = iota
	Thydian   Type = iota

	// 8.1 Aerycryllic scale
	Aerycryllic Type = iota
	Gadyllic    Type = iota
	Solyllic    Type = iota
	Zylyllic    Type = iota
	Mixodyllic  Type = iota
	Soryllic    Type = iota
	Godyllic    Type = iota
	Epiphyllic  Type = iota

	// 8.2 Pynyllic scale
	Pynyllic  Type = iota
	Bocryllic Type = iota
	Kogyllic  Type = iota
	Raryllic  Type = iota
	Zycryllic Type = iota
	Mycryllic Type = iota
	Laptyllic Type = iota
	Pylyllic  Type = iota

	// 8.3 Pothyllic scale
	Pothyllic   Type = iota
	Phronyllic  Type = iota
	Stynyllic   Type = iota
	Rathyllic   Type = iota
	Aeryptyllic Type = iota
	Zydyllic    Type = iota
	Katolyllic  Type = iota
	Rythyllic   Type = iota

	// 8.4 Locryllic scale
	Locryllic   Type = iota
	Bylyllic    Type = iota
	Sogyllic    Type = iota
	Ionycryllic Type = iota
	Koptyllic   Type = iota
	Epyryllic   Type = iota
	Soptyllic   Type = iota
	Aeolylyllic Type = iota

	// 8.5 Aeracryllic scale
	Aeracryllic Type = iota
	Epygyllic   Type = iota
	Thonyllic   Type = iota
	Lanyllic    Type = iota
	Phrynyllic  Type = iota
	Lycryllic   Type = iota
	Ionyptyllic Type = iota
	Epathyllic  Type = iota

	// 8.6 Dydyllic scale
	Dydyllic  Type = iota
	Thogyllic Type = iota
	Rygyllic  Type = iota
	Bycryllic Type = iota
	Zacryllic Type = iota
	Panyllic  Type = iota
	Dyryllic  Type = iota
	Zathyllic Type = iota

	// 8.7 Dagyllic scale
	Dagyllic    Type = iota
	Katalyllic  Type = iota
	Katoryllic  Type = iota
	Dodyllic    Type = iota
	Zogyllic    Type = iota
	Madyllic    Type = iota
	Dycryllic   Type = iota
	Aeologyllic Type = iota

	// 8.8 Sydyllic scale
	Sydyllic   Type = iota
	Katogyllic Type = iota
	Zygyllic   Type = iota
	Aeralyllic Type = iota

	// 8.9 Bacryllic scale
	Bacryllic  Type = iota
	Aerygyllic Type = iota
	Dathyllic  Type = iota
	Boptyllic  Type = iota
	Bagyllic   Type = iota
	Mathyllic  Type = iota
	Styptyllic Type = iota
	Zolyllic   Type = iota

	// 8.10 Rocryllic scale
	Rocryllic  Type = iota
	Zyryllic   Type = iota
	Sagyllic   Type = iota
	Epinyllic  Type = iota
	Katagyllic Type = iota
	Ragyllic   Type = iota
	Gothyllic  Type = iota
	Lythyllic  Type = iota

	// 8.11 Ionocryllic scale
	Ionocryllic Type = iota
	Gocryllic   Type = iota
	Epiryllic   Type = iota
	Aeradyllic  Type = iota
	Staptyllic  Type = iota
	Danyllic    Type = iota
	Goptyllic   Type = iota
	Epocryllic  Type = iota

	// 8.12 Ionoptyllic scale
	Ionoptyllic Type = iota
	Aeoloryllic Type = iota
	Thydyllic   Type = iota
	Gycryllic   Type = iota
	Lyryllic    Type = iota
	Mogyllic    Type = iota
	Katodyllic  Type = iota
	Moptyllic   Type = iota

	// 8.13 Dolyllic scale
	Dolyllic    Type = iota
	Moryllic    Type = iota
	Bydyllic    Type = iota
	Pocryllic   Type = iota
	Phracryllic Type = iota
	Gyryllic    Type = iota
	Phrygyllic  Type = iota
	Dogyllic    Type = iota

	// 8.14 Thagyllic scale
	Thagyllic   Type = iota
	Thoptyllic  Type = iota
	Phraptyllic Type = iota
	Gylyllic    Type = iota
	Phralyllic  Type = iota
	Dygyllic    Type = iota
	Ronyllic    Type = iota
	Epogyllic   Type = iota

	// 8.15 Aeoladyllic scale
	Aeoladyllic Type = iota
	Kocryllic   Type = iota
	Lodyllic    Type = iota
	Bynyllic    Type = iota
	Kydyllic    Type = iota
	Bygyllic    Type = iota
	Phryptyllic Type = iota
	Ionayllic   Type = iota

	// 8.16 Phroryllic scale
	Phroryllic  Type = iota
	Thyphyllic  Type = iota
	Poptyllic   Type = iota
	Mixonyllic  Type = iota
	Paptyllic   Type = iota
	Storyllic   Type = iota
	Phrycryllic Type = iota
	Palyllic    Type = iota

	// 8.17 Phranyllic scale
	Phranyllic  Type = iota
	Stydyllic   Type = iota
	Zadyllic    Type = iota
	Zalyllic    Type = iota
	Zocryllic   Type = iota
	Katocryllic Type = iota
	Aerathyllic Type = iota
	Stoptyllic  Type = iota

	// 8.18 Lydyllic scale
	Lydyllic   Type = iota
	Radyllic   Type = iota
	Stagyllic  Type = iota
	Ionoryllic Type = iota
	Phrodyllic Type = iota
	Aeragyllic Type = iota
	Banyllic   Type = iota
	Epothyllic Type = iota

	// 8.19 Zoryllic scale
	Zoryllic   Type = iota
	Phrolyllic Type = iota
	Kolyllic   Type = iota
	Thodyllic  Type = iota
	Socryllic  Type = iota
	Aeolyllic  Type = iota
	Zythyllic  Type = iota
	Aeoryllic  Type = iota

	// 8.20 Mixolydyllic
	Mixolydyllic  Type = iota
	Mixonyphyllic Type = iota
	Aeolanyllic   Type = iota
	Thocryllic    Type = iota
	Kygyllic      Type = iota
	Ionagyllic    Type = iota
	Gogyllic      Type = iota
	Phradyllic    Type = iota

	// 8.21 Ioniptyllic scale
	Ioniptyllic  Type = iota
	Kycryllic    Type = iota
	Aeolaptyllic Type = iota
	Rodyllic     Type = iota
	Ionathyllic  Type = iota
	Pythyllic    Type = iota
	Zonyllic     Type = iota
	Ryryllic     Type = iota

	// 8.22 Aeolothyllic scale
	Aeolothyllic Type = iota
	Ionyryllic   Type = iota
	Rydyllic     Type = iota
	Gonyllic     Type = iota
	Rolyllic     Type = iota
	Katydyllic   Type = iota
	Zyptyllic    Type = iota
	Modyllic     Type = iota

	// 8.23 Maptyllic scale
	Maptyllic   Type = iota
	Aeraptyllic Type = iota
	Katadyllic  Type = iota
	Magyllic    Type = iota
	Phrylyllic  Type = iota
	Epigyllic   Type = iota
	Molyllic    Type = iota
	Ponyllic    Type = iota

	// 8.24 Thyptyllic scale
	Thyptyllic  Type = iota
	Ionogyllic  Type = iota
	Aeolaryllic Type = iota
	Katygyllic  Type = iota
	Ganyllic    Type = iota
	Kyptyllic   Type = iota
	Salyllic    Type = iota
	Sanyllic    Type = iota

	// 8.25 Doptyllic scale
	Doptyllic     Type = iota
	Ionilyllic    Type = iota
	Manyllic      Type = iota
	Polyllic      Type = iota
	Stanyllic     Type = iota
	Mixotharyllic Type = iota
	Eporyllic     Type = iota
	Aerynyllic    Type = iota

	// 8.26 Lonyllic scale
	Lonyllic    Type = iota
	Sathyllic   Type = iota
	Layllic     Type = iota
	Saryllic    Type = iota
	Thacryllic  Type = iota
	Aeolynyllic Type = iota
	Thadyllic   Type = iota
	Lynyllic    Type = iota

	// 8.27 Aeolathyllic scale
	Aeolathyllic Type = iota
	Aeolocryllic Type = iota
	Phroptyllic  Type = iota
	Kodyllic     Type = iota
	Epaptyllic   Type = iota
	Ionoyllic    Type = iota
	Gyptyllic    Type = iota
	Aerythyllic  Type = iota

	// 8.28 Zagyllic scale
	Zagyllic    Type = iota
	Epacryllic  Type = iota
	Thorcryllic Type = iota
	Loptyllic   Type = iota
	Katylyllic  Type = iota
	Malyllic    Type = iota
	Mydyllic    Type = iota
	Thycryllic  Type = iota

	// 8.29 Gythyllic scale
	Gythyllic   Type = iota
	Pyryllic    Type = iota
	Rycryllic   Type = iota
	Phrathyllic Type = iota
	Badyllic    Type = iota
	Phrocryllic Type = iota
	Staryllic   Type = iota
	Zothyllic   Type = iota

	// 8.30 Tharyllic scale
	Tharyllic    Type = iota
	Sylyllic     Type = iota
	Lothyllic    Type = iota
	Daryllic     Type = iota
	Monyllic     Type = iota
	Styryllic    Type = iota
	Aeolacryllic Type = iota
	Raptyllic    Type = iota

	// 8.31 Kataryllic scale
	Kataryllic  Type = iota
	Aerocryllic Type = iota
	Zanyllic    Type = iota
	Aeolonyllic Type = iota
	Aeonyllic   Type = iota
	Kyryllic    Type = iota
	Sythyllic   Type = iota
	Katycryllic Type = iota

	// 8.32 Stogyllic scale
	Stogyllic   Type = iota
	Ionidyllic  Type = iota
	Stonyllic   Type = iota
	Stalyllic   Type = iota
	Poryllic    Type = iota
	Mocryllic   Type = iota
	Aeolyryllic Type = iota
	Baryllic    Type = iota

	// 8.33 Dalyllic scale
	Dalyllic     Type = iota
	Ionyphyllic  Type = iota
	Zaptyllic    Type = iota
	Garyllic     Type = iota
	Gathyllic    Type = iota
	Mixopyryllic Type = iota
	Ionacryllic  Type = iota
	Stylyllic    Type = iota

	// 8.34 Stycryllic scale
	Stycryllic  Type = iota
	Ionothyllic Type = iota
	Mythyllic   Type = iota
	Aerylyllic  Type = iota
	Bonyllic    Type = iota
	Tholyllic   Type = iota
	Katyryllic  Type = iota
	Sadyllic    Type = iota

	// 8.35 Stolyllic scale
	Stolyllic Type = iota
	Logyllic  Type = iota
	Dacryllic Type = iota
	Thynyllic Type = iota
	Gydyllic  Type = iota
	Eparyllic Type = iota
	Dynyllic  Type = iota
	Ionyllic  Type = iota

	// 8.36 Zaryllic scale
	Zaryllic    Type = iota
	Dythyllic   Type = iota
	Ionaryllic  Type = iota
	Laryllic    Type = iota
	Kataptyllic Type = iota
	Sonyllic    Type = iota
	Pathyllic   Type = iota
	Loryllic    Type = iota

	// 8.37 Aeronyllic scale
	Aeronyllic Type = iota
	Pycryllic  Type = iota
	Mygyllic   Type = iota
	Lylyllic   Type = iota
	Daptyllic  Type = iota
	Ioninyllic Type = iota
	Epaphyllic Type = iota
	Lolyllic   Type = iota

	// 8.38 Stacryllic scale
	Stacryllic Type = iota
	Doryllic   Type = iota
	Kadyllic   Type = iota
	Rynyllic   Type = iota
	Aerogyllic Type = iota
	Rothyllic  Type = iota
	Kagyllic   Type = iota
	Stathyllic Type = iota

	// 8.39 Thyryllic scale
	Thyryllic   Type = iota
	Gygyllic    Type = iota
	Sodyllic    Type = iota
	Goryllic    Type = iota
	Bothyllic   Type = iota
	Gynyllic    Type = iota
	Ionaptyllic Type = iota
	Phryryllic  Type = iota

	// 8.40 Racryllic scale
	Racryllic   Type = iota
	Epicryllic  Type = iota
	Stygyllic   Type = iota
	Syryllic    Type = iota
	Stythyllic  Type = iota
	Aerothyllic Type = iota
	Mixoryllic  Type = iota
	Thanyllic   Type = iota

	// 8.41 Roryllic
	Epotyllic Type = iota
	Epidyllic Type = iota
	Kaptyllic Type = iota

	// 8.42 MajorDiminished scale
	MajorDiminished Type = iota
	MinorDiminished Type = iota

	// 9.1 Aerycrygic scale
	Aerycrygic Type = iota
	Gadygic    Type = iota
	Solygic    Type = iota
	Zylygic    Type = iota
	Garygic    Type = iota
	Sorygic    Type = iota
	Godygic    Type = iota
	Epithygic  Type = iota
	Ionoptygic Type = iota

	// 9.2 Kalygic scale
	Kalygic    Type = iota
	Ionodygic  Type = iota
	Bythygic   Type = iota
	Epygic     Type = iota
	Marygic    Type = iota
	Gaptygic   Type = iota
	Aeroptygic Type = iota
	Mylygic    Type = iota
	Galygic    Type = iota

	// 9.3 Mixolydygic scale
	Mixolydygic Type = iota
	Ionycrygic  Type = iota
	Zoptygic    Type = iota
	Phrygygic   Type = iota
	Locrygic    Type = iota
	Gonygic     Type = iota
	Aeracrygic  Type = iota
	Aerathygic  Type = iota
	Dorygic     Type = iota

	// 9.4 Dycrygic scale
	Dycrygic Type = iota
	Aeolygic Type = iota
	Dydygic  Type = iota
	Tholygic Type = iota
	Rynygic  Type = iota
	Bycrygic Type = iota
	Zacrygic Type = iota
	Panygic  Type = iota
	Dyrygic  Type = iota

	// 9.5 Loptygic scale
	Loptygic  Type = iota
	Katylygic Type = iota
	Phradygic Type = iota
	Mixodygic Type = iota
	Katalygic Type = iota
	Katorygic Type = iota
	Dogygic   Type = iota
	Zodygic   Type = iota
	Madygic   Type = iota

	// 9.6 Bagygic scale
	Bagygic   Type = iota
	Mathygic  Type = iota
	Styptygic Type = iota
	Zolygic   Type = iota
	Sydygic   Type = iota
	Katygic   Type = iota
	Zyphygic  Type = iota
	Aeralygic Type = iota
	Ryptygic  Type = iota

	// 9.7 Apinygic scale
	Apinygic  Type = iota
	Katagygic Type = iota
	Radygic   Type = iota
	Gothygic  Type = iota
	Lythygic  Type = iota
	Bacrygic  Type = iota
	Aerygic   Type = iota
	Dathygic  Type = iota
	Boptygic  Type = iota

	// 9.8 Epyrygic scale
	Epyrygic  Type = iota
	Aeradygic Type = iota
	Staptygic Type = iota
	Danygic   Type = iota
	Goptygic  Type = iota
	Epocrygic Type = iota
	Rocrygic  Type = iota
	Zyrygic   Type = iota
	Sadygic   Type = iota

	// 9.9 Aeolorygic scale
	Aeolorygic Type = iota
	Thydygic   Type = iota
	Gycrygic   Type = iota
	Lyrygic    Type = iota
	Modygic    Type = iota
	Katodygic  Type = iota
	Moptygic   Type = iota
	Ionocrygic Type = iota
	Gocrygic   Type = iota

	// 9.10 Manygic scale
	Manygic    Type = iota
	Polygic    Type = iota
	Stanygic   Type = iota
	Thaptygic  Type = iota
	Eporygic   Type = iota
	Aerynygic  Type = iota
	Thyptygic  Type = iota
	Ionogygic  Type = iota
	Aeolarygic Type = iota

	// 9.11 Sathygic scale
	Sathygic   Type = iota
	Ladygic    Type = iota
	Sarygic    Type = iota
	Thacrygic  Type = iota
	Aeolynygic Type = iota
	Thadygic   Type = iota
	Lynygic    Type = iota
	Doptygic   Type = iota
	Ionilygic  Type = iota

	// 9.12 Phrygic scale
	Phrygic   Type = iota
	Aeranygic Type = iota
	Dothygic  Type = iota
	Lydygic   Type = iota
	Stadygic  Type = iota
	Byptygic  Type = iota
	Stodygic  Type = iota
	Zynygic   Type = iota
	Lonygic   Type = iota

	// 9.13 Zothygic scale
	Zothygic    Type = iota
	Aeolathygic Type = iota
	Aeolocrygic Type = iota
	Phroptygic  Type = iota
	Kodygic     Type = iota
	Eparygic    Type = iota
	Ionygic     Type = iota
	Gyptygic    Type = iota
	Aerythygic  Type = iota

	// 9.14 Aeolacrygic scale
	Aeolacrygic Type = iota
	Raptygic    Type = iota
	Gythygic    Type = iota
	Pyrygic     Type = iota
	Rycrygic    Type = iota
	Phrathygic  Type = iota
	Badygic     Type = iota
	Phrocrygic  Type = iota
	Starygic    Type = iota

	// 9.15 Kyrygic scale
	Kyrygic    Type = iota
	Sythygic   Type = iota
	Katycrygic Type = iota
	Tharygic   Type = iota
	Sylygic    Type = iota
	Lothygic   Type = iota
	Darygic    Type = iota
	Monygic    Type = iota
	Styrygic   Type = iota

	// 9.16 Porygic scale
	Porygic    Type = iota
	Mocrygic   Type = iota
	Aeolyrygic Type = iota
	Barygic    Type = iota
	Katarygic  Type = iota
	Aerocrygic Type = iota
	Zanygic    Type = iota
	Aeolonygic Type = iota
	Aeolanygic Type = iota

	// 9.17 Kaptygic scale
	Kaptygic  Type = iota
	Sacrygic  Type = iota
	Padygic   Type = iota
	Epilygic  Type = iota
	Kynygic   Type = iota
	Stophygic Type = iota
	Ionidygic Type = iota
	Stonygic  Type = iota
	Stalygic  Type = iota

	// 9.18 Koptygic scale
	Koptygic Type = iota
	Raphygic Type = iota
	Zycrygic Type = iota
	Mycrygic Type = iota
	Laptygic Type = iota
	Pylygic  Type = iota
	Rodygic  Type = iota
	Epolygic Type = iota
	Epidygic Type = iota

	// 9.19 Phronygic scale
	Phronygic Type = iota
	Stynygic  Type = iota
	Zydygic   Type = iota

	// 10.1 Aerycryllian scale
	Aerycryllian Type = iota
	Gadyllian    Type = iota
	Solyllian    Type = iota
	Zyphyllian   Type = iota
	Garyllian    Type = iota
	Soryllian    Type = iota
	Godyllian    Type = iota
	Epityllian   Type = iota
	Ionyllian    Type = iota
	Aeoryllian   Type = iota

	// 10.2 Katoryllian scale
	Katoryllian Type = iota
	Dodyllian   Type = iota
	Zogyllian   Type = iota
	Madyllian   Type = iota
	Dycryllian  Type = iota
	Aeogyllian  Type = iota
	Dydyllian   Type = iota
	Thogyllian  Type = iota
	Rygyllian   Type = iota
	Bathyllian  Type = iota

	// 10.3 Sydyllian scale
	Sydyllian    Type = iota
	Katogyllian  Type = iota
	Mixodyllian  Type = iota
	Aeradyllian  Type = iota
	Ryptyllian   Type = iota
	Loptyllian   Type = iota
	Kataphyllian Type = iota
	Phradyllian  Type = iota
	Dagyllian    Type = iota
	Katyllian    Type = iota

	// 10.4 Gothyllian scale
	Gothyllian  Type = iota
	Lythyllian  Type = iota
	Bacryllian  Type = iota
	Aerygyllian Type = iota
	Dathyllian  Type = iota
	Boptyllian  Type = iota
	Bagyllian   Type = iota
	Mathyllian  Type = iota
	Styptyllian Type = iota
	Zolyllian   Type = iota

	// 10.5 Staptyllian scale
	Staptyllian Type = iota
	Danyllian   Type = iota
	Goptyllian  Type = iota
	Epocryllian Type = iota
	Rocryllian  Type = iota
	Zyryllian   Type = iota
	Sagyllian   Type = iota
	Epinyllian  Type = iota
	Katagyllian Type = iota
	Ragyllian   Type = iota

	// 10.6 Thydyllian
	Thydyllian  Type = iota
	Epiryllian  Type = iota
	Lyryllian   Type = iota
	Mogyllian   Type = iota
	Katodyllian Type = iota

	// 11.1 Aerycratic scale
	Aerycratic Type = iota
	Monatic    Type = iota
	Solatic    Type = iota
	Zylatic    Type = iota
	Mixolatic  Type = iota
	Soratic    Type = iota
	Godatic    Type = iota
	Eptatic    Type = iota
	Ionatic    Type = iota
	Aeolatic   Type = iota
	Thydatic   Type = iota

	// 12.1 Chromatic scale
	Chromatic Type = iota
)

// AllTypes returns list of mode types
func AllTypes() Types {
	return allTypes
}
