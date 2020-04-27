package models

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type User struct {
	Name  string
	Id    int
	Age   int
	Hand  *Hand
	Chips int
}

func NewUser(name string, age int) (newUser User) {
	hand := NewHand()
	newUser = User{name, -1, age, &hand, 100}
	return
}

func (user User) Hit(gameId int, faceUp bool) (game Game) {
	fmt.Printf("user %d hits!", user.Id)
	game = GetGameDb().Get(gameId)
	fmt.Println(game.GameState)
	if game.GameState != playing {
		panic(fmt.Sprintf("Game is already over. User %s", game.GameState))
	}
	//pop an element from the cards array
	index := utils.GetRandomNumber(0, len(game.Cards)-1)
	card := game.Cards[index]
	game.Cards[index] = game.Cards[len(game.Cards)-1]
	game.Cards = game.Cards[:len(game.Cards)-1]
	
	card.isFaceUp = faceUp
	// assign the new cards to the user's hand
	hand := game.User.Hand
	if hand.Score+card.value(*hand) > 21 {
		game.GameState = lost
	}
	hand.Cards = append(hand.Cards, card)
	hand.Score += card.value(*hand)
	// finish turn
	game.isUserTurn = false
	fmt.Println("user", game.User.Hand, game.GameState)
	//call dealers turn
	GetGameDb().Update(game)
	game = game.Dealer.Play(game)
	return
}

func (user User) Stand(gameId int) (game Game) {
	fmt.Printf("user %d stands!", user.Id)
	game = GetGameDb().Get(gameId)
	game.LastUserAction = stand
	game.isUserTurn = false
	//call dealers turn
	GetGameDb().Update(game)
	game = game.Dealer.Play(game)
	return
}

func (user User) DoubleDown(gameId int) (game Game) {
	fmt.Println("user doubleDowns!")
	game = GetGameDb().Get(gameId)
	if len(game.User.Hand.Cards) != 2 {
		panic(`User can't Double Down. This move is only 
		available when he has only the first 2 starting cards`)
	} else if game.User.Chips < game.Bet {
		panic(`User can't Double Down. His current chips balance is lower than necessary`)
	}
	game.User.Chips -= game.Bet
	game.Bet += game.Bet
	game.LastUserAction = doubleDown
	game = game.User.Hit(gameId, true)
	game.isUserTurn = false
	//call dealers turn
	GetGameDb().Update(game)
	game = game.Dealer.Play(game)
	return
}
