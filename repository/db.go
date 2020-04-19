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
		instance[PlayerConst][1] = Player{"Alisson", 1, 18, nil}

		cards := make([]Card, 0)
		hand := Hand{cards, 0}
		instance[GameConst][1] = EasyHouse{"test", easy, &hand}

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
		return nil
	}
	return &player
}

func getHouse(dif Difficuty, opName string) (house House) {

	switch dif {
	case easy:
		house = EasyHouse{opName, easy, nil}
	case medium:
		house = MediumHouse{opName, medium, nil}
	case hard:
		house = HardHouse{opName, hard, nil}
	}

	return
}
