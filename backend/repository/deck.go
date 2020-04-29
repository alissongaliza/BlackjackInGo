package models

import (
	"github.com/alissongaliza/BlackjackInGo/backend/utils"
)

type SuitType string

const (
	hearts   SuitType = "hearts"
	spades   SuitType = "spades"
	clubs    SuitType = "clubs"
	diamonds SuitType = "diamonds"
)

type Card struct {
	Suit     SuitType
	Name     string
	isNumber bool
	isFaceUp bool
}

type Hand struct {
	Cards []Card
	Score int
}

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
