package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	score -= 10
	flScore := math.Floor(float64(score) / 2)
	return int(flScore)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := []int{}
	for i := 0; i < 4; i++ {
		rolls = append(rolls, rand.Intn(6)+1)
	}
	slices.Sort(rolls)
	rolls = slices.Delete(rolls, 0, 1)
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10,
	}
	character.Hitpoints += Modifier(character.Constitution)
	return character
}
