package blackjack
import "strings"

var cardValues = map[string]int{
    "ace": 11, "eight": 8, "two": 2, "nine": 9,
    "three": 3, "ten": 10, "four": 4, "jack": 10,
    "five": 5, "queen": 10, "six": 6, "king": 10,
    "seven": 7,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    normalized := strings.ToLower(card)
    return cardValues[normalized]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	v1 := ParseCard(card1)
	v2 := ParseCard(card2)
	vDealer := ParseCard(dealerCard)
	sum := v1 + v2

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sum == 21:
		if vDealer < 10 {
			return "W"
		}
		return "S"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16:
		if vDealer >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
