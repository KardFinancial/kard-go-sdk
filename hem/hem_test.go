package hem

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testVector struct {
	Name          string `json:"name"`
	Input         string `json:"input"`
	Normalized    string `json:"normalized"`
	SHA256Hex     string `json:"sha256_hex"`
	UID2Validated bool   `json:"uid2_validated,omitempty"`
}

func loadTestVectors(t *testing.T) []testVector {
	t.Helper()
	data, err := os.ReadFile("test_vectors.json")
	require.NoError(t, err)
	var vs []testVector
	require.NoError(t, json.Unmarshal(data, &vs))
	require.NotEmpty(t, vs, "test_vectors.json yielded no vectors")
	return vs
}

func TestNormalizeEmail(t *testing.T) {
	for _, v := range loadTestVectors(t) {
		t.Run(v.Name, func(t *testing.T) {
			got, err := normalizeEmail(v.Input)
			require.NoError(t, err)
			require.Equal(t, v.Normalized, got)
		})
	}
}

func TestGenerateHEM(t *testing.T) {
	for _, v := range loadTestVectors(t) {
		t.Run(v.Name, func(t *testing.T) {
			got, err := GenerateHEM(v.Input)
			require.NoError(t, err)
			require.Equal(t, v.SHA256Hex, got)
		})
	}
}

func TestGenerateHEMRejectsInvalid(t *testing.T) {
	cases := map[string]string{
		"empty":          "",
		"whitespaceOnly": " ",
		"atOnly":         "@",
		"noAt":           "user",
		"emptyDomain":    "user@",
		"emptyLocal":     "@domain.com",
		"multipleAt":     "a@b@c.com",
	}
	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := GenerateHEM(in)
			require.Error(t, err)
		})
	}
}
