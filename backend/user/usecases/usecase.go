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
	player := models.Player{uuc.gameRepo.CreateHand()}
	user := models.User{Age: age, Name: name, Id: id, Chips: 100, Player: player}
	return uuc.userRepo.CreateUser(user)
}

func (uuc userUsecase) GetUser(userId int) models.User {
	return uuc.userRepo.GetUser(userId)
}

func (uuc userUsecase) ListUser(name string) []models.User {
	return uuc.userRepo.ListUser(name)
}

func (uuc userUsecase) UpdateUser(user models.User) models.User {
	return uuc.userRepo.UpdateUser(user)
}

func (uuc userUsecase) IsUserValid(userId int) bool {
	user := uuc.userRepo.GetUser(userId)
	return user.Age > 18
}

func (uuc userUsecase) Hit(gameId int, faceUp bool) (game models.Game) {
	game = uuc.gameRepo.GetGame(gameId)
	fmt.Printf("user %d hits!", game.User.Id)
	if game.GameState != utils.Playing {
		panic(fmt.Sprintf("Game is already over. User %s", game.GameState))
	}
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	card.IsFaceUp = faceUp
	// assign the new cards to the user's hand
	hand := game.User.Hand
	hand.Cards = append(hand.Cards, card)
	if card.IsFaceUp {
		hand.Score = hand.RecalculateScore()
	}
	// check if burst happened
	if hand.Score > 21 {
		game.GameState = utils.Lost
	}
	game = uuc.gameRepo.UpdateGame(game)
	fmt.Println("user hand", game.User.Hand, game.GameState)
	return
}

func (uuc userUsecase) Stand(gameId int) (game models.Game) {
	game = uuc.gameRepo.GetGame(gameId)
	fmt.Printf("user %d stands!", game.User.Id)
	game.LastUserAction = utils.Stand
	uuc.gameRepo.UpdateGame(game)
	//call dealers turn
	game = game.Dealer.Play(game)
	return
}

func (uuc userUsecase) DoubleDown(gameId int) (game models.Game) {
	fmt.Println("user doubleDowns!")
	game = uuc.gameRepo.GetGame(gameId)
	if len(game.User.Hand.Cards) != 2 {
		panic(`User can't Double Down. This move is only 
		available when he only has the first 2 starting cards`)
	} else if game.User.Chips < game.Bet {
		panic(`User can't Double Down. His current chips balance is lower than necessary`)
	}
	game.User.Chips -= game.Bet
	game.Bet += game.Bet
	game.LastUserAction = utils.DoubleDown
	game = uuc.Hit(gameId, true)
	//call dealers turn
	// game = uuc.dealerUsecase.
	return
}
