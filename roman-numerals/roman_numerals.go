package romannumerals

import (
	"errors"
	"strings"
)

// ToRomanNumeral converts an integer to a Roman numeral.
func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input out of range (1-3999)")
	}

	var result strings.Builder

	// Define the values and corresponding symbols for Roman numerals.
	romanSymbols := []struct {
		value  int
		symbol string
	}{
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

	// Iterate over the values and symbols, appending symbols to the result.
	for _, rs := range romanSymbols {
		for input >= rs.value {
			result.WriteString(rs.symbol)
			input -= rs.value
		}
	}

	return result.String(), nil
}
