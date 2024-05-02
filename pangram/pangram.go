package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	lowercase := ""
	for r := 'a'; r <= 'z'; r++ {
		lowercase += string(r)
	}

	for _, char := range input {
		if strings.Contains(lowercase, string(char)) {
			lowercase = strings.ReplaceAll(lowercase, string(char), "")
		} else {
			continue
		}
	}
	return len(lowercase) == 0
}
