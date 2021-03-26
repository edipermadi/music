-- schema for postgres

CREATE table pitches
(
    id        BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number    INTEGER          NOT NULL,
    octave    INTEGER          NOT NULL,
    frequency DOUBLE PRECISION NOT NULL,
    label     TEXT             NOT NULL,
    name      TEXT             NOT NULL
);

CREATE
INDEX ON pitches (number);

CREATE
INDEX ON pitches (octave);

CREATE TABLE accidentals
(
    id           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    displacement INT  NOT NULL,
    label        TEXT NOT NULL,
    name         TEXT NOT NULL
);

CREATE TABLE notes
(
    id            BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number        INTEGER NOT NULL,
    accidental_id BIGINT  NOT NULL REFERENCES accidentals (id),
    label         TEXT    NOT NULL,
    name          TEXT    NOT NULL
);

CREATE
INDEX ON notes (number);

CREATE TABLE note_pitches
(
    id       BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    octave   INTEGER NOT NULL,
    note_id  BIGINT  NOT NULL REFERENCES notes (id),
    pitch_id BIGINT  NOT NULL REFERENCES pitches (id)
);

CREATE
UNIQUE INDEX ON note_pitches (note_id, pitch_id);

CREATE TABLE scales
(
    id            BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    cardinality   INTEGER NOT NULL,
    transposition JSONB   NOT NULL,
    label         TEXT    NOT NULL,
    name          TEXT    NOT NULL
);
