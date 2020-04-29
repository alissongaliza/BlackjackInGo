package models

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/backend/utils"
)

type Difficuty string

const (
	easy   Difficuty = "easy"
	broken Difficuty = "broken"
)

type Dealer interface {
	Play(Game) Game
}

type EasyDealer struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type BrokenDealer struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

func (easy *EasyDealer) Play(currentGame Game) Game {
	if currentGame.GameState != playing {
		return currentGame
	}
	if easy.Hand.Score <= 17 {
		return Hit(currentGame.Id, true)
	} else {
		return Stand(currentGame.Id)
	}
}

func (broken *BrokenDealer) Play(currentGame Game) Game {
	if broken.Hand.Score <= 17 &&
		currentGame.User.Hand.Score > currentGame.User.Hand.Score {
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

	card.isFaceUp = faceUp
	// assign the new cards to the dealers's hand
	dealer := game.Dealer.(*EasyDealer)
	hand := dealer.Hand
	if hand.Score+card.value(*hand) > 21 {
		game.GameState = lost
	}
	hand.Cards = append(hand.Cards, card)
	hand.Score += card.value(*hand)
	// finish turn
	game.isUserTurn = true
	game.LastDealerAction = hit
	fmt.Println("dealer", dealer.Hand, game.GameState)
	GetGameDb().Update(game)
	return
}

func Stand(gameId int) (game Game) {
	fmt.Println("Dealer stands!")
	game = GetGameDb().Get(gameId)
	game.LastDealerAction = stand
	game.isUserTurn = true
	game.LastDealerAction = stand
	dealer := game.Dealer.(*EasyDealer)
	fmt.Println("dealer", dealer.Hand, game.GameState)
	GetGameDb().Update(game)
	return
}
