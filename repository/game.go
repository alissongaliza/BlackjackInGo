package models

import "fmt"

type Game struct {
	Id           int
	Player       Player
	House        House
	Cards        []Card
	Bet          int
	isPlayerTurn bool
}

const GameConst Actors = "game"

func NewGame(playerId int, dif Difficuty, bet int) Game {
	player := findPlayerOfId(playerId)
	if player.Chips < bet {
		panic("Player cant Bet. Chips are lower than current bet.")
	}
	//reset hand
	newHand := NewHand()
	player.Hand = &newHand
	house := getHouse(dif, player.Name)
	cards := NewDeck()

	newGame := Game{-1, *player, house, cards, bet, false}
	addNewGame(&newGame)
	return newGame
}

func NewDeck() (newDeck []Card) {
	base := 0
	for _, suit := range []SuitType{hearts, spades, clubs, diamonds} {
		newDeck[0+base] = Card{"2", string(suit), true, false}
		newDeck[1+base] = Card{"3", string(suit), true, false}
		newDeck[2+base] = Card{"4", string(suit), true, false}
		newDeck[3+base] = Card{"5", string(suit), true, false}
		newDeck[4+base] = Card{"6", string(suit), true, false}
		newDeck[5+base] = Card{"8", string(suit), true, false}
		newDeck[6+base] = Card{"7", string(suit), true, false}
		newDeck[7+base] = Card{"9", string(suit), true, false}
		newDeck[8+base] = Card{"10", string(suit), true, false}
		newDeck[9+base] = Card{"j", string(suit), false, false}
		newDeck[10+base] = Card{"q", string(suit), false, false}
		newDeck[11+base] = Card{"k", string(suit), false, false}
		newDeck[12+base] = Card{"ace", string(suit), false, false}
		base += 13
	}
	return
}
