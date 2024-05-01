package isogram

import (
	"slices"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	okay := []uint8{'-', ' '}
	result := true
	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				if slices.Contains(okay, word[i]) {
					continue
				}
				result = false
			}
		}
	}
	return result
}
