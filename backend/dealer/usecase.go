package dealer

import (
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type UseCase interface {
	CreateDealer(dif utils.Difficulty) models.Dealer
	AutoPlay(models.Game) models.Game
	Hit(gameId int, faceUp bool) (game models.Game)
	Stand(gameId int) (game models.Game)
}
