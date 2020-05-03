package models

import "github.com/alissongaliza/BlackjackInGo/utils"

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

	newGame = Game{-1, user, dealer, cards, bet, 0, utils.NoAction, utils.NoAction, utils.Playing}
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
