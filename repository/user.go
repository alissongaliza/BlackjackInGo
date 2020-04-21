package models

type User interface {
	Hit(gameId int, faceUp bool)
	Stand(gameId int)
	DoubleDown(gameId int)
}
