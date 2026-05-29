package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int;
    
	switch card {
        case "ace":
        	value = 11
        	break
        case "two":
        	value = 2
        	break
        case "three":
        	value = 3
        	break
        case "four":
        	value = 4
        	break
        case "five":
        	value = 5
        	break
        case "six":
        	value = 6
        	break
        case "seven":
        	value = 7
        	break
        case "eight":
        	value = 8
        	break
        case "nine":
        	value = 9
        	break
        case "ten", "jack", "queen", "king":
        	value = 10
        	break
        default:
        	value = 0
        	break
    }

    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Val := ParseCard(card1)
    card2Val := ParseCard(card2)
    dealerCardVal := ParseCard(dealerCard)

    cardSum := card1Val + card2Val
    
	if card1Val == 11 && card2Val == 11 {
        return "P"
    }

    if cardSum == 21 {
        if dealerCardVal < 10 {
        	return "W"
        }

        return "S"
    }

    if cardSum >= 17 && cardSum <= 20 {
        return "S"
    }

    if cardSum >= 12 && cardSum <= 16 {
        if dealerCardVal <= 6 {
            return "S"
        }

        return "H"
    }

    return "H"
}
