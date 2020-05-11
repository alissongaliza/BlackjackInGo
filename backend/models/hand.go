package models

type Hand struct {
	Cards []Card
	Score int
}

func (hand Hand) RecalculateScore() int {
	score := 0
	for _, card := range hand.Cards {
		score += card.Value(hand)
	}
	return score
}
