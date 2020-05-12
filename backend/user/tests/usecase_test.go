package tests

import (
	"testing"

	"github.com/alissongaliza/BlackjackInGo/utils"

	dealerUsecase "github.com/alissongaliza/BlackjackInGo/backend/dealer/usecases"
	gameRep "github.com/alissongaliza/BlackjackInGo/backend/game/repositories"
	gameUsecase "github.com/alissongaliza/BlackjackInGo/backend/game/usecases"
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	userRep "github.com/alissongaliza/BlackjackInGo/backend/user/repositories"
	userUsecase "github.com/alissongaliza/BlackjackInGo/backend/user/usecases"
)

var (
	userRepo    = userRep.NewInMemoryUserDb()
	gameRepo    = gameRep.NewInMemoryGameDb()
	dealerUcase = dealerUsecase.NewDealerUsecase(gameRepo, userRepo)
	userUcase   = userUsecase.NewUserUsecase(gameRepo, userRepo, dealerUcase)
	gameUcase   = gameUsecase.NewGameUsecase(gameRepo, dealerUcase, userUcase, userRepo)
)

func mockGame(name string, age int, diff utils.Difficulty, bet int, blackJack bool) models.Game {
	mockUser := userUcase.CreateUser(name, age)
	mockGame, _ := gameUcase.CreateGame(mockUser, diff, bet)
	mockGame, _ = gameUcase.StartNewGame(mockGame)
	if blackJack {
		// manually setting the combination for a blackjack
		mockGame.User.Hand.Cards[0] = models.Card{Suit: utils.Diamonds,
			Name: "10", IsNumber: true, IsFaceUp: true}
		mockGame.User.Hand.Cards[1] = models.Card{Suit: utils.Diamonds,
			Name: "Ace", IsNumber: false, IsFaceUp: true}
	} else {
		mockGame.User.Hand.Cards[0] = models.Card{Suit: utils.Diamonds,
			Name: "2", IsNumber: true, IsFaceUp: true}
		mockGame.User.Hand.Cards[1] = models.Card{Suit: utils.Spades,
			Name: "2", IsNumber: true, IsFaceUp: true}
	}
	mockGame.User.Hand.RecalculateScore()
	return mockGame
}
func TestHit(t *testing.T) {
	game1 := mockGame("alisson", 18, utils.Easy, 50, false)
	resultGame1, _ := userUcase.Hit(game1, true)
	if !(len(resultGame1.User.Hand.Cards) > len(game1.User.Hand.Cards) &&
		resultGame1.User.Hand.Score > game1.User.Hand.Score) &&
		resultGame1.LastUserAction == utils.Hit &&
		resultGame1.GameState == utils.Playing {
		t.Error("User Hit faceUp test failed")
	} else {
		t.Log("User Hit faceUp test passed")
	}

	game2 := mockGame("alisson", 18, utils.Easy, 50, false)
	resultGame2, _ := userUcase.Hit(game2, false)
	if !(len(resultGame2.User.Hand.Cards) > len(game2.User.Hand.Cards) &&
		resultGame2.User.Hand.Score == game2.User.Hand.Score) &&
		resultGame1.LastUserAction == utils.Hit &&
		resultGame1.GameState == utils.Playing {
		t.Error("User Hit faceDown test failed")
	} else {
		t.Log("User Hit faceDown test passed")
	}
}

func TestStand(t *testing.T) {
	game1 := mockGame("alisson", 18, utils.Easy, 50, false)
	resultGame1 := userUcase.Stand(game1)
	if !(len(resultGame1.User.Hand.Cards) == len(game1.User.Hand.Cards) &&
		resultGame1.User.Hand.Score == game1.User.Hand.Score) &&
		resultGame1.LastUserAction == utils.Stand &&
		resultGame1.GameState == utils.Playing {
		t.Error("User Stand test failed")
	} else {
		t.Log("User Stand test passed")
	}
}

func TestDoubleDown(t *testing.T) {
	game1 := mockGame("alisson", 18, utils.Easy, 50, false)
	resultGame1, _ := userUcase.DoubleDown(game1)
	if !(len(resultGame1.User.Hand.Cards) > len(game1.User.Hand.Cards) &&
		resultGame1.User.Hand.Score > game1.User.Hand.Score) &&
		resultGame1.LastUserAction == utils.DoubleDown &&
		resultGame1.GameState == utils.Playing {
		t.Error("User DoubleDown ok test failed")
	} else {
		t.Log("User DoubleDown ok test passed")
	}

	game2 := mockGame("alisson", 18, utils.Easy, 40, false)
	// simulating adding a third card to test doubledown 2 card max rule
	card := models.Card{Suit: utils.Diamonds, Name: "3", IsNumber: true, IsFaceUp: true}
	game2.User.Hand.Cards = append(game2.User.Hand.Cards, card)
	_, err := userUcase.DoubleDown(game2)
	if err != nil &&
		err.Error() == `User can't Double Down. This move is only available when he only has the first 2 starting cards` {
		t.Log("User DoubleDown 2 cards max test passed")
	} else {
		t.Error("User DoubleDown 2 cards max test  failed")
	}

	game3 := mockGame("alisson", 18, utils.Easy, 60, false)
	_, err2 := userUcase.DoubleDown(game3)
	if err2 != nil &&
		err2.Error() == `User can't Double Down. His current chips balance is lower than necessary` {
		t.Log("User DoubleDown not enough chips passed")
	} else {
		t.Error("User DoubleDown not enough chips  failed")
	}
}

func TestIsUserValid(t *testing.T) {
	user1 := userUcase.CreateUser("alisson", 18)
	if isValid := userUcase.IsUserValid(user1); !isValid {
		t.Error("User valid test failed")
	} else {
		t.Log("User valid test passed")
	}
	user2 := userUcase.CreateUser("alisson", 17)
	if isValid := userUcase.IsUserValid(user2); isValid {
		t.Error("User invalid test failed")
	} else {
		t.Log("User invalid test passed")
	}
}

func TestCreateUser(t *testing.T) {
	var (
		name    = "alisson"
		age     = 18
		hand    = gameRepo.CreateHand()
		initial = 100
	)
	user1 := userUcase.CreateUser(name, age)
	if user1.Age == age &&
		user1.Name == name &&
		len(user1.Hand.Cards) == len(hand.Cards) &&
		user1.Hand.Score == 0 &&
		user1.Id > 0 &&
		user1.Chips == initial {
		t.Log("User create test passed")
	} else {
		t.Error("User create test failed")
	}
}
