package models

type Game struct {
	id    int
	user  User
	house User
	cards [52]Card
}

const GameConst Actors = "game"

// func New(user User, nextId int) Game {
// 	// newHouse User :=
// 	// newGame game = game{user=user, id=nextId, house=newHouse, cards=cards}
// }
