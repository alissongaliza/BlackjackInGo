package models

import (
	"fmt"
	"sync"
)

// var settings = postgresql.ConnectionURL{
// 	Host:     "db",
// 	Database: "BlackJackDB",
// 	User:     "bj",
// 	Password: "bjIsFun",
// }

// var Sess, Err = postgresql.Open(settings)
type DB map[Actors]map[int]interface{}

type Actors string

var once sync.Once

var instance DB

func GetDb() DB {
	once.Do(func() {

		instance = make(DB, 2)
		instance[PlayerConst] = make(map[int]interface{})
		instance[GameConst] = make(map[int]interface{})
		instance[PlayerConst][1] = Player{"Alisson", 1, 18, nil, 100}

		player := instance[PlayerConst][1].(Player)
		house := getHouse("easy", player.Name)
		cards := NewDeck()
		instance[GameConst][1] = Game{1, player, house, cards, 0, false}

	})

	return instance
}

func getMapKeys(m map[int]interface{}) []int {

	keys := make([]int, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func findNextId(key Actors) int {
	db := GetDb()
	if db == nil {
		return -1
	}

	collection := db[key]
	keys := getMapKeys(collection)

	// find max index
	max := 0
	for k, e := range keys {
		if k == 0 || e > max {
			max = e
		}
	}
	max += 1
	return max
}

func addNewPlayer(newPlayer *Player) bool {
	db := GetDb()

	nextId := findNextId(PlayerConst)
	players := db[PlayerConst]
	newPlayer.Id = nextId
	players[nextId] = *newPlayer
	return true
}

func addNewGame(newGame *Game) bool {
	db := GetDb()

	nextId := findNextId(GameConst)
	games := db[GameConst]
	newGame.Id = nextId
	games[nextId] = *newGame
	return true
}

func IsPlayerValid(playerId int) bool {
	fmt.Println("IsPlayerValid")
	db := GetDb()
	// fmt.Println(db)

	if value, present := db[PlayerConst][playerId].(Player); present && value.Age >= 18 {
		return true
	}
	return false
}

func findPlayerOfId(id int) *Player {
	db := GetDb()

	player, ok := db[PlayerConst][id].(Player)
	if !ok {
		panic(fmt.Sprintf("Player of id %d not found", id))
	}
	return &player
}

func findGameOfId(id int) *Game {
	db := GetDb()

	game, ok := db[GameConst][id].(Game)
	if !ok {
		panic(fmt.Sprintf("Game of id %d not found", id))
	}
	return &game
}

func getHouse(dif Difficuty, opName string) (house House) {
	hand := NewHand()
	switch dif {
	case easy:
		house = EasyHouse{opName, easy, &hand}
	case medium:
		house = MediumHouse{opName, medium, &hand}
	case hard:
		house = HardHouse{opName, hard, &hand}
	}

	return
}

func StartGame(game *Game) Game {
	game.House.hit(game.Id, true)
	game.Player.hit(game.Id, true)
	game.House.hit(game.Id, false)
	game.Player.hit(game.Id, true)
	return *game
}
