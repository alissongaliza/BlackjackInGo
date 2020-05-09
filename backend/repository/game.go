package models

import (
	"github.com/alissongaliza/BlackjackInGo/utils"
)

const GameConst utils.Actors = "game"

func NewGame(userId int, dif utils.Difficulty, bet int) (newGame Game) {
	user := GetUserDb().Get(userId)
	if user.Chips < bet {
		panic("User cant Bet. Chips are lower than current bet.")
	}
	//reset hand
	newHand := NewHand()
	user.Hand = &newHand
	dealer := getDealer(dif, user.Name)
	cards := NewDeck()

	newGame = Game{-1, user, dealer, cards, bet, 0, utils.NoAction, utils.NoAction, utils.Setup}
	return
}

func NewDeck() (newDeck []Card) {
	for _, suit := range []utils.SuitType{utils.Hearts, utils.Spades, utils.Clubs, utils.Diamonds} {
		newDeck = append(newDeck,
			Card{suit, "2", true, false},
			Card{suit, "3", true, false},
			Card{suit, "4", true, false},
			Card{suit, "5", true, false},
			Card{suit, "6", true, false},
			Card{suit, "8", true, false},
			Card{suit, "7", true, false},
			Card{suit, "9", true, false},
			Card{suit, "10", true, false},
			Card{suit, "J", false, false},
			Card{suit, "Q", false, false},
			Card{suit, "K", false, false},
			Card{suit, "ACE", false, false},
		)
	}
	return
}

func StartGame(gameId int) (game Game) {
	game = GetGameDb().Get(gameId)
	game.GameState = utils.Playing
	GetGameDb().Update(game)
	game.User.Hit(game.Id, true)
	game.User.Hit(game.Id, true)
	game.Dealer.Hit(game.Id, true)
	game.Dealer.Hit(game.Id, false)
	//set the dealer's last given card face down and recalculate score
	dealerHand := game.Dealer.Hand
	dealerHand.Cards[1].IsFaceUp = false
	dealerHand.Score = recalculateHandScore(*dealerHand)
	GetGameDb().Update(game)
	return
}

func calculatePayouts(game *Game) {
	user := game.User
	winnings := 0
	switch game.GameState {
	case utils.Won:
		{
			if game.LastUserAction == utils.DoubleDown {
				winnings = game.Bet * 4
			} else {
				winnings = game.Bet * 2
			}
		}
	case utils.Drew:
		{
			if user.Hand.Score == 21 {
				winnings = int(float64(game.Bet) * 1.5)
			} else {
				winnings = game.Bet
			}
		}
	}
	user.Chips += winnings
	game.Payout = winnings
}
