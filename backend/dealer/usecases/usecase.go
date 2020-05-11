package usecases

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/backend/dealer"
	"github.com/alissongaliza/BlackjackInGo/backend/game"
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type dealerUsecase struct {
	gameRepo game.Repository
}

func NewDealerUsecase(gameRepo game.Repository) dealer.UseCase {
	return &dealerUsecase{
		gameRepo: gameRepo,
	}
}

func (duc dealerUsecase) CreateDealer(dif utils.Difficulty) models.Dealer {
	hand := duc.gameRepo.CreateHand()
	switch dif {
	case utils.Easy:
		return models.Dealer{DealerActions: &models.EasyDealer{},
			Player: models.Player{Hand: hand}, Difficulty: utils.Easy,
		}
	case utils.Broken:
		return models.Dealer{DealerActions: &models.BrokenDealer{},
			Player: models.Player{Hand: hand}, Difficulty: utils.Broken,
		}
	default:
		panic("Invalid dificulty")
	}
}

func (duc dealerUsecase) AutoPlay(currentGame models.Game) models.Game {
	return currentGame.Dealer.Play(currentGame)
}

func (duc dealerUsecase) Hit(gameId int, faceUp bool) (game models.Game) {
	fmt.Println("Dealer hit reached")
	game = duc.gameRepo.GetGame(gameId)
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	card.IsFaceUp = faceUp
	// assign the new cards to the dealers's hand
	dealer := game.Dealer
	hand := dealer.Hand
	hand.Cards = append(hand.Cards, card)
	if card.IsFaceUp {
		hand.Score = hand.RecalculateScore()
	}
	if hand.Score > 21 {
		//user won
		game.GameState = utils.Won
	}
	game.LastDealerAction = utils.Hit
	duc.gameRepo.UpdateGame(game)
	return
}

func (duc dealerUsecase) Stand(gameId int) (game models.Game) {
	fmt.Println("Dealer stands!")
	game = duc.gameRepo.GetGame(gameId)
	game.LastDealerAction = utils.Stand
	if game.User.Hand.Score > game.Dealer.Hand.Score || game.Dealer.Hand.Score > 21 {
		//user won
		game.GameState = utils.Won
	} else if game.User.Hand.Score == game.Dealer.Hand.Score {
		game.GameState = utils.Drew
	} else {
		// dealer won
		game.GameState = utils.Lost
	}
	return duc.gameRepo.UpdateGame(game)
}

// func (*BrokenDealer) Play(currentGame Game) Game {
// 	duc := dealerUsecase.New
// 	currentGame.Dealer.Hand.Cards[1].IsFaceUp = true
// 	currentGame.Dealer.Hand.Score()
// 	for currentGame.GameState == utils.Playing {
// 		if currentGame.LastDealerAction == utils.Stand {
// 			currentGame = dus.Stand(currentGame.Id)
// 		} else if currentGame.Dealer.Player.Hand.Score <= 17 &&
// 			currentGame.User.Hand.Score > currentGame.User.Hand.Score {
// 			currentGame = dus.Hit(currentGame.Id, true)
// 		} else {
// 			currentGame = dus.Stand(currentGame.Id)
// 		}
// 	}
// 	return currentGame
// }

// func (*EasyDealer) Play(currentGame Game) Game {
// 	duc := dealerUsecase.New
// 	currentGame.Dealer.Hand.Cards[1].IsFaceUp = true
// 	currentGame.Dealer.Hand.Score()
// 	for currentGame.GameState == utils.Playing {
// 		if currentGame.LastDealerAction == utils.Stand {
// 			currentGame = dus.Stand(currentGame.Id)
// 		} else if currentGame.Dealer.Player.Hand.Score <= 17 {
// 			currentGame = dus.Hit(currentGame.Id, true)
// 		} else {
// 			currentGame = dus.Stand(currentGame.Id)
// 		}
// 	}
// 	calculatePayouts(&currentGame)
// 	return currentGame
// }
