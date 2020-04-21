package models

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type Player struct {
	Name  string
	Id    int
	Age   int
	Hand  *Hand
	Chips int
}

const PlayerConst Actors = "player"

func NewPlayer(name string, age int) (newPlayer Player) {
	hand := NewHand()
	newPlayer = Player{name, -1, age, &hand, 100}
	addNewPlayer(&newPlayer)
	return newPlayer
}

func (player Player) Hit(gameId int, faceUp bool) {
	fmt.Println("player hit!")
	game := findGameOfId(gameId)
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards))
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	if faceUp {
		card.isFaceUp = true
	}
	// assign the new cards to the player's hand
	hand := game.Player.Hand
	hand.Cards = append(hand.Cards, card)
	hand.Score += card.value(*hand)
	// finish turn
	game.isPlayerTurn = false
	fmt.Println(hand)
}

func (player Player) Stand(gameId int) {
	fmt.Println("player stand!")
}

func (player Player) DoubleDown(gameId int) {
	fmt.Println("player doubleDown!")
}
