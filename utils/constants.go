package utils

type SuitType string
type Difficuty string
type GameState string
type Action string
type Actors string

const (
	Won     GameState = "won"
	Lost    GameState = "lost"
	Drew    GameState = "drew"
	Playing GameState = "playing"

	NoAction   Action = "noAction"
	Hit        Action = "hit"
	Stand      Action = "stand"
	DoubleDown Action = "doubleDown"

	Hearts   SuitType = "hearts"
	Spades   SuitType = "spades"
	Clubs    SuitType = "clubs"
	Diamonds SuitType = "diamonds"

	Easy   Difficuty = "easy"
	Broken Difficuty = "broken"
)
