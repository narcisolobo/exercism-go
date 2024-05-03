package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newScores := make(map[string]int)

	for score, letters := range in {
		for _, letter := range letters {
			newScores[strings.ToLower(letter)] = score
		}
	}
	return newScores
}
