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

func (player Player) Hit(gameId int, faceUp bool) (game Game) {
	game = GetGameDb().Get(gameId)
	fmt.Println("player hits!", gameId, game.GameState, game.Player.Hand)
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
	fmt.Println(game.Player.Hand, game.GameState)
	GetGameDb().Update(game)
	return
}

func (player Player) Stand(gameId int) (game Game) {
	fmt.Println("player stands!")
	game = GetGameDb().Get(gameId)
	game.LastPlayerAction = stand
	game.isPlayerTurn = false
	GetGameDb().Update(game)
	return
}

func (player Player) DoubleDown(gameId int) (game Game) {
	fmt.Println("player doubleDowns!")
	game = GetGameDb().Get(gameId)
	if len(game.Player.Hand.Cards) != 2 {
		panic(`Player can't Double Down. This move is only 
		available when he has only the first 2 starting cards`)
	} else if game.Player.Chips < game.Bet {
		panic(`Player can't Double Down. His current chips balance is lower than necessary`)
	}
	game.Player.Chips -= game.Bet
	game.Bet += game.Bet
	game.LastPlayerAction = doubleDown
	game = game.Player.Hit(gameId, true)
	game.isPlayerTurn = false
	GetGameDb().Update(game)
	return
}
