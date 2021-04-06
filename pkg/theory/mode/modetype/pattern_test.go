package modetype_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/mode/modetype"
)

func TestMode_Transposition(t *testing.T) {
	for _, givenType := range modetype.AllTypes() {
		pattern := givenType.Transposition()
		t.Logf("mode %s has transposition %#v", givenType, pattern)
	}
}
