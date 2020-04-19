package models

import "fmt"

type Player struct {
	Name string
	Id   int
	Age  int
	Hand *Hand
}

const PlayerConst Actors = "player"

func NewPlayer(name string, age int) (newPlayer Player) {
	cards := make([]Card, 0)
	hand := Hand{cards, 0}
	newPlayer = Player{name, -1, age, &hand}
	addNewPlayer(&newPlayer)
	return newPlayer
}

func (player Player) hit(gameId int) {
	fmt.Println("player hit!")

}

func (player Player) stand(gameId int) {
	fmt.Println("player stand!")

}

func (player Player) doubleDown(gameId int) {
	fmt.Println("player hit!")

}

// func (player *Player) SetId(id int) {
// 	player.Id = id
// }

// func (player *Player) SetName(name string) {
// 	player.Name = name
// }

// func (player Player) GetId() int {
// 	return player.Id
// }

// func (player Player) GetName() string {
// 	return player.Name
// }
