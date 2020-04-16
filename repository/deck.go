package models

type SuitType string

const (
	hearts   SuitType = "hearts"
	spades   SuitType = "spades"
	clubs    SuitType = "clubs"
	diamonds SuitType = "diamonds"
)

type Card struct {
	suit SuitType
	name string
}

type Hand struct {
	cards []*Card
	score int
}
