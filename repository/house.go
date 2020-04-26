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

type House interface {
	User
}

type EasyHouse struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type MediumHouse struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type HardHouse struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

func (easy *EasyHouse) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("EasyHouse hit reached")
	game = GetGameDb().Get(gameId)
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	if faceUp {
		card.isFaceUp = true
	}
	// assign the new cards to the houses's hand
	house := game.House.(*EasyHouse)
	hand := house.Hand
	hand.Cards = append(hand.Cards, card)
	hand.Score += card.value(*hand)
	// finish turn
	game.isPlayerTurn = true
	fmt.Println(house.Hand)
	GetGameDb().Update(game)
	return
}

func (easy EasyHouse) Stand(gameId int) (game Game) {
	fmt.Println("House stands!")
	game = GetGameDb().Get(gameId)
	game.LastHouseAction = stand
	game.isPlayerTurn = true
	GetGameDb().Update(game)
	return
}

func (medium *MediumHouse) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("MediumHouse hit!")
	return
}
func (medium *MediumHouse) Stand(gameId int) (game Game) {
	fmt.Println("MediumHouse stand!")
	return
}

func (hard *HardHouse) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Println("HardHouse hit!")
	return
}

func (hard *HardHouse) Stand(gameId int) (game Game) {
	fmt.Println("HardHouse stand!")
	return
}
