package models

import "github.com/alissongaliza/BlackjackInGo/utils"

func NewHand() Hand {
	return Hand{make([]Card, 0), 0}
}

func (card Card) value(hand Hand) int {
	if !card.IsFaceUp {
		return 0
	}

	if card.isNumber {
		return utils.StringToInt(card.Name)
	}

	if card.Name == "ACE" {
		if hand.Score+11 <= 21 {
			return 11
		}
		return 1

	}

	return 10
}

func recalculateHandScore(hand Hand) int {
	score := 0
	for _, card := range hand.Cards {
		score += card.value(hand)
	}
	return score
}
