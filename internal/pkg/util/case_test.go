package util

import (
	"testing"
)

func TestCaseLowerAndTitle(t *testing.T) {
	t.Log("testing lower case and title case...")
	input := "heLLo wORLD!!!"
	want := "Hello World!!!"

	result := Case(input, CaseLower, CaseTitle)
	if result != want {
		t.Error("result not match, want:", want, "but the result is:", result)
	}
}
