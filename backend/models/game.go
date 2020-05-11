package models

import "github.com/alissongaliza/BlackjackInGo/utils"

type Game struct {
	Id               int
	User             User
	Dealer           Dealer
	Cards            []Card
	Bet              int
	Payout           int
	LastUserAction   utils.Action
	LastDealerAction utils.Action
	GameState        utils.GameState
}
