package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	}
	return value
}

func DealerHasAceOrFace(card string) bool {
	dealerHasAceOrFace := false
	switch ParseCard(card) {
	case 10:
		dealerHasAceOrFace = true
	case 11:
		dealerHasAceOrFace = true
	}
	return dealerHasAceOrFace
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var play string
	holding := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case holding == 22:
		play = "P"
	case holding == 21 && !DealerHasAceOrFace(dealerCard):
		play = "W"
	case holding == 21 && DealerHasAceOrFace(dealerCard):
		play = "S"
	case holding >= 17 && holding <= 20:
		play = "S"
	case holding >= 12 && holding <= 16 && dealer <= 6:
		play = "S"
	case holding >= 12 && holding <= 16 && dealer >= 7:
		play = "H"
	case holding <= 11:
		play = "H"
	}

	return play
}
