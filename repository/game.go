package models

import "fmt"

type Game struct {
	Id     int
	Player Player
	House  House
	Cards  [52]Card
}

const GameConst Actors = "game"

func NewGame(playerId int, dif Difficuty) Game {
	fmt.Println("New game called")
	player := findPlayerOfId(playerId)
	house := getHouse(dif, player.Name)
	cards := getNewDeck()

	newGame := Game{-1, *player, house, cards}
	addNewGame(&newGame)
	return newGame
}

func getNewDeck() (newDeck [52]Card) {
	base := 0
	for _, suit := range []SuitType{hearts, spades, clubs, diamonds} {
		newDeck[0+base] = Card{"2", string(suit), true}
		newDeck[1+base] = Card{"3", string(suit), true}
		newDeck[2+base] = Card{"4", string(suit), true}
		newDeck[3+base] = Card{"5", string(suit), true}
		newDeck[4+base] = Card{"6", string(suit), true}
		newDeck[5+base] = Card{"8", string(suit), true}
		newDeck[6+base] = Card{"7", string(suit), true}
		newDeck[7+base] = Card{"9", string(suit), true}
		newDeck[8+base] = Card{"10", string(suit), true}
		newDeck[9+base] = Card{"j", string(suit), false}
		newDeck[10+base] = Card{"q", string(suit), false}
		newDeck[11+base] = Card{"k", string(suit), false}
		newDeck[12+base] = Card{"ace", string(suit), false}
		base += 13
	}
	return
}
