package modetype

var mapModeToString = map[Type]string{
	Invalid: "Invalid",

	// 3.1
	Minoric: "Minoric",

	// 4.1 Thaptic scale
	Thaptic:  "Thaptic",
	Lothic:   "Lothic",
	Phratic:  "Phratic",
	Aerathic: "Aerathic",

	// 4.2 Epathic scale
	Epathic: "Epathic",
	Mynic:   "Mynic",
	Rothic:  "Rothic",
	Eporic:  "Eporic",

	// 4.3 Zyphic scale
	Zyphic: "Zyphic",
	Epogic: "Epogic",
	Lanic:  "Lanic",
	Pyrric: "Pyrric",

	// 4.4 Aeoloric scale
	Aeoloric: "Aeoloric",
	Gonic:    "Gonic",
	Dalic:    "Dalic",
	Dygic:    "Dygic",

	// 4.5 Daric scale
	Daric:   "Daric",
	Lonic:   "Lonic",
	Phradic: "Phradic",
	Bolic:   "Bolic",

	// 4.6 Saric scale
	Saric:    "Saric",
	Zoptic:   "Zoptic",
	Aeraphic: "Aeraphic",
	Byptic:   "Byptic",

	// 4.7 Aeolic scale
	Aeolic:    "Aeolic",
	Koptic:    "Koptic",
	Mixolyric: "Mixolyric",
	Lydic:     "Lydic",

	// 4.8 Saric scale
	Stathic: "Stathic",
	Dadic:   "Dadic",

	// 4.9 Phrynic scale
	Phrynic: "Phrynic",

	// 5.1 Epathitonic scale
	Epathitonic: "Epathitonic",
	Mynitonic:   "Mynitonic",
	Rocritonic:  "Rocritonic",
	Pentatonic:  "Pentatonic",
	Thaptitonic: "Thaptitonic",

	// 5.2 Magitonic scale
	Magitonic:      "Magitonic",
	Daditonic:      "Daditonic",
	Aeolyphritonic: "Aeolyphritonic",
	Gycritonic:     "Gycritonic",
	Pyritonic:      "Pyritonic",

	// 5.3 Gathitonic scale
	Gathitonic:  "Gathitonic",
	Ionitonic:   "Ionitonic",
	Phrynitonic: "Phrynitonic",
	Stathitonic: "Stathitonic",
	Thalitonic:  "Thalitonic",

	// 5.4 Zolitonic scale
	Zolitonic:    "Zolitonic",
	Epogitonic:   "Epogitonic",
	Lanitonic:    "Lanitonic",
	Paptitonic:   "Paptitonic",
	Ionacritonic: "Ionacritonic",

	// 5.5 Phraditonic scale
	Phraditonic:  "Phraditonic",
	Aeoloritonic: "Aeoloritonic",
	Gonitonic:    "Gonitonic",
	Dalitonic:    "Dalitonic",
	Dygitonic:    "Dygitonic",

	// 5.6 Aeracritonic scale
	Aeracritonic: "Aeracritonic",
	Byptitonic:   "Byptitonic",
	Daritonic:    "Daritonic",
	Lonitonic:    "Lonitonic",
	Ionycritonic: "Ionycritonic",

	// 5.7 Lothitonic scale
	Lothitonic:   "Lothitonic",
	Phratonic:    "Phratonic",
	Aerathitonic: "Aerathitonic",
	Saritonic:    "Saritonic",
	Zoptitonic:   "Zoptitonic",

	// 5.8 Dolitonic scale
	Dolitonic:   "Dolitonic",
	Poritonic:   "Poritonic",
	Aerylitonic: "Aerylitonic",
	Zagitonic:   "Zagitonic",
	Lagitonic:   "Lagitonic",

	// 5.9 Molitonic scale
	Molitonic:   "Molitonic",
	Staptitonic: "Staptitonic",
	Mothitonic:  "Mothitonic",
	Aeritonic:   "Aeritonic",
	Ragitonic:   "Ragitonic",

	// 5.10 Ionaditonic scale
	Ionaditonic:   "Ionaditonic",
	Bocritonic:    "Bocritonic",
	Gythitonic:    "Gythitonic",
	Pagitonic:     "Pagitonic",
	Aeolythitonic: "Aeolythitonic",

	// 5.11 Zacritonic scale
	Zacritonic:  "Zacritonic",
	Laritonic:   "Laritonic",
	Thacritonic: "Thacritonic",
	Styditonic:  "Styditonic",
	Loritonic:   "Loritonic",

	// 5.12 Aeolyritonic scale
	Aeolyritonic: "Aeolyritonic",
	Goritonic:    "Goritonic",
	Aeoloditonic: "Aeoloditonic",
	Doptitonic:   "Doptitonic",
	Aeraphitonic: "Aeraphitonic",

	// 5.13 Zathitonic scale
	Zathitonic:   "Zathitonic",
	Raditonic:    "Raditonic",
	Stonitonic:   "Stonitonic",
	Syptitonic:   "Syptitonic",
	Ionythitonic: "Ionythitonic",

	// 5.14 Aeolanitonic scale
	Aeolanitonic: "Aeolanitonic",
	Danitonic:    "Danitonic",
	Ionaritonic:  "Ionaritonic",
	Dynitonic:    "Dynitonic",
	Zyditonic:    "Zyditonic",

	// 5.15 Aeolacritonic scale
	Aeolacritonic: "Aeolacritonic",
	Zythitonic:    "Zythitonic",
	Dyritonic:     "Dyritonic",
	Koptitonic:    "Koptitonic",
	Thocritonic:   "Thocritonic",

	// 5.16 Lycritonic scale
	Lycritonic: "Lycritonic",
	Daptitonic: "Daptitonic",
	Kygitonic:  "Kygitonic",
	Mocritonic: "Mocritonic",
	Zynitonic:  "Zynitonic",

	// 5.17 Epygitonic scale
	Epygitonic: "Epygitonic",
	Zaptitonic: "Zaptitonic",
	Kagitonic:  "Kagitonic",
	Zogitonic:  "Zogitonic",
	Epyritonic: "Epyritonic",

	// 5.18 Zothitonic scale
	Zothitonic:    "Zothitonic",
	Phrolitonic:   "Phrolitonic",
	Ionagitonic:   "Ionagitonic",
	Aeolapritonic: "Aeolapritonic",
	Kyritonic:     "Kyritonic",

	// 5.19 Ionyptitonic scale
	Ionyptitonic: "Ionyptitonic",
	Gyritonic:    "Gyritonic",
	Zalitonic:    "Zalitonic",
	Stolitonic:   "Stolitonic",
	Bylitonic:    "Bylitonic",

	// 5.20 Thoditonic scale
	Thoditonic:  "Thoditonic",
	Dogitonic:   "Dogitonic",
	Phralitonic: "Phralitonic",
	Garitonic:   "Garitonic",
	Soptitonic:  "Soptitonic",

	// 5.21 Kataritonic scale
	Kataritonic: "Kataritonic",
	Sylitonic:   "Sylitonic",
	Thonitonic:  "Thonitonic",
	Phropitonic: "Phropitonic",
	Staditonic:  "Staditonic",

	// 5.22 Lyditonic scale
	Lyditonic:  "Lyditonic",
	Mythitonic: "Mythitonic",
	Sogitonic:  "Sogitonic",
	Gothitonic: "Gothitonic",
	Rothitonic: "Rothitonic",

	// 5.23 Zylitonic scale
	Zylitonic:    "Zylitonic",
	Zoditonic:    "Zoditonic",
	Zaritonic:    "Zaritonic",
	Phrythitonic: "Phrythitonic",
	Rolitonic:    "Rolitonic",

	// 5.24 Ranitonic scale
	Ranitonic:    "Ranitonic",
	Laditonic:    "Laditonic",
	Poditonic:    "Poditonic",
	Ionothitonic: "Ionothitonic",
	Kanitonic:    "Kanitonic",

	// 5.25 Ryphitonic scale
	Ryphitonic:    "Ryphitonic",
	Gylitonic:     "Gylitonic",
	Aeolycritonic: "Aeolycritonic",
	Pynitonic:     "Pynitonic",
	Zanitonic:     "Zanitonic",

	// 5.26 Phronitonic scale
	Phronitonic: "Phronitonic",
	Banitonic:   "Banitonic",
	Aeronitonic: "Aeronitonic",
	Golitonic:   "Golitonic",
	Dyptitonic:  "Dyptitonic",

	// 5.27 Aerynitonic scale
	Aerynitonic:  "Aerynitonic",
	Palitonic:    "Palitonic",
	Stothitonic:  "Stothitonic",
	Aerophitonic: "Aerophitonic",
	Katagitonic:  "Katagitonic",

	// 5.28 Ionoditonic scale
	Ionoditonic: "Ionoditonic",
	Bogitonic:   "Bogitonic",
	Mogitonic:   "Mogitonic",
	Docritonic:  "Docritonic",
	Epaditonic:  "Epaditonic",

	// 5.29 Mixitonic scale
	Mixitonic:    "Mixitonic",
	Phrothitonic: "Phrothitonic",
	Katycritonic: "Katycritonic",
	Ionalitonic:  "Ionalitonic",
	Loptitonic:   "Loptitonic",

	// 5.30 Thyritonic scale
	Thyritonic:  "Thyritonic",
	Thoptitonic: "Thoptitonic",
	Bycritonic:  "Bycritonic",
	Pathitonic:  "Pathitonic",
	Myditonic:   "Myditonic",

	// 5.31 Bolitonic scale
	Bolitonic:   "Bolitonic",
	Bothitonic:  "Bothitonic",
	Kataditonic: "Kataditonic",
	Koditonic:   "Koditonic",
	Tholitonic:  "Tholitonic",

	// 6.1 Epathimic scale
	Epathimic: "Epathimic",
	Mynimic:   "Mynimic",
	Rocrimic:  "Rocrimic",
	Eporimic:  "Eporimic",
	Thaptimic: "Thaptimic",
	Lothimic:  "Lothimic",

	// 6.2 Dyrimic scale
	Dyrimic:    "Dyrimic",
	Koptimic:   "Koptimic",
	Thocrimic:  "Thocrimic",
	Aeolanimic: "Aeolanimic",
	Danimic:    "Danimic",
	Ionarimic:  "Ionarimic",

	// 6.3 Daptimic scale
	Daptimic: "Daptimic",
	Kygimic:  "Kygimic",
	Mocrimic: "Mocrimic",
	Zynimic:  "Zynimic",
	Aeolimic: "Aeolimic",
	Zythimic: "Zythimic",

	// 6.4 Epygimic scale
	Epygimic: "Epygimic",
	Zaptimic: "Zaptimic",
	Kagimic:  "Kagimic",
	Zogimic:  "Zogimic",
	Epyrimic: "Epyrimic",
	Lycrimic: "Lycrimic",

	// 6.5 Bylimic scale
	Bylimic:     "Bylimic",
	Zothimic:    "Zothimic",
	Phrolimic:   "Phrolimic",
	Ionagimic:   "Ionagimic",
	Aeolaphimic: "Aeolaphimic",
	Kycrimic:    "Kycrimic",

	// 6.6 Garimic scale
	Garimic:    "Garimic",
	Soptimic:   "Soptimic",
	Ionyptimic: "Ionyptimic",
	Gyrimic:    "Gyrimic",
	Zalimic:    "Zalimic",
	Stolimic:   "Stolimic",

	// 6.7 Thonimic scale
	Thonimic: "Thonimic",
	Stadimic: "Stadimic",
	Thodimic: "Thodimic",

	// 6.8 Mythimic scale
	Mythimic:  "Mythimic",
	Sogimic:   "Sogimic",
	Gogimic:   "Gogimic",
	Rothimic:  "Rothimic",
	Katarimic: "Katarimic",
	Sylimic:   "Sylimic",

	// 6.9 Mixolimic scale
	Mixolimic:   "Mixolimic",
	Dadimic:     "Dadimic",
	Aeolyphimic: "Aeolyphimic",
	Gycrimic:    "Gycrimic",
	Pyrimic:     "Pyrimic",
	Lydimic:     "Lydimic",

	// 6.10 Ionacrimic scale
	Ionacrimic: "Ionacrimic",
	Gathimic:   "Gathimic",
	Ionynimic:  "Ionynimic",
	Phrynimic:  "Phrynimic",
	Stathimic:  "Stathimic",
	Thatimic:   "Thatimic",

	// 6.11 Dalimic scale
	Dalimic:  "Dalimic",
	Dygimic:  "Dygimic",
	Zolimic:  "Zolimic",
	Epogimic: "Epogimic",
	Lanimic:  "Lanimic",
	Paptimic: "Paptimic",

	// 6.12 Darmic scale
	Darmic:     "Darmic",
	Lonimic:    "Lonimic",
	Ionycrimic: "Ionycrimic",
	Phradimic:  "Phradimic",
	Aeolorimic: "Aeolorimic",
	Gonimic:    "Gonimic",

	// 6.13 Phracrimic scale
	Phracrimic: "Phracrimic",
	Aerathimic: "Aerathimic",
	Sarimic:    "Sarimic",
	Zoptimic:   "Zoptimic",
	Zeracrimic: "Zeracrimic",
	Byptimic:   "Byptimic",

	// 6.14 Starimic scale
	Starimic:   "Starimic",
	Phrathimic: "Phrathimic",
	Saptimic:   "Saptimic",
	Aerodimic:  "Aerodimic",
	Macrimic:   "Macrimic",
	Rogimic:    "Rogimic",

	// 6.15 Bygimic scale
	Bygimic:    "Bygimic",
	Thycrimic:  "Thycrimic",
	Aeoladimic: "Aeoladimic",
	Dylimic:    "Dylimic",
	Eponimic:   "Eponimic",
	Katygimic:  "Katygimic",

	// 6.16 Stalimic scale
	Stalimic:    "Stalimic",
	Stoptimic:   "Stoptimic",
	Zygimic:     "Zygimic",
	Kataptimic:  "Kataptimic",
	Aeolaptimic: "Aeolaptimic",
	Pothimic:    "Pothimic",

	// 6.17 Rycrimic scale
	Rycrimic:  "Rycrimic",
	Ronimic:   "Ronimic",
	Stycrimic: "Stycrimic",
	Katorimic: "Katorimic",
	Epythimic: "Epythimic",
	Kaptimic:  "Kaptimic",

	// 6.18 Katythimic scale
	Katythimic: "Katythimic",
	Madimic:    "Madimic",
	Aerygimic:  "Aerygimic",
	Pylimic:    "Pylimic",
	Ionathimic: "Ionathimic",
	Morimic:    "Morimic",

	// 6.19 Aerycrimic scale
	Aerycrimic: "Aerycrimic",
	Ganimic:    "Ganimic",
	Eparimic:   "Eparimic",
	Lyrimic:    "Lyrimic",
	Phraptimic: "Phraptimic",
	Bacrimic:   "Bacrimic",

	// 6.20 Phralimic scale
	Phralimic:  "Phralimic",
	Phrogimic:  "Phrogimic",
	Rathimic:   "Rathimic",
	Katocrimic: "Katocrimic",
	Phryptimic: "Phryptimic",
	Katynimic:  "Katynimic",

	// 6.21 Solimic scale
	Solimic:    "Solimic",
	Ionolimic:  "Ionolimic",
	Ionophimic: "Ionophimic",
	Aeologimic: "Aeologimic",
	Zadimic:    "Zadimic",
	Sygimic:    "Sygimic",

	// 6.22 Thogimic scale
	Thogimic:    "Thogimic",
	Rythimic:    "Rythimic",
	Donimic:     "Donimic",
	Aeoloptimic: "Aeoloptimic",
	Panimic:     "Panimic",
	Lodimic:     "Lodimic",

	// 6.23 Laptimic scale
	Laptimic:   "Laptimic",
	Lygimic:    "Lygimic",
	Logimic:    "Logimic",
	Lalimic:    "Lalimic",
	Sothimic:   "Sothimic",
	Phrocrimic: "Phrocrimic",

	// 6.24 Modimic scale
	Modimic:    "Modimic",
	Barimic:    "Barimic",
	Poptimic:   "Poptimic",
	Sagimic:    "Sagimic",
	Aelothimic: "Aelothimic",
	Socrimic:   "Socrimic",

	// 6.25 Syrimic scale
	Syrimic:    "Syrimic",
	Stodimic:   "Stodimic",
	Ionocrimic: "Ionocrimic",
	Zycrimic:   "Zycrimic",
	Ionygimic:  "Ionygimic",
	Katathimic: "Katathimic",

	// 6.26 Bolimic scale
	Bolimic:   "Bolimic",
	Bothimic:  "Bothimic",
	Katadimic: "Katadimic",
	Kodimic:   "Kodimic",
	Tholimic:  "Tholimic",
	Ralimic:   "Ralimic",

	// 6.27 Kanimic scale
	Kanimic:    "Kanimic",
	Zylimic:    "Zylimic",
	Zodimic:    "Zodimic",
	Zarimic:    "Zarimic",
	Phrythimic: "Phrythimic",
	Rorimic:    "Rorimic",

	// 6.28 Pynimic scale
	Pynimic:    "Pynimic",
	Zanimic:    "Zanimic",
	Ranimic:    "Ranimic",
	Ladimic:    "Ladimic",
	Podimic:    "Podimic",
	Ionothimic: "Ionothimic",

	// 6.29 Kytrimic scale
	Kytrimic:    "Kytrimic",
	Golimic:     "Golimic",
	Dyptimic:    "Dyptimic",
	Ryrimic:     "Ryrimic",
	Gylimic:     "Gylimic",
	Aeolycrimic: "Aeolycrimic",

	// 6.30 Palimic scale
	Palimic:   "Palimic",
	Stothimic: "Stothimic",
	Aeronimic: "Aeronimic",
	Katagimic: "Katagimic",
	Phronimic: "Phronimic",
	Banimic:   "Banimic",

	// 6.31 Ionodimic scale
	Ionodimic: "Ionodimic",
	Bogimic:   "Bogimic",
	Mogimic:   "Mogimic",
	Docrimic:  "Docrimic",
	Epadimic:  "Epadimic",
	Aerynimic: "Aerynimic",

	// 6.32 Mydimic scale
	Mydimic:    "Mydimic",
	Thyptimic:  "Thyptimic",
	Phrothimic: "Phrothimic",
	Katycrimic: "Katycrimic",
	Ionalimic:  "Ionalimic",
	Loptimic:   "Loptimic",

	// 6.33 Zagimic scale
	Zagimic:   "Zagimic",
	Lagimic:   "Lagimic",
	Thyrimic:  "Thyrimic",
	Thothimic: "Thothimic",
	Bycrimic:  "Bycrimic",
	Pathimic:  "Pathimic",

	// 6.34 Mothimic scale
	Mothimic:  "Mothimic",
	Aeranimic: "Aeranimic",
	Ragimic:   "Ragimic",
	Dolimic:   "Dolimic",
	Porimic:   "Porimic",
	Aerylimic: "Aerylimic",

	// 6.35 Bocrimic scale
	Bocrimic:    "Bocrimic",
	Gythimic:    "Gythimic",
	Pagimic:     "Pagimic",
	Aeolythimic: "Aeolythimic",
	Molimic:     "Molimic",
	Staptimic:   "Staptimic",

	// 6.36 Zacrimic scale
	Zacrimic:  "Zacrimic",
	Larimic:   "Larimic",
	Thacrimic: "Thacrimic",
	Stydimic:  "Stydimic",
	Lorimic:   "Lorimic",
	Ionadimic: "Ionadimic",

	// 6.37 Ionythimic scale
	Ionythimic: "Ionythimic",
	Aerythimic: "Aerythimic",

	// 6.38 Dynimic scale
	Dynimic:  "Dynimic",
	Zydimic:  "Zydimic",
	Zathimic: "Zathimic",
	Radimic:  "Radimic",
	Stonimic: "Stonimic",
	Syptimic: "Syptimic",

	// 6.39 Ponimic scale
	Ponimic:  "Ponimic",
	Kadimic:  "Kadimic",
	Gynimic:  "Gynimic",
	Thydimic: "Thydimic",
	Polimic:  "Polimic",
	Thanimic: "Thanimic",

	// 6.40 Lathimic scale
	Lathimic:   "Lathimic",
	Aeralimic:  "Aeralimic",
	Kynimic:    "Kynimic",
	Stynimic:   "Stynimic",
	Epytimic:   "Epytimic",
	Katoptimic: "Katoptimic",

	// 6.41 Galimic scale
	Galimic:   "Galimic",
	Kathimic:  "Kathimic",
	Lylimic:   "Lylimic",
	Epalimic:  "Epalimic",
	Epacrimic: "Epacrimic",
	Sathimic:  "Sathimic",

	// 6.42 Katanimic scale
	Katanimic:  "Katanimic",
	Katyrimic:  "Katyrimic",
	Rynimic:    "Rynimic",
	Pogimic:    "Pogimic",
	Aeraptimic: "Aeraptimic",
	Epylimic:   "Epylimic",

	// 6.43 Manimic scale
	Manimic:  "Manimic",
	Marimic:  "Marimic",
	Locrimic: "Locrimic",
	Rylimic:  "Rylimic",
	Epatimic: "Epatimic",
	Byrimic:  "Byrimic",

	// 6.44 Kocrimic scale
	Kocrimic:  "Kocrimic",
	Korimic:   "Korimic",
	Lynimic:   "Lynimic",
	Malimic:   "Malimic",
	Synimic:   "Synimic",
	Phragimic: "Phragimic",

	// 6.45 Mycrimic scale
	Mycrimic:   "Mycrimic",
	Ionorimic:  "Ionorimic",
	Phrydimic:  "Phrydimic",
	Zyptimic:   "Zyptimic",
	Katothimic: "Katothimic",
	Phrylimic:  "Phrylimic",

	// 6.46 Aerothimic scale
	Aerothimic: "Aerothimic",
	Stagimic:   "Stagimic",
	Dorimic:    "Dorimic",
	Phrycrimic: "Phrycrimic",
	Kyptimic:   "Kyptimic",
	Ionylimic:  "Ionylimic",

	// 6.47 Epynimic scale
	Epynimic:   "Epynimic",
	Ionogimic:  "Ionogimic",
	Kydimic:    "Kydimic",
	Gaptimic:   "Gaptimic",
	Tharimic:   "Tharimic",
	Ionaphimic: "Ionaphimic",

	// 6.48 Thoptimic scale
	Thoptimic:  "Thoptimic",
	Bagimic:    "Bagimic",
	Kyrimic:    "Kyrimic",
	Sonimic:    "Sonimic",
	Aeolonimic: "Aeolonimic",
	Rygimic:    "Rygimic",

	// 6.49 Thagimic scale
	Thagimic:  "Thagimic",
	Kolimic:   "Kolimic",
	Dycrimic:  "Dycrimic",
	Epycrimic: "Epycrimic",
	Gocrimic:  "Gocrimic",
	Katolimic: "Katolimic",

	// 6.50 Dagimic scale
	Dagimic:    "Dagimic",
	Aeolydimic: "Aeolydimic",
	Parimic:    "Parimic",
	Ionaptimic: "Ionaptimic",
	Thylimic:   "Thylimic",
	Lolimic:    "Lolimic",

	// 6.51 Thalimic scale
	Thalimic:   "Thalimic",
	Stygimic:   "Stygimic",
	Aeolygimic: "Aeolygimic",
	Aerogimic:  "Aerogimic",
	Dacrimic:   "Dacrimic",
	Baptimic:   "Baptimic",

	// 6.52 Stythimic scale
	Stythimic: "Stythimic",
	Kothimic:  "Kothimic",
	Pygimic:   "Pygimic",
	Rodimic:   "Rodimic",
	Sorimic:   "Sorimic",
	Monimic:   "Monimic",

	// 6.53 Aeragimic scale
	Aeragimic: "Aeragimic",
	Epothimic: "Epothimic",
	Salimic:   "Salimic",
	Lyptimic:  "Lyptimic",
	Katonimic: "Katonimic",
	Gygimic:   "Gygimic",

	// 6.54 Aeradimic scale
	Aeradimic: "Aeradimic",
	Zyrimic:   "Zyrimic",
	Stylimic:  "Stylimic",

	// 6.55 Lythimic scale
	Lythimic:  "Lythimic",
	Dodimic:   "Dodimic",
	Katalimic: "Katalimic",

	// 6.56 Boptimic scale
	Boptimic:    "Boptimic",
	Stogimic:    "Stogimic",
	Thynimic:    "Thynimic",
	Aeolathimic: "Aeolathimic",
	Bythimic:    "Bythimic",
	Padimic:     "Padimic",

	// 6.57 Dathimic scale
	Dathimic: "Dathimic",
	Epagimic: "Epagimic",
	Raptimic: "Raptimic",
	Epolimic: "Epolimic",
	Sythimic: "Sythimic",
	Sydimic:  "Sydimic",

	// 6.58 Gacrimic
	Gacrimic:    "Gacrimic",
	Borimic:     "Borimic",
	Sycrimic:    "Sycrimic",
	Gadimic:     "Gadimic",
	Aeolocrimic: "Aeolocrimic",
	Phrygimic:   "Phrygimic",

	// 6.59 WholeTone
	WholeTone: "WholeTone",

	// 7.1 Lydian scale
	Lydian:     "Lydian",
	Mixolydian: "Mixolydian",
	Aeolian:    "Aeolian",
	Locrian:    "Locrian",
	Ionian:     "Ionian",
	Dorian:     "Dorian",
	Phrygian:   "Phrygian",

	// 7.2 Ionythian scale
	Ionythian: "Ionythian",
	Aeolyrian: "Aeolyrian",
	Gorian:    "Gorian",
	Aeolodian: "Aeolodian",
	Doptian:   "Doptian",
	Aeraphian: "Aeraphian",
	Zacrian:   "Zacrian",

	// 7.3 Ionarian scale
	Ionarian: "Ionarian",
	Dynian:   "Dynian",
	Zydian:   "Zydian",
	Zathian:  "Zathian",
	Radian:   "Radian",
	Stonian:  "Stonian",
	Syptian:  "Syptian",

	// 7.4 Aeolacrian scale
	Aeolacrian: "Aeolacrian",
	Zythian:    "Zythian",
	Dyrian:     "Dyrian",
	Koptian:    "Koptian",
	Thocrian:   "Thocrian",
	Aeolanian:  "Aeolanian",
	Danian:     "Danian",

	// 7.5 Zogian scale
	Zogian:  "Zogian",
	Epyrian: "Epyrian",
	Lycrian: "Lycrian",
	Daptian: "Daptian",
	Kygian:  "Kygian",
	Mocrian: "Mocrian",
	Zynian:  "Zynian",

	// 7.6 Phrolian scale
	Phrolian: "Phrolian",
	Ionagian: "Ionagian",
	Aeodian:  "Aeodian",
	Kycrian:  "Kycrian",
	Epygian:  "Epygian",
	Zaptian:  "Zaptian",
	Kagian:   "Kagian",

	// 7.7 Soptian scale
	Soptian:   "Soptian",
	Ionyptian: "Ionyptian",
	Gyrian:    "Gyrian",
	Zalian:    "Zalian",
	Stolian:   "Stolian",
	Bylian:    "Bylian",
	Zothian:   "Zothian",

	// 7.8 Thonian scale
	Thonian:    "Thonian",
	Phrorian:   "Phrorian",
	Stadian:    "Stadian",
	Thodian:    "Thodian",
	Dogian:     "Dogian",
	Mixopyrian: "Mixopyrian",
	Garian:     "Garian",

	// 7.9 Epathian scale
	Epathian: "Epathian",
	Mythian:  "Mythian",
	Sogian:   "Sogian",
	Gogian:   "Gogian",
	Rothian:  "Rothian",
	Katarian: "Katarian",
	Stylian:  "Stylian",

	// 7.10 Stathian scale
	Stathian:    "Stathian",
	Mixonyphian: "Mixonyphian",
	Magian:      "Magian",
	Dadian:      "Dadian",
	Aeolylian:   "Aeolylian",
	Gycrian:     "Gycrian",
	Pyrian:      "Pyrian",

	// 7.11 Epogian scale
	Epogian:   "Epogian",
	Lanian:    "Lanian",
	Paptian:   "Paptian",
	Ionacrian: "Ionacrian",
	Gathian:   "Gathian",
	Ionyphian: "Ionyphian",
	Phrynian:  "Phrynian",

	// 7.12 Ionycrian scale
	Ionycrian: "Ionycrian",
	Phradian:  "Phradian",
	Aeolorian: "Aeolorian",
	Gonian:    "Gonian",
	Dalian:    "Dalian",
	Dygian:    "Dygian",
	Zolian:    "Zolian",

	// 7.13 Aerathian scale
	Aerathian: "Aerathian",
	Sarian:    "Sarian",
	Zoptian:   "Zoptian",
	Aeracrian: "Aeracrian",
	Byptian:   "Byptian",
	Darian:    "Darian",
	Lonian:    "Lonian",

	// 7.14 Aeopian scale
	Aeopian:  "Aeopian",
	Rygian:   "Rygian",
	Epynian:  "Epynian",
	Ionogian: "Ionogian",
	Kydian:   "Kydian",
	Gaptian:  "Gaptian",
	Tharian:  "Tharian",

	// 7.15 Epycrian scale
	Epycrian: "Epycrian",
	Gocrian:  "Gocrian",
	Katolian: "Katolian",
	Thoptian: "Thoptian",
	Bagian:   "Bagian",
	Kyrian:   "Kyrian",
	Sonian:   "Sonian",

	// 7.16 Parian scale
	Parian:    "Parian",
	Ionaptian: "Ionaptian",
	Thylian:   "Thylian",
	Lolian:    "Lolian",
	Thagian:   "Thagian",
	Kolian:    "Kolian",
	Dycrian:   "Dycrian",

	// 7.17 Stygian scale
	Stygian:   "Stygian",
	Aeolygian: "Aeolygian",
	Aerogian:  "Aerogian",
	Dacrian:   "Dacrian",
	Baptian:   "Baptian",
	Dagian:    "Dagian",
	Aeolydian: "Aeolydian",

	// 7.18 Stythian scale
	Stythian: "Stythian",
	Kothian:  "Kothian",
	Pygian:   "Pygian",
	Rodian:   "Rodian",
	Sorian:   "Sorian",
	Monian:   "Monian",
	Thalian:  "Thalian",

	// 7.19 Zorian scale
	Zorian:   "Zorian",
	Aeragian: "Aeragian",
	Epothian: "Epothian",
	Salian:   "Salian",
	Lyptian:  "Lyptian",
	Katonian: "Katonian",
	Gyphian:  "Gyphian",

	// 7.20 Thacrian scale
	Thacrian:   "Thacrian",
	Dodian:     "Dodian",
	Aeolyptian: "Aeolyptian",
	Aeolonian:  "Aeolonian",
	Aeradian:   "Aeradian",
	Aeolagian:  "Aeolagian",
	Zyrian:     "Zyrian",

	// 7.21 Aeolathian scale
	Aeolathian: "Aeolathian",
	Bythian:    "Bythian",
	Padian:     "Padian",
	Rolian:     "Rolian",
	Pydian:     "Pydian",
	Thygian:    "Thygian",
	Katalian:   "Katalian",

	// 7.22 Saptian scale
	Saptian:  "Saptian",
	Aerodian: "Aerodian",
	Macrian:  "Macrian",
	Rogian:   "Rogian",
	Boptian:  "Boptian",
	Stogian:  "Stogian",
	Thynian:  "Thynian",

	// 7.23 Thycrian scale
	Thycrian:  "Thycrian",
	Aeoladian: "Aeoladian",
	Dylian:    "Dylian",
	Eponian:   "Eponian",
	Katygian:  "Katygian",
	Starian:   "Starian",
	Phrathian: "Phrathian",

	// 7.24 Stalian scale
	Stalian:    "Stalian",
	Stoptian:   "Stoptian",
	Zygian:     "Zygian",
	Kataptian:  "Kataptian",
	Aeolaptian: "Aeolaptian",
	Pothian:    "Pothian",
	Bygian:     "Bygian",

	// 7.25 Morian scale
	Morian:   "Morian",
	Rycrian:  "Rycrian",
	Ronian:   "Ronian",
	Stycrian: "Stycrian",
	Katorian: "Katorian",
	Epythian: "Epythian",
	Kaptian:  "Kaptian",

	// 7.26 Phraptian scale
	Phraptian: "Phraptian",
	Bacrian:   "Bacrian",
	Katythian: "Katythian",
	Madian:    "Madian",
	Aerygian:  "Aerygian",
	Pylian:    "Pylian",
	Ionathian: "Ionathian",

	// 7.27 Katocrian scale
	Katocrian: "Katocrian",
	Phryptian: "Phryptian",
	Katynian:  "Katynian",
	Aerycrian: "Aerycrian",
	Ganian:    "Ganian",
	Eparian:   "Eparian",
	Lyrian:    "Lyrian",

	// 7.28 Ionopian scale
	Ionopian:  "Ionopian",
	Aeologian: "Aeologian",
	Zadian:    "Zadian",
	Sygian:    "Sygian",
	Phralian:  "Phralian",
	Phrogian:  "Phrogian",
	Rathian:   "Rathian",

	// 7.29 Rythian scale
	Rythian:    "Rythian",
	Donian:     "Donian",
	Aeoloptian: "Aeoloptian",
	Panian:     "Panian",
	Lodian:     "Lodian",
	Solian:     "Solian",
	Ionolian:   "Ionolian",

	// 7.30 Laptian scale
	Laptian:   "Laptian",
	Lygian:    "Lygian",
	Logian:    "Logian",
	Lalian:    "Lalian",
	Sothian:   "Sothian",
	Phrocrian: "Phrocrian",
	Thogian:   "Thogian",

	// 7.31 Katathian scale
	Katathian:   "Katathian",
	Modian:      "Modian",
	Barian:      "Barian",
	Mixolocrian: "Mixolocrian",
	Sagian:      "Sagian",
	Aeolothian:  "Aeolothian",
	Socrian:     "Socrian",

	// 7.32 Tholian scale
	Tholian:   "Tholian",
	Ralian:    "Ralian",
	Syrian:    "Syrian",
	Stodian:   "Stodian",
	Ionocrian: "Ionocrian",
	Zycrian:   "Zycrian",
	Ionygian:  "Ionygian",

	// 7.33 Zarian scale
	Zarian:    "Zarian",
	Phrythian: "Phrythian",
	Rorian:    "Rorian",
	Bolian:    "Bolian",
	Bothian:   "Bothian",
	Katadian:  "Katadian",
	Kodian:    "Kodian",

	// 7.34 Ranian scale
	Ranian:    "Ranian",
	Ladian:    "Ladian",
	Podian:    "Podian",
	Ionothian: "Ionothian",
	Kanian:    "Kanian",
	Zylian:    "Zylian",
	Zodian:    "Zodian",

	// 7.35 Golian scale
	Golian:     "Golian",
	Dyptian:    "Dyptian",
	Ryphian:    "Ryphian",
	Gylian:     "Gylian",
	Aeolycrian: "Aeolycrian",
	Pynian:     "Pynian",
	Zanian:     "Zanian",

	// 7.36 Palian scale
	Palian:   "Palian",
	Stothian: "Stothian",
	Aerorian: "Aerorian",
	Katagian: "Katagian",
	Phronian: "Phronian",
	Banian:   "Banian",
	Aeronian: "Aeronian",

	// 7.37 Loptian scale
	Loptian:  "Loptian",
	Ionodian: "Ionodian",
	Bogian:   "Bogian",
	Mogian:   "Mogian",
	Docrian:  "Docrian",
	Epadian:  "Epadian",
	Aerynian: "Aerynian",

	// 7.38 Bycrian scale
	Bycrian:   "Bycrian",
	Pathian:   "Pathian",
	Mydian:    "Mydian",
	Thyptian:  "Thyptian",
	Phrothian: "Phrothian",
	Katycrian: "Katycrian",
	Ionalian:  "Ionalian",

	// 7.39 Dolian scale
	Dolian:     "Dolian",
	Porian:     "Porian",
	Aerylian:   "Aerylian",
	Zagian:     "Zagian",
	Lagian:     "Lagian",
	Tyrian:     "Tyrian",
	Mixonorian: "Mixonorian",

	// 7.40 Pagian scale
	Pagian:     "Pagian",
	Aeolythian: "Aeolythian",
	Molian:     "Molian",
	Staptian:   "Staptian",
	Mothian:    "Mothian",
	Aeranian:   "Aeranian",
	Ragian:     "Ragian",

	// 7.41 Larian scale
	Larian:      "Larian",
	Lythian:     "Lythian",
	Stydian:     "Stydian",
	Lorian:      "Lorian",
	Ionadian:    "Ionadian",
	Bocrian:     "Bocrian",
	Mixolythian: "Mixolythian",

	// 7.42 Thadian scale
	Thadian:   "Thadian",
	Sanian:    "Sanian",
	Ionydian:  "Ionydian",
	Epydian:   "Epydian",
	Katydian:  "Katydian",
	Mathian:   "Mathian",
	Aeryptian: "Aeryptian",

	// 7.43 Pythian scale
	Pythian:  "Pythian",
	Katylian: "Katylian",
	Bydian:   "Bydian",
	Bynian:   "Bynian",
	Galian:   "Galian",
	Zonian:   "Zonian",
	Myrian:   "Myrian",

	// 7.44 Katogian scale
	Katogian: "Katogian",
	Stacrian: "Stacrian",
	Styrian:  "Styrian",
	Ionyrian: "Ionyrian",
	Phrodian: "Phrodian",
	Pycrian:  "Pycrian",
	Gyptian:  "Gyptian",

	// 7.45 Katacrian scale
	Katacrian: "Katacrian",
	Sodian:    "Sodian",
	Bathian:   "Bathian",
	Mylian:    "Mylian",
	Godian:    "Godian",
	Thorian:   "Thorian",
	Zocrian:   "Zocrian",

	// 7.46 Stanian scale
	Stanian:   "Stanian",
	Epanian:   "Epanian",
	Konian:    "Konian",
	Stocrian:  "Stocrian",
	Kalian:    "Kalian",
	Phroptian: "Phroptian",
	Dydian:    "Dydian",

	// 7.47 Katyptian scale
	Katyptian: "Katyptian",
	Epodian:   "Epodian",
	Mygian:    "Mygian",
	Pacrian:   "Pacrian",
	Aerocrian: "Aerocrian",
	Aeolarian: "Aeolarian",
	Kythian:   "Kythian",

	// 7.48 Bonian scale
	Bonian:   "Bonian",
	Badian:   "Badian",
	Katodian: "Katodian",
	Sadian:   "Sadian",
	Dothian:  "Dothian",
	Moptian:  "Moptian",
	Aeryrian: "Aeryrian",

	// 7.49 Epagian scale
	Epagian:  "Epagian",
	Raptian:  "Raptian",
	Epolian:  "Epolian",
	Sythian:  "Sythian",
	Sydian:   "Sydian",
	Epocrian: "Epocrian",
	Kylian:   "Kylian",

	// 7.50 Gacrian scale
	Gacrian:    "Gacrian",
	Borian:     "Borian",
	Sycrian:    "Sycrian",
	Gadian:     "Gadian",
	Aeolocrian: "Aeolocrian",
	Mixodorian: "Mixodorian",
	Dathian:    "Dathian",

	// 7.51 Katoptian scale
	Katoptian: "Katoptian",
	Ponian:    "Ponian",
	Kadian:    "Kadian",
	Gynian:    "Gynian",
	Thyphian:  "Thyphian",
	Polian:    "Polian",
	Thanian:   "Thanian",

	// 7.52 Epacrian scale
	Epacrian: "Epacrian",
	Sathian:  "Sathian",
	Lathian:  "Lathian",
	Aeralian: "Aeralian",
	Kynian:   "Kynian",
	Stynian:  "Stynian",
	Epyphian: "Epyphian",

	// 7.53 Pogian scale
	Pogian:    "Pogian",
	Aeraptian: "Aeraptian",
	Epylian:   "Epylian",
	Gamian:    "Gamian",
	Kathian:   "Kathian",
	Lylian:    "Lylian",
	Epalian:   "Epalian",

	// 7.54 Eporian scale
	Eporian:  "Eporian",
	Rylian:   "Rylian",
	Epaptian: "Epaptian",
	Byrian:   "Byrian",
	Katanian: "Katanian",
	Katyrian: "Katyrian",
	Rynian:   "Rynian",

	// 7.55 Korian scale
	Korian:   "Korian",
	Lynian:   "Lynian",
	Malian:   "Malian",
	Synian:   "Synian",
	Phragian: "Phragian",
	Manian:   "Manian",
	Marian:   "Marian",

	// 7.56 Mycrian scale
	Mycrian:   "Mycrian",
	Ionorian:  "Ionorian",
	Phrydian:  "Phrydian",
	Zyptian:   "Zyptian",
	Katothian: "Katothian",
	Phrylian:  "Phrylian",
	Kocrian:   "Kocrian",

	// 7.57 Ionanian scale
	Ionanian:  "Ionanian",
	Aerothian: "Aerothian",
	Stagian:   "Stagian",
	Lothian:   "Lothian",
	Phrycrian: "Phrycrian",
	Kyptian:   "Kyptian",
	Ionylian:  "Ionylian",

	// 7.58 Gydian scale
	Gydian:   "Gydian",
	Kogian:   "Kogian",
	Rarian:   "Rarian",
	Aerolian: "Aerolian",
	Karian:   "Karian",
	Myptian:  "Myptian",
	Rydian:   "Rydian",

	// 7.59 Aeolynian scale
	Aeolynian: "Aeolynian",
	Aeroptian: "Aeroptian",
	Phryrian:  "Phryrian",
	Gothian:   "Gothian",
	Storian:   "Storian",
	Pyptian:   "Pyptian",
	Thydian:   "Thydian",

	// 8.1 Aerycryllic scale
	Aerycryllic: "Aerycryllic",
	Gadyllic:    "Gadyllic",
	Solyllic:    "Solyllic",
	Zylyllic:    "Zylyllic",
	Mixodyllic:  "Mixodyllic",
	Soryllic:    "Soryllic",
	Godyllic:    "Godyllic",
	Epiphyllic:  "Epiphyllic",

	// 8.2 Pynyllic scale
	Pynyllic:  "Pynyllic",
	Bocryllic: "Bocryllic",
	Kogyllic:  "Kogyllic",
	Raryllic:  "Raryllic",
	Zycryllic: "Zycryllic",
	Mycryllic: "Mycryllic",
	Laptyllic: "Laptyllic",
	Pylyllic:  "Pylyllic",

	// 8.3 Pothyllic scale
	Pothyllic:   "Pothyllic",
	Phronyllic:  "Phronyllic",
	Stynyllic:   "Stynyllic",
	Rathyllic:   "Rathyllic",
	Aeryptyllic: "Aeryptyllic",
	Zydyllic:    "Zydyllic",
	Katolyllic:  "Katolyllic",
	Rythyllic:   "Rythyllic",

	// 8.4 Locryllic scale
	Locryllic:   "Locryllic",
	Bylyllic:    "Bylyllic",
	Sogyllic:    "Sogyllic",
	Ionycryllic: "Ionycryllic",
	Koptyllic:   "Koptyllic",
	Epyryllic:   "Epyryllic",
	Soptyllic:   "Soptyllic",
	Aeolylyllic: "Aeolylyllic",

	// 8.5 Aeracryllic scale
	Aeracryllic: "Aeracryllic",
	Epygyllic:   "Epygyllic",
	Thonyllic:   "Thonyllic",
	Lanyllic:    "Lanyllic",
	Phrynyllic:  "Phrynyllic",
	Lycryllic:   "Lycryllic",
	Ionyptyllic: "Ionyptyllic",
	Epathyllic:  "Epathyllic",

	// 8.6 Dydyllic scale
	Dydyllic:  "Dydyllic",
	Thogyllic: "Thogyllic",
	Rygyllic:  "Rygyllic",
	Bycryllic: "Bycryllic",
	Zacryllic: "Zacryllic",
	Panyllic:  "Panyllic",
	Dyryllic:  "Dyryllic",
	Zathyllic: "Zathyllic",

	// 8.7 Dagyllic scale
	Dagyllic:    "Dagyllic",
	Katalyllic:  "Katalyllic",
	Katoryllic:  "Katoryllic",
	Dodyllic:    "Dodyllic",
	Zogyllic:    "Zogyllic",
	Madyllic:    "Madyllic",
	Dycryllic:   "Dycryllic",
	Aeologyllic: "Aeologyllic",

	// 8.8 Sydyllic scale
	Sydyllic:   "Sydyllic",
	Katogyllic: "Katogyllic",
	Zygyllic:   "Zygyllic",
	Aeralyllic: "Aeralyllic",

	// 8.9 Bacryllic scale
	Bacryllic:  "Bacryllic",
	Aerygyllic: "Aerygyllic",
	Dathyllic:  "Dathyllic",
	Boptyllic:  "Boptyllic",
	Bagyllic:   "Bagyllic",
	Mathyllic:  "Mathyllic",
	Styptyllic: "Styptyllic",
	Zolyllic:   "Zolyllic",

	// 8.10 Rocryllic scale
	Rocryllic:  "Rocryllic",
	Zyryllic:   "Zyryllic",
	Sagyllic:   "Sagyllic",
	Epinyllic:  "Epinyllic",
	Katagyllic: "Katagyllic",
	Ragyllic:   "Ragyllic",
	Gothyllic:  "Gothyllic",
	Lythyllic:  "Lythyllic",

	// 8.11 Ionocryllic scale
	Ionocryllic: "Ionocryllic",
	Gocryllic:   "Gocryllic",
	Epiryllic:   "Epiryllic",
	Aeradyllic:  "Aeradyllic",
	Staptyllic:  "Staptyllic",
	Danyllic:    "Danyllic",
	Goptyllic:   "Goptyllic",
	Epocryllic:  "Epocryllic",

	// 8.12 Ionoptyllic scale
	Ionoptyllic: "Ionoptyllic",
	Aeoloryllic: "Aeoloryllic",
	Thydyllic:   "Thydyllic",
	Gycryllic:   "Gycryllic",
	Lyryllic:    "Lyryllic",
	Mogyllic:    "Mogyllic",
	Katodyllic:  "Katodyllic",
	Moptyllic:   "Moptyllic",

	// 8.13 Dolyllic scale
	Dolyllic:    "Dolyllic",
	Moryllic:    "Moryllic",
	Bydyllic:    "Bydyllic",
	Pocryllic:   "Pocryllic",
	Phracryllic: "Phracryllic",
	Gyryllic:    "Gyryllic",
	Phrygyllic:  "Phrygyllic",
	Dogyllic:    "Dogyllic",

	// 8.14 Thagyllic scale
	Thagyllic:   "Thagyllic",
	Thoptyllic:  "Thoptyllic",
	Phraptyllic: "Phraptyllic",
	Gylyllic:    "Gylyllic",
	Phralyllic:  "Phralyllic",
	Dygyllic:    "Dygyllic",
	Ronyllic:    "Ronyllic",
	Epogyllic:   "Epogyllic",

	// 8.15 Aeoladyllic scale
	Aeoladyllic: "Aeoladyllic",
	Kocryllic:   "Kocryllic",
	Lodyllic:    "Lodyllic",
	Bynyllic:    "Bynyllic",
	Kydyllic:    "Kydyllic",
	Bygyllic:    "Bygyllic",
	Phryptyllic: "Phryptyllic",
	Ionayllic:   "Ionayllic",

	// 8.16 Phroryllic scale
	Phroryllic:  "Phroryllic",
	Thyphyllic:  "Thyphyllic",
	Poptyllic:   "Poptyllic",
	Mixonyllic:  "Mixonyllic",
	Paptyllic:   "Paptyllic",
	Storyllic:   "Storyllic",
	Phrycryllic: "Phrycryllic",
	Palyllic:    "Palyllic",

	// 8.17 Phranyllic scale
	Phranyllic:  "Phranyllic",
	Stydyllic:   "Stydyllic",
	Zadyllic:    "Zadyllic",
	Zalyllic:    "Zalyllic",
	Zocryllic:   "Zocryllic",
	Katocryllic: "Katocryllic",
	Aerathyllic: "Aerathyllic",
	Stoptyllic:  "Stoptyllic",

	// 8.18 Lydyllic scale
	Lydyllic:   "Lydyllic",
	Radyllic:   "Radyllic",
	Stagyllic:  "Stagyllic",
	Ionoryllic: "Ionoryllic",
	Phrodyllic: "Phrodyllic",
	Aeragyllic: "Aeragyllic",
	Banyllic:   "Banyllic",
	Epothyllic: "Epothyllic",

	// 8.19 Zoryllic scale
	Zoryllic:   "Zoryllic",
	Phrolyllic: "Phrolyllic",
	Kolyllic:   "Kolyllic",
	Thodyllic:  "Thodyllic",
	Socryllic:  "Socryllic",
	Aeolyllic:  "Aeolyllic",
	Zythyllic:  "Zythyllic",
	Aeoryllic:  "Aeoryllic",

	// 8.20 Mixolydyllic
	Mixolydyllic:  "Mixolydyllic",
	Mixonyphyllic: "Mixonyphyllic",
	Aeolanyllic:   "Aeolanyllic",
	Thocryllic:    "Thocryllic",
	Kygyllic:      "Kygyllic",
	Ionagyllic:    "Ionagyllic",
	Gogyllic:      "Gogyllic",
	Phradyllic:    "Phradyllic",

	// 8.21 Ioniptyllic scale
	Ioniptyllic:  "Ioniptyllic",
	Kycryllic:    "Kycryllic",
	Aeolaptyllic: "Aeolaptyllic",
	Rodyllic:     "Rodyllic",
	Ionathyllic:  "Ionathyllic",
	Pythyllic:    "Pythyllic",
	Zonyllic:     "Zonyllic",
	Ryryllic:     "Ryryllic",

	// 8.22 Aeolothyllic scale
	Aeolothyllic: "Aeolothyllic",
	Ionyryllic:   "Ionyryllic",
	Rydyllic:     "Rydyllic",
	Gonyllic:     "Gonyllic",
	Rolyllic:     "Rolyllic",
	Katydyllic:   "Katydyllic",
	Zyptyllic:    "Zyptyllic",
	Modyllic:     "Modyllic",

	// 8.23 Maptyllic scale
	Maptyllic:   "Maptyllic",
	Aeraptyllic: "Aeraptyllic",
	Katadyllic:  "Katadyllic",
	Magyllic:    "Magyllic",
	Phrylyllic:  "Phrylyllic",
	Epigyllic:   "Epigyllic",
	Molyllic:    "Molyllic",
	Ponyllic:    "Ponyllic",

	// 8.24 Thyptyllic scale
	Thyptyllic:  "Thyptyllic",
	Ionogyllic:  "Ionogyllic",
	Aeolaryllic: "Aeolaryllic",
	Katygyllic:  "Katygyllic",
	Ganyllic:    "Ganyllic",
	Kyptyllic:   "Kyptyllic",
	Salyllic:    "Salyllic",
	Sanyllic:    "Sanyllic",

	// 8.25 Doptyllic scale
	Doptyllic:     "Doptyllic",
	Ionilyllic:    "Ionilyllic",
	Manyllic:      "Manyllic",
	Polyllic:      "Polyllic",
	Stanyllic:     "Stanyllic",
	Mixotharyllic: "Mixotharyllic",
	Eporyllic:     "Eporyllic",
	Aerynyllic:    "Aerynyllic",

	// 8.26 Lonyllic scale
	Lonyllic:    "Lonyllic",
	Sathyllic:   "Sathyllic",
	Layllic:     "Layllic",
	Saryllic:    "Saryllic",
	Thacryllic:  "Thacryllic",
	Aeolynyllic: "Aeolynyllic",
	Thadyllic:   "Thadyllic",
	Lynyllic:    "Lynyllic",

	// 8.27 Aeolathyllic scale
	Aeolathyllic: "Aeolathyllic",
	Aeolocryllic: "Aeolocryllic",
	Phroptyllic:  "Phroptyllic",
	Kodyllic:     "Kodyllic",
	Epaptyllic:   "Epaptyllic",
	Ionoyllic:    "Ionoyllic",
	Gyptyllic:    "Gyptyllic",
	Aerythyllic:  "Aerythyllic",

	// 8.28 Zagyllic scale
	Zagyllic:    "Zagyllic",
	Epacryllic:  "Epacryllic",
	Thorcryllic: "Thorcryllic",
	Loptyllic:   "Loptyllic",
	Katylyllic:  "Katylyllic",
	Malyllic:    "Malyllic",
	Mydyllic:    "Mydyllic",
	Thycryllic:  "Thycryllic",

	// 8.29 Gythyllic scale
	Gythyllic:   "Gythyllic",
	Pyryllic:    "Pyryllic",
	Rycryllic:   "Rycryllic",
	Phrathyllic: "Phrathyllic",
	Badyllic:    "Badyllic",
	Phrocryllic: "Phrocryllic",
	Staryllic:   "Staryllic",
	Zothyllic:   "Zothyllic",

	// 8.30 Tharyllic scale
	Tharyllic:    "Tharyllic",
	Sylyllic:     "Sylyllic",
	Lothyllic:    "Lothyllic",
	Daryllic:     "Daryllic",
	Monyllic:     "Monyllic",
	Styryllic:    "Styryllic",
	Aeolacryllic: "Aeolacryllic",
	Raptyllic:    "Raptyllic",

	// 8.31 Kataryllic scale
	Kataryllic:  "Kataryllic",
	Aerocryllic: "Aerocryllic",
	Zanyllic:    "Zanyllic",
	Aeolonyllic: "Aeolonyllic",
	Aeonyllic:   "Aeonyllic",
	Kyryllic:    "Kyryllic",
	Sythyllic:   "Sythyllic",
	Katycryllic: "Katycryllic",

	// 8.32 Stogyllic scale
	Stogyllic:   "Stogyllic",
	Ionidyllic:  "Ionidyllic",
	Stonyllic:   "Stonyllic",
	Stalyllic:   "Stalyllic",
	Poryllic:    "Poryllic",
	Mocryllic:   "Mocryllic",
	Aeolyryllic: "Aeolyryllic",
	Baryllic:    "Baryllic",

	// 8.33 Dalyllic scale
	Dalyllic:     "Dalyllic",
	Ionyphyllic:  "Ionyphyllic",
	Zaptyllic:    "Zaptyllic",
	Garyllic:     "Garyllic",
	Gathyllic:    "Gathyllic",
	Mixopyryllic: "Mixopyryllic",
	Ionacryllic:  "Ionacryllic",
	Stylyllic:    "Stylyllic",

	// 8.34 Stycryllic scale
	Stycryllic:  "Stycryllic",
	Ionothyllic: "Ionothyllic",
	Mythyllic:   "Mythyllic",
	Aerylyllic:  "Aerylyllic",
	Bonyllic:    "Bonyllic",
	Tholyllic:   "Tholyllic",
	Katyryllic:  "Katyryllic",
	Sadyllic:    "Sadyllic",

	// 8.35 Stolyllic scale
	Stolyllic: "Stolyllic",
	Logyllic:  "Logyllic",
	Dacryllic: "Dacryllic",
	Thynyllic: "Thynyllic",
	Gydyllic:  "Gydyllic",
	Eparyllic: "Eparyllic",
	Dynyllic:  "Dynyllic",
	Ionyllic:  "Ionyllic",

	// 8.36 Zaryllic scale
	Zaryllic:    "Zaryllic",
	Dythyllic:   "Dythyllic",
	Ionaryllic:  "Ionaryllic",
	Laryllic:    "Laryllic",
	Kataptyllic: "Kataptyllic",
	Sonyllic:    "Sonyllic",
	Pathyllic:   "Pathyllic",
	Loryllic:    "Loryllic",

	// 8.37 Aeronyllic scale
	Aeronyllic: "Aeronyllic",
	Pycryllic:  "Pycryllic",
	Mygyllic:   "Mygyllic",
	Lylyllic:   "Lylyllic",
	Daptyllic:  "Daptyllic",
	Ioninyllic: "Ioninyllic",
	Epaphyllic: "Epaphyllic",
	Lolyllic:   "Lolyllic",

	// 8.38 Stacryllic scale
	Stacryllic: "Stacryllic",
	Doryllic:   "Doryllic",
	Kadyllic:   "Kadyllic",
	Rynyllic:   "Rynyllic",
	Aerogyllic: "Aerogyllic",
	Rothyllic:  "Rothyllic",
	Kagyllic:   "Kagyllic",
	Stathyllic: "Stathyllic",

	// 8.39 Thyryllic scale
	Thyryllic:   "Thyryllic",
	Gygyllic:    "Gygyllic",
	Sodyllic:    "Sodyllic",
	Goryllic:    "Goryllic",
	Bothyllic:   "Bothyllic",
	Gynyllic:    "Gynyllic",
	Ionaptyllic: "Ionaptyllic",
	Phryryllic:  "Phryryllic",

	// 8.40 Racryllic scale
	Racryllic:   "Racryllic",
	Epicryllic:  "Epicryllic",
	Stygyllic:   "Stygyllic",
	Syryllic:    "Syryllic",
	Stythyllic:  "Stythyllic",
	Aerothyllic: "Aerothyllic",
	Mixoryllic:  "Mixoryllic",
	Thanyllic:   "Thanyllic",

	// 8.41 Roryllic
	Epotyllic: "Epotyllic",
	Epidyllic: "Epidyllic",
	Kaptyllic: "Kaptyllic",

	// 8.42 MajorDiminished scale
	MajorDiminished: "MajorDiminished",
	MinorDiminished: "MinorDiminished",

	// 9.1 Aerycrygic scale
	Aerycrygic: "Aerycrygic",
	Gadygic:    "Gadygic",
	Solygic:    "Solygic",
	Zylygic:    "Zylygic",
	Garygic:    "Garygic",
	Sorygic:    "Sorygic",
	Godygic:    "Godygic",
	Epithygic:  "Epithygic",
	Ionoptygic: "Ionoptygic",

	// 9.2 Kalygic scale
	Kalygic:    "Kalygic",
	Ionodygic:  "Ionodygic",
	Bythygic:   "Bythygic",
	Epygic:     "Epygic",
	Marygic:    "Marygic",
	Gaptygic:   "Gaptygic",
	Aeroptygic: "Aeroptygic",
	Mylygic:    "Mylygic",
	Galygic:    "Galygic",

	// 9.3 Mixolydygic scale
	Mixolydygic: "Mixolydygic",
	Ionycrygic:  "Ionycrygic",
	Zoptygic:    "Zoptygic",
	Phrygygic:   "Phrygygic",
	Locrygic:    "Locrygic",
	Gonygic:     "Gonygic",
	Aeracrygic:  "Aeracrygic",
	Aerathygic:  "Aerathygic",
	Dorygic:     "Dorygic",

	// 9.4 Dycrygic scale
	Dycrygic: "Dycrygic",
	Aeolygic: "Aeolygic",
	Dydygic:  "Dydygic",
	Tholygic: "Tholygic",
	Rynygic:  "Rynygic",
	Bycrygic: "Bycrygic",
	Zacrygic: "Zacrygic",
	Panygic:  "Panygic",
	Dyrygic:  "Dyrygic",

	// 9.5 Loptygic scale
	Loptygic:  "Loptygic",
	Katylygic: "Katylygic",
	Phradygic: "Phradygic",
	Mixodygic: "Mixodygic",
	Katalygic: "Katalygic",
	Katorygic: "Katorygic",
	Dogygic:   "Dogygic",
	Zodygic:   "Zodygic",
	Madygic:   "Madygic",

	// 9.6 Bagygic scale
	Bagygic:   "Bagygic",
	Mathygic:  "Mathygic",
	Styptygic: "Styptygic",
	Zolygic:   "Zolygic",
	Sydygic:   "Sydygic",
	Katygic:   "Katygic",
	Zyphygic:  "Zyphygic",
	Aeralygic: "Aeralygic",
	Ryptygic:  "Ryptygic",

	// 9.7 Apinygic scale
	Apinygic:  "Apinygic",
	Katagygic: "Katagygic",
	Radygic:   "Radygic",
	Gothygic:  "Gothygic",
	Lythygic:  "Lythygic",
	Bacrygic:  "Bacrygic",
	Aerygic:   "Aerygic",
	Dathygic:  "Dathygic",
	Boptygic:  "Boptygic",

	// 9.8 Epyrygic scale
	Epyrygic:  "Epyrygic",
	Aeradygic: "Aeradygic",
	Staptygic: "Staptygic",
	Danygic:   "Danygic",
	Goptygic:  "Goptygic",
	Epocrygic: "Epocrygic",
	Rocrygic:  "Rocrygic",
	Zyrygic:   "Zyrygic",
	Sadygic:   "Sadygic",

	// 9.9 Aeolorygic scale
	Aeolorygic: "Aeolorygic",
	Thydygic:   "Thydygic",
	Gycrygic:   "Gycrygic",
	Lyrygic:    "Lyrygic",
	Modygic:    "Modygic",
	Katodygic:  "Katodygic",
	Moptygic:   "Moptygic",
	Ionocrygic: "Ionocrygic",
	Gocrygic:   "Gocrygic",

	// 9.10 Manygic scale
	Manygic:    "Manygic",
	Polygic:    "Polygic",
	Stanygic:   "Stanygic",
	Thaptygic:  "Thaptygic",
	Eporygic:   "Eporygic",
	Aerynygic:  "Aerynygic",
	Thyptygic:  "Thyptygic",
	Ionogygic:  "Ionogygic",
	Aeolarygic: "Aeolarygic",

	// 9.11 Sathygic scale
	Sathygic:   "Sathygic",
	Ladygic:    "Ladygic",
	Sarygic:    "Sarygic",
	Thacrygic:  "Thacrygic",
	Aeolynygic: "Aeolynygic",
	Thadygic:   "Thadygic",
	Lynygic:    "Lynygic",
	Doptygic:   "Doptygic",
	Ionilygic:  "Ionilygic",

	// 9.12 Phrygic scale
	Phrygic:   "Phrygic",
	Aeranygic: "Aeranygic",
	Dothygic:  "Dothygic",
	Lydygic:   "Lydygic",
	Stadygic:  "Stadygic",
	Byptygic:  "Byptygic",
	Stodygic:  "Stodygic",
	Zynygic:   "Zynygic",
	Lonygic:   "Lonygic",

	// 9.13 Zothygic scale
	Zothygic:    "Zothygic",
	Aeolathygic: "Aeolathygic",
	Aeolocrygic: "Aeolocrygic",
	Phroptygic:  "Phroptygic",
	Kodygic:     "Kodygic",
	Eparygic:    "Eparygic",
	Ionygic:     "Ionygic",
	Gyptygic:    "Gyptygic",
	Aerythygic:  "Aerythygic",

	// 9.14 Aeolacrygic scale
	Aeolacrygic: "Aeolacrygic",
	Raptygic:    "Raptygic",
	Gythygic:    "Gythygic",
	Pyrygic:     "Pyrygic",
	Rycrygic:    "Rycrygic",
	Phrathygic:  "Phrathygic",
	Badygic:     "Badygic",
	Phrocrygic:  "Phrocrygic",
	Starygic:    "Starygic",

	// 9.15 Kyrygic scale
	Kyrygic:    "Kyrygic",
	Sythygic:   "Sythygic",
	Katycrygic: "Katycrygic",
	Tharygic:   "Tharygic",
	Sylygic:    "Sylygic",
	Lothygic:   "Lothygic",
	Darygic:    "Darygic",
	Monygic:    "Monygic",
	Styrygic:   "Styrygic",

	// 9.16 Porygic scale
	Porygic:    "Porygic",
	Mocrygic:   "Mocrygic",
	Aeolyrygic: "Aeolyrygic",
	Barygic:    "Barygic",
	Katarygic:  "Katarygic",
	Aerocrygic: "Aerocrygic",
	Zanygic:    "Zanygic",
	Aeolonygic: "Aeolonygic",
	Aeolanygic: "Aeolanygic",

	// 9.17 Kaptygic scale
	Kaptygic:  "Kaptygic",
	Sacrygic:  "Sacrygic",
	Padygic:   "Padygic",
	Epilygic:  "Epilygic",
	Kynygic:   "Kynygic",
	Stophygic: "Stophygic",
	Ionidygic: "Ionidygic",
	Stonygic:  "Stonygic",
	Stalygic:  "Stalygic",

	// 9.18 Koptygic scale
	Koptygic: "Koptygic",
	Raphygic: "Raphygic",
	Zycrygic: "Zycrygic",
	Mycrygic: "Mycrygic",
	Laptygic: "Laptygic",
	Pylygic:  "Pylygic",
	Rodygic:  "Rodygic",
	Epolygic: "Epolygic",
	Epidygic: "Epidygic",

	// 9.19 Phronygic scale
	Phronygic: "Phronygic",
	Stynygic:  "Stynygic",
	Zydygic:   "Zydygic",

	// 10.1 Aerycryllian scale
	Aerycryllian: "Aerycryllian",
	Gadyllian:    "Gadyllian",
	Solyllian:    "Solyllian",
	Zyphyllian:   "Zyphyllian",
	Garyllian:    "Garyllian",
	Soryllian:    "Soryllian",
	Godyllian:    "Godyllian",
	Epityllian:   "Epityllian",
	Ionyllian:    "Ionyllian",
	Aeoryllian:   "Aeoryllian",

	// 10.2 Katoryllian scale
	Katoryllian: "Katoryllian",
	Dodyllian:   "Dodyllian",
	Zogyllian:   "Zogyllian",
	Madyllian:   "Madyllian",
	Dycryllian:  "Dycryllian",
	Aeogyllian:  "Aeogyllian",
	Dydyllian:   "Dydyllian",
	Thogyllian:  "Thogyllian",
	Rygyllian:   "Rygyllian",
	Bathyllian:  "Bathyllian",

	// 10.3 Sydyllian scale
	Sydyllian:    "Sydyllian",
	Katogyllian:  "Katogyllian",
	Mixodyllian:  "Mixodyllian",
	Aeradyllian:  "Aeradyllian",
	Ryptyllian:   "Ryptyllian",
	Loptyllian:   "Loptyllian",
	Kataphyllian: "Kataphyllian",
	Phradyllian:  "Phradyllian",
	Dagyllian:    "Dagyllian",
	Katyllian:    "Katyllian",

	// 10.4 Gothyllian scale
	Gothyllian:  "Gothyllian",
	Lythyllian:  "Lythyllian",
	Bacryllian:  "Bacryllian",
	Aerygyllian: "Aerygyllian",
	Dathyllian:  "Dathyllian",
	Boptyllian:  "Boptyllian",
	Bagyllian:   "Bagyllian",
	Mathyllian:  "Mathyllian",
	Styptyllian: "Styptyllian",
	Zolyllian:   "Zolyllian",

	// 10.5 Staptyllian scale
	Staptyllian: "Staptyllian",
	Danyllian:   "Danyllian",
	Goptyllian:  "Goptyllian",
	Epocryllian: "Epocryllian",
	Rocryllian:  "Rocryllian",
	Zyryllian:   "Zyryllian",
	Sagyllian:   "Sagyllian",
	Epinyllian:  "Epinyllian",
	Katagyllian: "Katagyllian",
	Ragyllian:   "Ragyllian",

	// 10.6 Thydyllian
	Thydyllian:  "Thydyllian",
	Epiryllian:  "Epiryllian",
	Lyryllian:   "Lyryllian",
	Mogyllian:   "Mogyllian",
	Katodyllian: "Katodyllian",

	// 11.1 Aerycratic scale
	Aerycratic: "Aerycratic",
	Monatic:    "Monatic",
	Solatic:    "Solatic",
	Zylatic:    "Zylatic",
	Mixolatic:  "Mixolatic",
	Soratic:    "Soratic",
	Godatic:    "Godatic",
	Eptatic:    "Eptatic",
	Ionatic:    "Ionatic",
	Aeolatic:   "Aeolatic",
	Thydatic:   "Thydatic",

	// 12 notes
	Chromatic: "Chromatic",
}

// String retuns stringified mode name
func (t Type) String() string {
	return mapModeToString[t]
}

// GoString satisfies GoStringer
func (t Type) GoString() string {
	return t.String()
}
