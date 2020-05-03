package models

import (
	"fmt"
)

func StartGame(gameId int) (game Game) {
	fmt.Println("StartGame reached", gameId)
	game = GetGameDb().Get(gameId)
	game.User.Hit(game.Id, true)
	game.User.Hit(game.Id, true)
	//set the dealer's last given card face down and recalculate score
	dealerHand := game.Dealer.Hand
	dealerHand.Score -= dealerHand.Cards[1].value(*dealerHand)
	dealerHand.Cards[1].IsFaceUp = false
	GetGameDb().Update(game)
	return
}
