package models

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

func NewUser(name string, age int) (newUser User) {
	hand := NewHand()
	//dummy id
	newUser = User{Player: Player{&hand}, Name: name, Id: -1, Age: age, Chips: 100}
	return
}

func (user User) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Printf("user %d hits!", user.Id)
	game = GetGameDb().Get(gameId)
	if game.GameState != utils.Playing {
		panic(fmt.Sprintf("Game is already over. User %s", game.GameState))
	}
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]

	card.IsFaceUp = faceUp
	// assign the new cards to the user's hand
	hand := game.User.Hand
	hand.Cards = append(hand.Cards, card)
	if card.IsFaceUp {
		hand.Score += card.value(*hand)
	}
	fmt.Println("user", game.User.Hand, game.GameState)
	// check if burst happened
	if hand.Score > 21 {
		game.GameState = utils.Lost
	}
	GetGameDb().Update(game)
	return
}

func (user User) Stand(gameId int) (game Game) {
	fmt.Printf("user %d stands!", user.Id)
	game = GetGameDb().Get(gameId)
	game.LastUserAction = utils.Stand
	GetGameDb().Update(game)
	//call dealers turn
	game = game.Dealer.Play(game)
	return
}

func (user User) DoubleDown(gameId int) (game Game) {
	fmt.Println("user doubleDowns!")
	game = GetGameDb().Get(gameId)
	if len(game.User.Hand.Cards) != 2 {
		panic(`User can't Double Down. This move is only 
		available when he only has the first 2 starting cards`)
	} else if game.User.Chips < game.Bet {
		panic(`User can't Double Down. His current chips balance is lower than necessary`)
	}
	game.User.Chips -= game.Bet
	game.Bet += game.Bet
	game.LastUserAction = utils.DoubleDown
	game = game.User.Hit(gameId, true)
	//call dealers turn
	GetGameDb().Update(game)
	game = game.Dealer.Play(game)
	return
}
