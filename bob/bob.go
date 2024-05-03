package bob

import (
	"strings"
	"unicode"
)

func isYelling(remark string) bool {
	hasLetters := false
	for _, char := range remark {
		if unicode.IsLetter(char) {
			hasLetters = true
			if !unicode.IsUpper(char) {
				return false
			}
		}
	}
	return hasLetters
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var reply string
	switch {
	case len(remark) == 0:
		reply = "Fine. Be that way!"
	case remark[len(remark)-1] == '?' && !isYelling(remark):
		reply = "Sure."
	case remark[len(remark)-1] == '?' && isYelling(remark):
		reply = "Calm down, I know what I'm doing!"
	case remark[len(remark)-1] != '?' && isYelling(remark):
		reply = "Whoa, chill out!"
	default:
		reply = "Whatever."
	}
	return reply
}
