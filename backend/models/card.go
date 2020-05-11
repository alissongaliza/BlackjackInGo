package models

import "github.com/alissongaliza/BlackjackInGo/utils"

type Card struct {
	Suit     utils.SuitType
	Name     string
	IsNumber bool
	IsFaceUp bool
}

func (card Card) Value(hand Hand) int {
	if !card.IsFaceUp {
		return 0
	}

	if card.IsNumber {
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
