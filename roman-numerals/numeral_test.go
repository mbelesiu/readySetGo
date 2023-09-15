package romannumerals

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q, wante %q", got, want)
	}
}