package models

import (
	"github.com/alissongaliza/BlackjackInGo/backend/utils"
)

func NewHand() Hand {
	return Hand{make([]Card, 0), 0}
}

func (card Card) value(hand Hand) int {
	if !card.isFaceUp {
		return 0
	}

	if card.isNumber {
		return utils.StringToInt(card.Name)
	}

	if card.Name == "ace" {
		if hand.Score+11 <= 21 {
			return 11
		}
		return 1

	}

	return 10
}
