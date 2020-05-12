package user

import "github.com/alissongaliza/BlackjackInGo/backend/models"

type UseCase interface {
	Hit(gameId int, faceUp bool) (game models.Game)
	Stand(gameId int) (game models.Game)
	DoubleDown(gameId int) (game models.Game)
	IsUserValid(userId int) bool
	CreateUser(name string, age int) models.User
}
