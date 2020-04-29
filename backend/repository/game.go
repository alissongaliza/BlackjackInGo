package models

type Game struct {
	Id               int
	User             User
	Dealer           Dealer
	Cards            []Card
	Bet              int
	isUserTurn       bool
	LastUserAction   Action
	LastDealerAction Action
	GameState        GameState
}

type GameState string
type Action string

const (
	won        GameState = "won"
	lost       GameState = "lost"
	drew       GameState = "drew"
	playing    GameState = "playing"
	noAction   Action    = "noAction"
	hit        Action    = "hit"
	stand      Action    = "stand"
	doubleDown Action    = "doubleDown"
)

const GameConst Actors = "game"

func NewGame(userId int, dif Difficuty, bet int) (newGame Game) {
	user := GetUserDb().Get(userId)
	if user.Chips < bet {
		panic("User cant Bet. Chips are lower than current bet.")
	}
	//reset hand
	newHand := NewHand()
	user.Hand = &newHand
	dealer := getDealer(dif, user.Name)
	cards := NewDeck()

	newGame = Game{-1, user, dealer, cards, bet, false, noAction, noAction, playing}
	return
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
