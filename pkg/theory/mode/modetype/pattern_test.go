package modetype_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/mode/modetype"
)

func TestMode_IntervalPattern(t *testing.T) {
	for _, givenType := range modetype.AllTypes() {
		pattern := givenType.IntervalPattern()
		t.Logf("mode %s has interval pattern %#v", givenType, pattern)
	}
}
