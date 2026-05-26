package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, exists := cards[card]

	if exists {
		return value
	}

	return 0
}
// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerTotal := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
    case card1 == "ace" && card2 == "ace":
		return "P"

	case playerTotal == 21:
		switch dealerValue {
		case 10, 11:
			return "S"
		default:
			return "W"
		}

	case playerTotal >= 17 && playerTotal <= 20:
		return "S"

	case playerTotal <= 11:
		return "H"

	case playerTotal >= 12 && playerTotal <= 16:
		switch {
		case dealerValue >= 7:
			return "H"
		default:
			return "S"
		}
	}

	return "S"
}
