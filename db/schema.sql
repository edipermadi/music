-- schema for postgres

CREATE table pitches
(
    id        BIGINT GENERATED ALWAYS AS IDENTITY,
    number    INTEGER          NOT NULL,
    octave    INTEGER          NOT NULL,
    frequency DOUBLE PRECISION NOT NULL,
    label     TEXT             NOT NULL,
    name      TEXT             NOT NULL
);

CREATE INDEX ON pitches (number);
CREATE INDEX ON pitches (octave);