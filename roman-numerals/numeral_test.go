package romannumerals

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted into I", 1, "I"},
		{"2 gets converted into II", 2, "II"},
		{"3 gets converted into III", 3, "III"},
		{"4 gets converted into IV", 4, "IV"},
		{"5 gets converted into V", 5, "V"},
		{"6 gets converted into VI", 6, "VI"},
		{"7 gets converted into VII", 7, "VII"},
		{"8 gets converted into VIII", 8, "VIII"},
		{"9 gets converted into VIII", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf("got %q, wante %q", got, want)
			}
		})
	}

}
