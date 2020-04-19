package models

import (
	"strconv"
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
}

type Hand struct {
	Cards []Card
	Score int
}

func (card Card) value(hand Hand) int {
	if card.isNumber {
		// convert from string to int
		if s, err := strconv.Atoi(card.Name); err == nil {
			return s
		}
	}

	if card.Name == "ace" {
		if hand.Score+11 <= 21 {
			return 11
		}
		return 1

	}

	return 10
}
