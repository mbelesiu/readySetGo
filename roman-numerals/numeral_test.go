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
