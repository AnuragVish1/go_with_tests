package propertybased

import (
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	HinduArabic uint16
	Roman       string
}{
	{HinduArabic: 1, Roman: "I"},
	{HinduArabic: 2, Roman: "II"},
	{HinduArabic: 3, Roman: "III"},
	{HinduArabic: 4, Roman: "IV"},
	{HinduArabic: 5, Roman: "V"},
	{HinduArabic: 6, Roman: "VI"},
	{HinduArabic: 7, Roman: "VII"},
	{HinduArabic: 8, Roman: "VIII"},
	{HinduArabic: 9, Roman: "IX"},
	{HinduArabic: 10, Roman: "X"},
	{HinduArabic: 14, Roman: "XIV"},
	{HinduArabic: 18, Roman: "XVIII"},
	{HinduArabic: 20, Roman: "XX"},
	{HinduArabic: 39, Roman: "XXXIX"},
	{HinduArabic: 40, Roman: "XL"},
	{HinduArabic: 47, Roman: "XLVII"},
	{HinduArabic: 49, Roman: "XLIX"},
	{HinduArabic: 50, Roman: "L"},
	{HinduArabic: 100, Roman: "C"},
	{HinduArabic: 90, Roman: "XC"},
	{HinduArabic: 400, Roman: "CD"},
	{HinduArabic: 500, Roman: "D"},
	{HinduArabic: 900, Roman: "CM"},
	{HinduArabic: 1000, Roman: "M"},
	{HinduArabic: 1984, Roman: "MCMLXXXIV"},
	{HinduArabic: 3999, Roman: "MMMCMXCIX"},
	{HinduArabic: 2014, Roman: "MMXIV"},
	{HinduArabic: 1006, Roman: "MVI"},
	{HinduArabic: 798, Roman: "DCCXCVIII"},
}

func TestRoman(t *testing.T) {

	for _, test := range cases {
		t.Run("Converting", func(t *testing.T) {
			got, _ := ToRoman(test.HinduArabic)

			if got != test.Roman {
				t.Errorf("Got %q want %q", got, test.Roman)
			}
		})
	}

}

func TestToNum(t *testing.T) {
	for _, test := range cases[:4] {
		t.Run("converting roman number to modern number", func(t *testing.T) {
			got, _ := ToNum(test.Roman)
			want := test.HinduArabic

			if got != want {
				t.Errorf("Got %v want %v", got, want)
			}
		})
	}
}

func TestPropertyBased(t *testing.T) {
	validateFunc :=
		func(num uint16) bool {
			t.Log(num)
			roman, err := ToRoman(num)
			if err != nil {
				log.Print("Value higher than the max limit")
			}
			hinduArabic, err := ToNum(roman)
			if err != nil {
				t.Error(err.Error())
			}

			return num == hinduArabic
		}
	if err := quick.Check(validateFunc, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("Not the same bro something is wrong", err)
	}
}
