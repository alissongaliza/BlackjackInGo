package user

import "github.com/alissongaliza/BlackjackInGo/backend/models"

type UseCase interface {
	Hit(game models.Game, faceUp bool) (models.Game, error)
	Stand(game models.Game) models.Game
	DoubleDown(game models.Game) (models.Game, error)
	IsUserValid(user models.User) (bool, error)
	CreateUser(name string, age int) models.User
}
