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

func (easy *EasyHouse) Hit(gameId int, faceUp bool) {
	fmt.Println("EasyHouse hit reached")
	game := GetGameDb().Get(gameId)
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
}

func (easy *EasyHouse) Stand(gameId int) {
	fmt.Println("EasyHouse stand!")
}

func (easy *EasyHouse) DoubleDown(gameId int) {
	fmt.Println("EasyHouse doubleDown!")
}

func (medium *MediumHouse) Hit(gameId int, faceUp bool) {
	fmt.Println("MediumHouse hit!")
}
func (medium *MediumHouse) Stand(gameId int) {
	fmt.Println("MediumHouse stand!")
}
func (medium *MediumHouse) DoubleDown(gameId int) {
	fmt.Println("MediumHouse doubleDown!")
}

func (hard *HardHouse) Hit(gameId int, faceUp bool) {
	fmt.Println("HardHouse hit!")
}

func (hard *HardHouse) Stand(gameId int) {
	fmt.Println("HardHouse stand!")
}

func (hard *HardHouse) DoubleDown(gameId int) {
	fmt.Println("HardHouse doubleDown!")
}
