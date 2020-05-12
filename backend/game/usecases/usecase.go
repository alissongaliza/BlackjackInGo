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
	userRepo      user.Repository
}

func NewGameUsecase(gameRepo game.Repository, dealerUsecase dealer.UseCase,
	userUsecase user.UseCase, userRepo user.Repository) game.UseCase {
	return &gameUsecase{
		gameRepo:      gameRepo,
		dealerUsecase: dealerUsecase,
		userUsecase:   userUsecase,
		userRepo:      userRepo,
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

func (guc gameUsecase) StartNewGame(game models.Game) models.Game {
	game.GameState = utils.Playing
	game.User.Chips -= game.Bet
	// due to not handling user db and game db the right way
	game = guc.gameRepo.UpdateGame(game)
	game.User = guc.userRepo.UpdateUser(game.User)
	game = guc.userUsecase.Hit(game, true)
	game = guc.userUsecase.Hit(game, true)
	game = guc.dealerUsecase.Hit(game, true)
	game = guc.dealerUsecase.Hit(game, false)
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
