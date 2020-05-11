package usecases

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/backend/dealer"
	"github.com/alissongaliza/BlackjackInGo/backend/game"
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/backend/user"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type gameUsecase struct {
	gameRepo      game.Repository
	dealerUsecase dealer.UseCase
	userUsecase   user.UseCase
}

func NewGameUsecase(gameRepo game.Repository, dealerUsecase dealer.UseCase, userUsecase user.UseCase) game.UseCase {
	return &gameUsecase{
		gameRepo:      gameRepo,
		dealerUsecase: dealerUsecase,
		userUsecase:   userUsecase,
	}
}

func (guc gameUsecase) CreateGame(user models.User, dif utils.Difficulty, bet int) models.Game {
	id := guc.gameRepo.GetNextValidGameId()
	cards := guc.gameRepo.CreateDeck()
	dealer := guc.dealerUsecase.CreateDealer(dif)
	game := models.Game{
		Id: id, Bet: bet, User: user, Dealer: dealer,
		GameState: utils.Setup, Cards: cards, LastDealerAction: utils.NoAction,
		LastUserAction: utils.NoAction, Payout: 0,
	}
	return guc.gameRepo.CreateGame(game)
}

func (guc gameUsecase) GetGame(gameId int) models.Game {
	return guc.gameRepo.GetGame(gameId)
}

func (guc gameUsecase) ListGame(userId int) []models.Game {
	return guc.gameRepo.ListGame(userId)
}

func (guc gameUsecase) UpdateGame(game models.Game) models.Game {
	return guc.gameRepo.UpdateGame(game)
}

func (guc gameUsecase) IsGameValid(gameId int) bool {
	game := guc.gameRepo.GetGame(gameId)
	return game.Id > 0
}

func (guc gameUsecase) CreateHand() models.Hand {
	return guc.CreateHand()
}

func (guc gameUsecase) StartNewGame(game models.Game) models.Game {
	game.GameState = utils.Playing
	game.User.Chips -= game.Bet
	// due to not handling user db and game db the right way
	game = guc.gameRepo.UpdateGame(game)
	game.User = guc.userUsecase.UpdateUser(game.User)
	game = guc.userUsecase.Hit(game.Id, true)
	game = guc.userUsecase.Hit(game.Id, true)
	game = guc.dealerUsecase.Hit(game.Id, true)
	game = guc.dealerUsecase.Hit(game.Id, false)
	//set the dealer's last given card face down and recalculate score
	dealerHand := &game.Dealer.Hand
	dealerHand.Cards[1].IsFaceUp = false
	dealerHand.Score = dealerHand.RecalculateScore()
	return guc.gameRepo.UpdateGame(game)
}

func (guc gameUsecase) ContinueGame(game models.Game) models.Game {
	if game.GameState != utils.Playing {
		panic(fmt.Sprintf("Game is already over. User %s", game.GameState))
	}
	if game.LastUserAction == utils.Hit {
		// means that its the player's turn
		return game
	} else {
		return guc.dealerUsecase.AutoPlay(game)
	}
}
