package models

import (
	"fmt"
	"sync"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type PlayerDb map[int]*Player

var playerInstance PlayerDb

var playerOnce sync.Once

func GetPlayerDb() PlayerDb {
	playerOnce.Do(func() {

		playerInstance = make(PlayerDb)

		hand1 := NewHand()
		playerInstance[1] = &Player{"Alisson", 1, 18, &hand1, 100}

	})

	return playerInstance
}

func assignPlayerId(player *Player) {
	players := GetPlayerDb()
	if players == nil {
		player.Id = 0
	}

	keys := utils.GetMapIntKeys(players)
	player.Id = utils.FindMaxIndex(keys)
}

func (players PlayerDb) Create(newPlayer Player) Player {
	assignPlayerId(&newPlayer)
	players[newPlayer.Id] = &newPlayer
	return newPlayer
}

func IsPlayerValid(playerId int) bool {
	players := GetPlayerDb()

	if player, present := players[playerId]; present && player.Age >= 18 {
		return true
	}
	return false
}

func (db PlayerDb) Get(id int) Player {
	players := GetPlayerDb()

	player, ok := players[id]
	if !ok {
		panic(fmt.Sprintf("Player of id %d not found", id))
	}
	return *player
}

func (db *PlayerDb) Update(player Player) Player {
	players := GetPlayerDb()

	players[player.Id] = &player
	return player
}
