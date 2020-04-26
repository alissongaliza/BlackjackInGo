package models

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type Difficuty string

const (
	easy   Difficuty = "easy"
	medium Difficuty = "medium"
	hard   Difficuty = "hard"
)

type Dealer interface {
	User
	Play(Game) Game
}

type EasyDealer struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type MediumDealer struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type HardDealer struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

func (easy *EasyDealer) Play(currentGame Game) Game {
	if easy.Hand.Score <= 17 {
		return easy.Hit(currentGame.Id, true)
	} else {
		return easy.Stand(currentGame.Id)
	}
}

func (easy *EasyDealer) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("EasyDealer hit reached")
	game = GetGameDb().Get(gameId)
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	if faceUp {
		card.isFaceUp = true
	}
	// assign the new cards to the dealers's hand
	dealer := game.Dealer.(*EasyDealer)
	hand := dealer.Hand
	hand.Cards = append(hand.Cards, card)
	hand.Score += card.value(*hand)
	// finish turn
	game.isPlayerTurn = true
	fmt.Println(dealer.Hand)
	GetGameDb().Update(game)
	return
}

func (easy EasyDealer) Stand(gameId int) (game Game) {
	fmt.Println("Dealer stands!")
	game = GetGameDb().Get(gameId)
	game.LastDealerAction = stand
	game.isPlayerTurn = true
	GetGameDb().Update(game)
	return
}

func (medium *MediumDealer) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("MediumDealer hit!")
	return
}
func (medium *MediumDealer) Stand(gameId int) (game Game) {
	fmt.Println("MediumDealer stand!")
	return
}

func (hard *HardDealer) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("HardDealer hit!")
	return
}

func (hard *HardDealer) Stand(gameId int) (game Game) {
	fmt.Println("HardDealer stand!")
	return
}
