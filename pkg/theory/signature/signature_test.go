package signature_test

import (
	"testing"

	"github.com/edipermadi/music/pkg/theory/signature"
	"github.com/stretchr/testify/require"
)

func TestSignature_String(t *testing.T) {
	for _, givenSignature := range signature.AllSignatures() {
		require.NotEmpty(t, givenSignature.String())
	}
}

func TestSignature_Name(t *testing.T) {
	for _, givenSignature := range signature.AllSignatures() {
		require.NotEmpty(t, givenSignature.Name())
	}
}
