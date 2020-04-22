package models

type Game struct {
	Id           int
	Player       Player
	House        House
	Cards        []Card
	Bet          int
	isPlayerTurn bool
	GameState    GameState
}

type GameState string

const (
	won     GameState = "won"
	lost    GameState = "lost"
	draw    GameState = "draw"
	playing GameState = "playing"
)

const GameConst Actors = "game"

func NewGame(playerId int, dif Difficuty, bet int) Game {
	player := FindPlayerOfId(playerId)
	if player.Chips < bet {
		panic("Player cant Bet. Chips are lower than current bet.")
	}
	//reset hand
	newHand := NewHand()
	player.Hand = &newHand
	house := getHouse(dif, player.Name)
	cards := NewDeck()

	newGame := Game{-1, *player, house, cards, bet, false, playing}
	addNewGame(&newGame)
	return newGame
}

func NewDeck() (newDeck []Card) {
	for _, suit := range []SuitType{hearts, spades, clubs, diamonds} {
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
			Card{suit, "j", false, false},
			Card{suit, "q", false, false},
			Card{suit, "k", false, false},
			Card{suit, "ace", false, false},
		)
	}
	return
}
