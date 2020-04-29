package models

type Player interface {
	Hit(gameId int, faceUp bool) (game Game)
	Stand(gameId int) (game Game)
}

type RealPlayer interface {
	Player
	DoubleDown(gameId int) (game Game)
}
