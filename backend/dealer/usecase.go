package dealer

import (
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type UseCase interface {
	CreateDealer(dif utils.Difficulty) models.Dealer
	AutoPlay(models.Game) models.Game
	Hit(game models.Game, faceUp bool) models.Game
	Stand(game models.Game) models.Game
}
