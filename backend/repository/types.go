package models

import "github.com/alissongaliza/BlackjackInGo/utils"

type PlayerActions interface {
	Hit(int, bool) (game Game)
	Stand(int) (game Game)
}

type RealPlayer interface {
	DoubleDown(int) (game Game)
}

type DealerActions interface {
	Play(Game) Game
}

type Player struct {
	Hand *Hand
	PlayerActions
}

type Dealer struct {
	Player
	Difficuty utils.Difficuty
}

type User struct {
	Player
	RealPlayer
	Name  string
	Id    int
	Age   int
	Chips int
}

type Game struct {
	Id     int
	User   User
	Dealer Dealer
	Cards  []Card
	Bet    int
	// isUserTurn       bool
	LastUserAction   utils.Action
	LastDealerAction utils.Action
	GameState        utils.GameState
}

type Card struct {
	Suit     utils.SuitType
	Name     string
	isNumber bool
	isFaceUp bool
}

type Hand struct {
	Cards []Card
	Score int
}

type base interface {
	Save(model base) base
	Get(modelId int)
	Update(model base) base
}