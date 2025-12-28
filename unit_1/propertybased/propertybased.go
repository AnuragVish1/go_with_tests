package propertybased

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

const MaxLimitNum uint16 = 3999

var romanNum = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {

}

func ToRoman(num uint16) (string, error) {

	if num > MaxLimitNum {
		return "", errors.New("overflow")
	}

	var result strings.Builder

	for _, roman_num := range romanNum {
		for num >= roman_num.Value {
			result.WriteString(roman_num.Symbol)
			num -= roman_num.Value
		}
	}

	return result.String(), nil
}

func ToNum(romanNumeral string) (uint16, error) {
	var total uint16 = 0

	if !isConsecutive(romanNumeral) {
		return 0, errors.New("found 3 letters to be consecutive")
	}

	for _, value := range romanNum {
		for strings.HasPrefix(romanNumeral, value.Symbol) {
			total += value.Value
			romanNumeral = strings.TrimPrefix(romanNumeral, value.Symbol)
		}
	}
	return total, nil
}

func isConsecutive(roman string) bool {
	founeded := map[rune]int{}

	for _, value := range roman {
		founeded[value]++
		if founeded[value] >= 3 {
			return false
		}
	}
	return true

}
