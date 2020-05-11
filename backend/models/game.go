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

func (game *Game) CalculatePayouts() {
	user := game.User
	winnings := 0
	switch game.GameState {
	case utils.Won:
		{
			if game.LastUserAction == utils.DoubleDown {
				winnings = game.Bet * 4
			} else {
				winnings = game.Bet * 2
			}
		}
	case utils.Drew:
		{
			if user.Hand.Score == 21 {
				winnings = int(float64(game.Bet) * 1.5)
			} else {
				winnings = game.Bet
			}
		}
	}
	user.Chips += winnings
	game.Payout = winnings
}
