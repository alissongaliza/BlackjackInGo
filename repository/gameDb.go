package models

import (
	"fmt"
	"sync"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type GameDb map[int]Game

var gameInstance GameDb

var gameOnce sync.Once

func GetGameDb() GameDb {
	gameOnce.Do(func() {

		gameInstance = make(GameDb)

		playerDb := GetPlayerDb()
		player := playerDb.Get(1)
		houseHand := NewHand()
		opponentName := fmt.Sprintf("%s's opponent", player.Name)
		house := EasyHouse{opponentName, easy, &houseHand}
		cards := NewDeck()
		gameInstance[1] = Game{1, player, &house, cards, 0, false, noAction, noAction, playing}

	})

	return gameInstance
}

func assignGameId(game *Game) {
	games := GetGameDb()
	if games == nil {
		game.Id = 0
	}

	keys := utils.GetMapIntKeys(games)
	game.Id = utils.FindMaxIndex(keys)
}

func (games GameDb) Create(newGame Game) Game {
	assignGameId(&newGame)
	games[newGame.Id] = newGame
	return newGame
}

func IsGameValid(gameId int) bool {
	games := GetGameDb()

	if _, present := games[gameId]; present {
		return true
	}
	return false
}

func (games GameDb) Get(id int) Game {
	game, ok := games[id]
	if !ok {
		panic(fmt.Sprintf("Game of id %d not found", id))
	}
	return game
}

func (games GameDb) Update(game Game) Game {
	games[game.Id] = game
	return game
}

func getHouse(dif Difficuty, opponentName string) (house House) {
	hand := NewHand()
	switch dif {
	case easy:
		house = &EasyHouse{opponentName, easy, &hand}
	case medium:
		house = &MediumHouse{opponentName, medium, &hand}
	case hard:
		house = &HardHouse{opponentName, hard, &hand}
	}
	return
}
