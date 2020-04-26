package models

type User interface {
	Hit(gameId int, faceUp bool) (game Game)
	Stand(gameId int) (game Game)
	DoubleDown(gameId int) (game Game)
}
