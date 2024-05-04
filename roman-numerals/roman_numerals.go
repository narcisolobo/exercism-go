package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "error", errors.New("invalid input")
	}

	var b strings.Builder

	if input > 1000 {
		m := input / 1000
		for i := 0; i < m; i++ {
			b.WriteRune('M')
		}
		input -= 1000 * m
	}

	if input > 100 {
		c := input / 100
		switch {
		case c == 9:
			b.WriteString("CM")
		case c >= 5 && c < 9:
			b.WriteRune('D')
			for i := 0; i < c%5; i++ {
				b.WriteRune('C')
			}
		case c == 4:
			b.WriteString("CD")
		case c >= 1 && c < 4:
			for i := 0; i < c; i++ {
				b.WriteRune('C')
			}
		}
		input -= 100 * c
	}

	if input > 10 {
		x := input / 10
		switch {
		case x == 9:
			b.WriteString("XC")
		case x >= 5 && x < 9:
			b.WriteRune('L')
			for i := 0; i < x%5; i++ {
				b.WriteRune('X')
			}
		case x == 4:
			b.WriteString("XL")
		case x >= 1 && x < 4:
			for i := 0; i < x; i++ {
				b.WriteRune('X')
			}
		}
		input -= 10 * x
	}

	if input >= 1 {
		switch {
		case input == 9:
			b.WriteString("IX")
		case input >= 5 && input < 9:
			b.WriteRune('V')
			for i := 0; i < input%5; i++ {
				b.WriteRune('I')
			}
		case input == 4:
			b.WriteString("IV")
		case input >= 1 && input < 4:
			for i := 0; i < input; i++ {
				b.WriteRune('I')
			}
		default:
			return "error", errors.New("invalid input")
		}
	}
	return b.String(), nil
}
