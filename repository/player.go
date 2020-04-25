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

func NewPlayer(name string, age int) (newPlayer Player) {
	hand := NewHand()
	newPlayer = Player{name, -1, age, &hand, 100}
	return
}

func (player *Player) Hit(gameId int, faceUp bool) {
	game := GetGameDb().Get(gameId)
	fmt.Println("player hit!", gameId, game.GameState)
	if game.GameState != playing {
		panic(fmt.Sprintf("Game is already over. Player %s", game.GameState))
	}
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	if faceUp {
		card.isFaceUp = true
	}
	// assign the new cards to the player's hand
	hand := game.Player.Hand
	if hand.Score+card.value(*hand) > 21 {
		game.GameState = lost
	}
	hand.Cards = append(hand.Cards, card)
	hand.Score += card.value(*hand)
	// finish turn
	game.isPlayerTurn = false
	fmt.Println(player.Hand, game.GameState)
}

func (player Player) Stand(gameId int) {
	fmt.Println("player stand!")
}

func (player Player) DoubleDown(gameId int) {
	fmt.Println("player doubleDown!")
}
