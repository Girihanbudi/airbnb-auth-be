package codegenerator

import "testing"

func TestRandomEncodedBytes(t *testing.T) {
	result := RandomEncodedBytes(16)
	if result == "" {
		t.Error("failed to generate code")
	}
}
