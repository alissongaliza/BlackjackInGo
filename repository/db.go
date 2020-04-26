package models

import "fmt"

type Actors string

type base interface {
	Save(model base) base
	Get(modelId int)
	Update(model base) base
}

func StartGame(gameId int) (game Game) {
	fmt.Println("StartGame reached", gameId)
	game = GetGameDb().Get(gameId)
	game.House.Hit(game.Id, true)
	game.Player.Hit(game.Id, true)
	game.House.Hit(game.Id, false)
	game.Player.Hit(game.Id, true)
	GetGameDb().Update(game)
	return
}
