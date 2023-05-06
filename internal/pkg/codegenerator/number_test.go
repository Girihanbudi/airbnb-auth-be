package codegenerator

import (
	"regexp"
	"testing"
)

func TestRandomEncodedNumbers6DigitNumber(t *testing.T) {
	result := RandomEncodedNumbers(6)
	if result == "" {
		t.Error("failed to generate code")
	}

	if len(result) != 6 {
		t.Error("expected to have 6 digit numbers")
	}
}

func TestRandomEncodedNumbersDigitOnlyResult(t *testing.T) {
	result := RandomEncodedNumbers(6)
	if result == "" {
		t.Error("failed to generate code")
	}

	re := regexp.MustCompile(`\d`)

	if !re.MatchString(result) {
		t.Error("expected to have digit only result")
	}
}
