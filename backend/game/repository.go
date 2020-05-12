package game

import (
	models "github.com/alissongaliza/BlackjackInGo/backend/models"
)

type Repository interface {
	CreateGame(game models.Game) models.Game
	GetGame(gameId int) (models.Game, error)
	ListGame(userId int) []models.Game
	UpdateGame(game models.Game) models.Game
	CreateHand() models.Hand
	CreateDeck() []models.Card
	GetNextValidGameId() int
}
