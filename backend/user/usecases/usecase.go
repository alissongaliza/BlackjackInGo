package usecases

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/backend/dealer"
	"github.com/alissongaliza/BlackjackInGo/backend/game"
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/backend/user"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type userUsecase struct {
	gameRepo      game.Repository
	userRepo      user.Repository
	dealerUsecase dealer.UseCase
}

func NewUserUsecase(gameRepo game.Repository, userRepo user.Repository, dealerUsecase dealer.UseCase) user.UseCase {
	return &userUsecase{
		gameRepo:      gameRepo,
		userRepo:      userRepo,
		dealerUsecase: dealerUsecase,
	}
}

func (uuc userUsecase) CreateUser(name string, age int) models.User {
	id := uuc.userRepo.GetNextValidUserId()
	player := models.Player{Hand: uuc.gameRepo.CreateHand()}
	user := models.User{Age: age, Name: name, Id: id, Chips: 100, Player: player}
	return uuc.userRepo.CreateUser(user)
}

func (uuc userUsecase) IsUserValid(user models.User) bool {
	var err error
	if user, err = uuc.userRepo.GetUser(user.Id); err != nil {
		return false
	}
	return user.Age >= 18
}

func (uuc userUsecase) Hit(game models.Game, faceUp bool) (models.Game, error) {
	// fmt.Printf("user %d hits!", game.User.Id)
	if game.GameState != utils.Playing {
		return models.Game{}, fmt.Errorf("Game is already over. User %s", game.GameState)
	}
	//pop an element from the cards slice
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]
	card.IsFaceUp = faceUp
	// assign the new cards to the user's hand
	hand := &game.User.Hand
	hand.Cards = append(hand.Cards, card)
	if card.IsFaceUp {
		hand.Score += card.Value(*hand)
	}
	// check if burst happened
	if hand.Score > 21 {
		game.GameState = utils.Lost
	}
	// fmt.Println("user hand", game.User.Hand, game.GameState)
	return uuc.gameRepo.UpdateGame(game), nil
}

func (uuc userUsecase) Stand(game models.Game) models.Game {
	// fmt.Printf("user %d stands!", game.User.Id)
	game.LastUserAction = utils.Stand
	uuc.gameRepo.UpdateGame(game)
	//call dealers turn
	return uuc.dealerUsecase.AutoPlay(game)
}

func (uuc userUsecase) DoubleDown(game models.Game) (models.Game, error) {
	// fmt.Println("user doubleDowns!")
	if len(game.User.Hand.Cards) != 2 {
		return models.Game{}, fmt.Errorf(`User can't Double Down. This move is only available when he only has the first 2 starting cards`)
	} else if game.User.Chips < game.Bet {
		return models.Game{}, fmt.Errorf(`User can't Double Down. His current chips balance is lower than necessary`)
	}
	game.User.Chips -= game.Bet
	game.Bet += game.Bet
	var err error
	if game, err = uuc.Hit(game, true); err != nil {
		return models.Game{}, err
	}
	game.LastUserAction = utils.DoubleDown
	//call dealers turn
	return uuc.dealerUsecase.AutoPlay(game), nil
}
