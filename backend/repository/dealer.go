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
	if currentGame.Dealer.Player.Hand.Score <= 17 &&
		currentGame.User.Hand.Score > currentGame.User.Hand.Score {
		return Hit(currentGame.Id, true)
	} else {
		return Stand(currentGame.Id)
	}
}

func (*EasyDealer) Play(currentGame Game) Game {
	if currentGame.GameState != utils.Playing {
		return currentGame
	}
	if currentGame.Dealer.Player.Hand.Score <= 17 {
		return Hit(currentGame.Id, true)
	} else {
		return Stand(currentGame.Id)
	}
}

func Hit(gameId int, faceUp bool) (game Game) {
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
		game.GameState = utils.Lost
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

func Stand(gameId int) (game Game) {
	fmt.Println("Dealer stands!")
	game = GetGameDb().Get(gameId)
	game.LastDealerAction = utils.Stand
	dealer := game.Dealer
	fmt.Println("dealer", dealer.Hand, game.GameState)
	GetGameDb().Update(game)
	return
}
