package data

// Pitch is a database model
type Pitch struct {
	ID        int     `json:"id" db:"id"`
	Octave    int     `json:"octave" db:"octave"`
	Frequency float32 `json:"frequency" db:"frequency"`
	Label     string  `json:"label" db:"label"`
	Name      string  `json:"name" db:"name"`
}

// Accidental is a database model
type Accidental struct {
	ID           int    `json:"id" db:"id"`
	Displacement int    `json:"displacement" db:"displacement"`
	Label        string `json:"label" db:"label"`
	Name         string `json:"name" db:"name"`
}

// Note is a database model
type Note struct {
	ID           int    `json:"id" db:"id"`
	AccidentalID int    `json:"accidental_id" db:"accidental_id"`
	Label        string `json:"label" db:"label"`
	Name         string `json:"name" db:"name"`
}

// Chord is a database model
type Chord struct {
	ID        int    `json:"id" db:"id"`
	Root      Note   `json:"root" db:"-"`
	Inversion int    `json:"inversion" db:"inversion"`
	Notes     []Note `json:"notes" db:"-"`
	Checksum  int    `json:"checksum" db:"checksum"`
	Label     string `json:"label" db:"-"`
	Name      string `json:"name" db:"-"`
}

// Mode is a model
type Mode struct {
	ID            int    `json:"id" db:"id"`
	Scale         Scale  `json:"scale" db:"-"`
	Transposition []int  `json:"transposition" db:""`
	Shift         int    `json:"shift" db:"-"`
	Tonic         Note   `json:"tonic" db:"-"`
	Notes         []Note `json:"notes" db:"-"`
	Checksum      int    `json:"checksum" db:"checksum"`
	Label         string `json:"label" db:"-"`
	Name          string `json:"name" db:"-"`
}

// Scale is a model
type Scale struct {
	ID            int    `json:"id" db:"id"`
	Transposition []int  `json:"transposition" db:""`
	Cardinality   int    `json:"cardinality" db:"cadinality"`
	Label         string `json:"label" db:"-"`
	Name          string `json:"name" db:"-"`
}
