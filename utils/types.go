package utils

type SuitType string
type Difficuty string
type GameState string
type Action string

type PlayerActions interface {
	Hit(gameId int, faceUp bool) (game Game)
	Stand(gameId int) (game Game)
}

type RealPlayer interface {
	DoubleDown(gameId int) (game Game)
}

type Player struct {
	Hand *Hand
	PlayerActions
}

type Dealer struct {
	Player
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
	LastUserAction   Action
	LastDealerAction Action
	GameState        GameState
}

type Card struct {
	Suit     SuitType
	Name     string
	isNumber bool
	isFaceUp bool
}

type Hand struct {
	Cards []Card
	Score int
}

type DealerActions interface {
	Play(Game) Game
}

type Actors string

type base interface {
	Save(model base) base
	Get(modelId int)
	Update(model base) base
}
