package models

import (
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type DealerActions interface {
	Play(Game) Game
}

type Dealer struct {
	DealerActions
	Player
	Difficulty utils.Difficulty
}

type EasyDealer Dealer
type BrokenDealer Dealer
