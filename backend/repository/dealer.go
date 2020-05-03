package models

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type EasyDealer struct {
}

type BrokenDealer struct {
}

func (*BrokenDealer) Play(currentGame Game) Game {
	currentGame.Dealer.Hand.Cards[1].IsFaceUp = true
	currentGame.Dealer.Hand.Score = recalculateHandScore(*currentGame.Dealer.Hand)
	for currentGame.GameState == utils.Playing {
		if currentGame.LastDealerAction == utils.Stand {
			currentGame = currentGame.Dealer.Stand(currentGame.Id)
		} else if currentGame.Dealer.Player.Hand.Score <= 17 &&
			currentGame.User.Hand.Score > currentGame.User.Hand.Score {
			currentGame = currentGame.Dealer.Hit(currentGame.Id, true)
		} else {
			currentGame = currentGame.Dealer.Stand(currentGame.Id)
		}
	}
	return currentGame
}

func (*EasyDealer) Play(currentGame Game) Game {
	currentGame.Dealer.Hand.Cards[1].IsFaceUp = true
	currentGame.Dealer.Hand.Score = recalculateHandScore(*currentGame.Dealer.Hand)
	for currentGame.GameState == utils.Playing {
		if currentGame.LastDealerAction == utils.Stand {
			currentGame = currentGame.Dealer.Stand(currentGame.Id)
		} else if currentGame.Dealer.Player.Hand.Score <= 17 {
			currentGame = currentGame.Dealer.Hit(currentGame.Id, true)
		} else {
			currentGame = currentGame.Dealer.Stand(currentGame.Id)
		}
	}
	calculatePayouts(&currentGame)
	return currentGame
}

func (Dealer) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("Dealer hit reached")
	game = GetGameDb().Get(gameId)
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	card.IsFaceUp = faceUp
	// assign the new cards to the dealers's hand
	dealer := game.Dealer
	hand := dealer.Hand
	if hand.Score+card.value(*hand) > 21 {
		//user won
		game.GameState = utils.Won
	}
	hand.Cards = append(hand.Cards, card)
	if card.IsFaceUp {
		hand.Score += card.value(*hand)
	}
	game.LastDealerAction = utils.Hit
	fmt.Println("dealer", dealer.Hand, game.GameState)
	GetGameDb().Update(game)
	return
}

func (Dealer) Stand(gameId int) (game Game) {
	fmt.Println("Dealer stands!")
	game = GetGameDb().Get(gameId)
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
	GetGameDb().Update(game)
	return
}
