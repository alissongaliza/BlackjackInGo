package models

type Player struct {
	Hand Hand
}

type User struct {
	Player
	Name  string
	Id    int
	Age   int
	Chips int
}
