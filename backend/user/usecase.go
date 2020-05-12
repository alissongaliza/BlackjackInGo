package user

import "github.com/alissongaliza/BlackjackInGo/backend/models"

type UseCase interface {
	Hit(game models.Game, faceUp bool) models.Game
	Stand(game models.Game) models.Game
	DoubleDown(game models.Game) models.Game
	IsUserValid(user models.User) bool
	CreateUser(name string, age int) models.User
}
