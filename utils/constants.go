package utils

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

const (
	hearts   SuitType = "hearts"
	spades   SuitType = "spades"
	clubs    SuitType = "clubs"
	diamonds SuitType = "diamonds"
)

const (
	easy   Difficuty = "easy"
	broken Difficuty = "broken"
)
