package signature_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/signature"
	"github.com/stretchr/testify/require"
)

func TestSignature_Number(t *testing.T) {
	for _, givenSignature := range signature.AllSignatures() {
		computedNumber := givenSignature.Number()
		require.NotEmpty(t, computedNumber)
		t.Logf("signature %s has number %d", givenSignature.Name(), computedNumber)
	}
}
