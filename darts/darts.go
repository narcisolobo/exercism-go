package darts

import (
	"math"
)

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	var score int

	switch {
	case distance > 10:
		score = 0
	case distance > 5 && distance <= 10:
		score = 1
	case distance > 1 && distance <= 5:
		score = 5
	case distance >= 0 && distance <= 1:
		score = 10
	}

	return score
}
