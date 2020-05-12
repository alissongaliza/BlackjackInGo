package game

import (
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type UseCase interface {
	CreateGame(user models.User, dif utils.Difficulty, bet int) (models.Game, error)
	StartNewGame(game models.Game) (models.Game, error)
	ContinueGame(game models.Game) (models.Game, error)
}
