package game

import (
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type UseCase interface {
	CreateGame(user models.User, dif utils.Difficulty, bet int) models.Game
	GetGame(gameId int) models.Game
	ListGame(userId int) []models.Game
	UpdateGame(game models.Game) models.Game
	CreateHand() models.Hand
	IsGameValid(gameId int) bool
	StartNewGame(game models.Game) models.Game
	ContinueGame(game models.Game) models.Game
	CalculatePayouts(game *models.Game)
}
