package models

import "fmt"

func StartGame(gameId int) (game Game) {
	fmt.Println("StartGame reached", gameId)
	game = GetGameDb().Get(gameId)
	game.User.Hit(game.Id, true)
	game.User.Hit(game.Id, true)
	//set the dealer's last given card face down
	dealer := game.Dealer.(*EasyDealer)
	dealer.Hand.Cards[1].isFaceUp = false
	GetGameDb().Update(game)
	return
}
