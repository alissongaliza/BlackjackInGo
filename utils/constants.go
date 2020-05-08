package utils

type SuitType string
type Difficulty string
type GameState string
type Action string
type Actors string

const (
	Won     GameState = "won"
	Lost    GameState = "lost"
	Drew    GameState = "drew"
	Playing GameState = "playing"
	Setup   GameState = "setup"

	NoAction   Action = "utils.NoAction"
	Hit        Action = "hit"
	Stand      Action = "stand"
	DoubleDown Action = "doubleDown"

	Hearts   SuitType = "hearts"
	Spades   SuitType = "spades"
	Clubs    SuitType = "clubs"
	Diamonds SuitType = "diamonds"

	Easy   Difficulty = "easy"
	Broken Difficulty = "broken"
)
