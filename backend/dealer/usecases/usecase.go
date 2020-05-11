package usecases

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/backend/dealer"
	"github.com/alissongaliza/BlackjackInGo/backend/game"
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type EasyDealer models.Dealer
type BrokenDealer models.Dealer

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
		return models.Dealer{DealerActions: &EasyDealer{},
			Player: models.Player{Hand: hand}, Difficulty: utils.Easy,
		}
	case utils.Broken:
		return models.Dealer{DealerActions: &BrokenDealer{},
			Player: models.Player{Hand: hand}, Difficulty: utils.Broken,
		}
	default:
		panic("Invalid dificulty")
	}
}

func (duc dealerUsecase) AutoPlay(currentGame models.Game) models.Game {
	currentGame.Dealer.Hand.Cards[1].IsFaceUp = true
	currentGame.Dealer.Hand.RecalculateScore()
	for currentGame.GameState == utils.Playing {
		if currentGame.LastDealerAction == utils.Stand {
			currentGame = duc.Stand(currentGame.Id)
		} else {
			if currentGame.Dealer.Difficulty == utils.Broken {
				if currentGame.Dealer.Player.Hand.Score <= 17 &&
					currentGame.User.Hand.Score > currentGame.User.Hand.Score {
					currentGame = duc.Hit(currentGame.Id, true)
				} else {
					currentGame = duc.Stand(currentGame.Id)
				}
			} else {
				if currentGame.Dealer.Player.Hand.Score <= 17 {
					currentGame = duc.Hit(currentGame.Id, true)
				} else {
					currentGame = duc.Stand(currentGame.Id)
				}
			}
		}
	}
	currentGame.CalculatePayouts()
	return currentGame
}

func (duc dealerUsecase) Hit(gameId int, faceUp bool) models.Game {
	fmt.Println("Dealer hit reached")
	game := duc.gameRepo.GetGame(gameId)
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	card.IsFaceUp = faceUp
	// assign the new cards to the dealers's hand
	hand := &game.Dealer.Hand
	hand.Cards = append(hand.Cards, card)
	if card.IsFaceUp {
		hand.Score = hand.RecalculateScore()
	}
	if hand.Score > 21 {
		//user won
		game.GameState = utils.Won
	}
	fmt.Println("Dealer hand", hand)
	game.LastDealerAction = utils.Hit
	return duc.gameRepo.UpdateGame(game)
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
