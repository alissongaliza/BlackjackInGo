package models

type User interface {
	hit(gameId int, faceUp bool)
	stand(gameId int)
	doubleDown(gameId int)
}
